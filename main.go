package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/common"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/command/v2"
	"github.com/liamawhite/cli-with-i18n/util/configv3"
	"github.com/liamawhite/cli-with-i18n/util/panichandler"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	log "github.com/Sirupsen/logrus"
)

type UI interface {
	DisplayError(err error)
	DisplayWarning(template string, templateValues ...map[string]interface{})
}

type DisplayUsage interface {
	DisplayUsage()
}

var ErrFailed = errors.New("command failed")
var ParseErr = errors.New("incorrect type for arg")

func main() {
	defer panichandler.HandlePanic()
	parse(os.Args[1:])
}

func parse(args []string) {
	parser := flags.NewParser(&common.Commands, flags.HelpFlag)
	parser.CommandHandler = executionWrapper
	extraArgs, err := parser.ParseArgs(args)
	if err == nil {
		return
	}

	if flagErr, ok := err.(*flags.Error); ok {
		switch flagErr.Type {
		case flags.ErrHelp, flags.ErrUnknownFlag, flags.ErrExpectedArgument, flags.ErrInvalidChoice:
			_, found := reflect.TypeOf(common.Commands).FieldByNameFunc(
				func(fieldName string) bool {
					field, _ := reflect.TypeOf(common.Commands).FieldByName(fieldName)
					return parser.Active != nil && parser.Active.Name == field.Tag.Get("command")
				},
			)

			if found && flagErr.Type == flags.ErrUnknownFlag && parser.Active.Name == "set-env" {
				newArgs := []string{}
				for _, arg := range args {
					if arg[0] == '-' {
						newArgs = append(newArgs, fmt.Sprintf("%s%s", v2.WorkAroundPrefix, arg))
					} else {
						newArgs = append(newArgs, arg)
					}
				}
				parse(newArgs)
				return
			}

			if flagErr.Type == flags.ErrUnknownFlag || flagErr.Type == flags.ErrExpectedArgument || flagErr.Type == flags.ErrInvalidChoice {
				fmt.Fprintf(os.Stderr, "Incorrect Usage: %s\n\n", flagErr.Error())
			}

			if found {
				parse([]string{"help", parser.Active.Name})
			} else {
				switch len(extraArgs) {
				case 0:
					parse([]string{"help"})
				case 1:
					if !isOption(extraArgs[0]) || (len(args) > 1 && extraArgs[0] == "-a") {
						parse([]string{"help", extraArgs[0]})
					} else {
						parse([]string{"help"})
					}
				default:
					if isCommand(extraArgs[0]) {
						parse([]string{"help", extraArgs[0]})
					} else {
						parse(extraArgs[1:])
					}
				}
			}

			if flagErr.Type == flags.ErrUnknownFlag || flagErr.Type == flags.ErrExpectedArgument || flagErr.Type == flags.ErrInvalidChoice {
				os.Exit(1)
			}
		case flags.ErrRequired:
			fmt.Fprintf(os.Stderr, "Incorrect Usage: %s\n\n", flagErr.Error())
			parse([]string{"help", args[0]})
			os.Exit(1)
		case flags.ErrMarshal:
			errMessage := strings.Split(flagErr.Message, ":")
			fmt.Fprintf(os.Stderr, "Incorrect Usage: %s\n\n", errMessage[0])
			parse([]string{"help", args[0]})
			os.Exit(1)
		case flags.ErrUnknownCommand:
			cmd.Main(os.Getenv("CF_TRACE"), os.Args)
		case flags.ErrCommandRequired:
			if common.Commands.VerboseOrVersion {
				parse([]string{"version"})
			} else {
				parse([]string{"help"})
			}
		default:
			fmt.Fprintf(os.Stderr, "Unexpected flag error\ntype: %s\nmessage: %s\n", flagErr.Type, flagErr.Error())
		}
	} else if err == ErrFailed {
		os.Exit(1)
	} else if err == ParseErr {
		fmt.Println()
		parse([]string{"help", args[0]})
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}
}

func isCommand(s string) bool {
	_, found := reflect.TypeOf(common.Commands).FieldByNameFunc(
		func(fieldName string) bool {
			field, _ := reflect.TypeOf(common.Commands).FieldByName(fieldName)
			return s == field.Tag.Get("command") || s == field.Tag.Get("alias")
		})

	return found
}

func isOption(s string) bool {
	return strings.HasPrefix(s, "-")
}

func executionWrapper(cmd flags.Commander, args []string) error {
	cfConfig, configErr := configv3.LoadConfig(configv3.FlagOverride{
		Verbose: common.Commands.VerboseOrVersion,
	})
	if configErr != nil {
		if _, ok := configErr.(translatableerror.EmptyConfigError); !ok {
			return configErr
		}
	}

	commandUI, err := ui.NewUI(cfConfig)
	if err != nil {
		return err
	}

	// TODO: when the line in the old code under `cf` which calls
	// configv3.LoadConfig() is finally removed, then we should replace the code
	// path above with the following:
	//
	// var configErrTemplate string
	// if configErr != nil {
	// 	if ce, ok := configErr.(translatableerror.EmptyConfigError); ok {
	// 		configErrTemplate = ce.Error()
	// 	} else {
	// 		return configErr
	// 	}
	// }

	// commandUI, err := ui.NewUI(cfConfig)
	// if err != nil {
	// 	return err
	// }

	// if configErr != nil {
	//   commandUI.DisplayWarning(configErrTemplate, map[string]interface{}{
	//   	"FilePath": configv3.ConfigFilePath(),
	//   })
	// }

	defer func() {
		configWriteErr := configv3.WriteConfig(cfConfig)
		if configWriteErr != nil {
			fmt.Fprintf(os.Stderr, "Error writing config: %s", configWriteErr.Error())
		}
	}()

	if extendedCmd, ok := cmd.(command.ExtendedCommander); ok {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.Level(cfConfig.LogLevel()))

		err = extendedCmd.Setup(cfConfig, commandUI)
		if err != nil {
			return handleError(err, commandUI)
		}
		return handleError(extendedCmd.Execute(args), commandUI)
	}

	return fmt.Errorf("command does not conform to ExtendedCommander")
}

func handleError(err error, commandUI UI) error {
	if err == nil {
		return nil
	}

	commandUI.DisplayError(err)

	if _, ok := err.(DisplayUsage); ok {
		return ParseErr
	}

	return ErrFailed
}
