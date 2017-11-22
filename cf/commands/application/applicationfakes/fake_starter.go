// This file was generated by counterfeiter
package applicationfakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/commands/application"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
)

type FakeStarter struct {
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 commandregistry.CommandMetadata
	}
	SetDependencyStub        func(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) ([]requirements.Requirement, error)
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	ExecuteStub        func(context flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	SetStartTimeoutInSecondsStub        func(timeout int)
	setStartTimeoutInSecondsMutex       sync.RWMutex
	setStartTimeoutInSecondsArgsForCall []struct {
		timeout int
	}
	ApplicationStartStub        func(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error)
	applicationStartMutex       sync.RWMutex
	applicationStartArgsForCall []struct {
		app       models.Application
		orgName   string
		spaceName string
	}
	applicationStartReturns struct {
		result1 models.Application
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStarter) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.recordInvocation("MetaData", []interface{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeStarter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeStarter) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeStarter) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.recordInvocation("SetDependency", []interface{}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeStarter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeStarter) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeStarter) SetDependencyReturns(result1 commandregistry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeStarter) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) ([]requirements.Requirement, error) {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.recordInvocation("Requirements", []interface{}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1, fake.requirementsReturns.result2
	}
}

func (fake *FakeStarter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeStarter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeStarter) RequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) Execute(context flags.FlagContext) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.recordInvocation("Execute", []interface{}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(context)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeStarter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeStarter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeStarter) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStarter) SetStartTimeoutInSeconds(timeout int) {
	fake.setStartTimeoutInSecondsMutex.Lock()
	fake.setStartTimeoutInSecondsArgsForCall = append(fake.setStartTimeoutInSecondsArgsForCall, struct {
		timeout int
	}{timeout})
	fake.recordInvocation("SetStartTimeoutInSeconds", []interface{}{timeout})
	fake.setStartTimeoutInSecondsMutex.Unlock()
	if fake.SetStartTimeoutInSecondsStub != nil {
		fake.SetStartTimeoutInSecondsStub(timeout)
	}
}

func (fake *FakeStarter) SetStartTimeoutInSecondsCallCount() int {
	fake.setStartTimeoutInSecondsMutex.RLock()
	defer fake.setStartTimeoutInSecondsMutex.RUnlock()
	return len(fake.setStartTimeoutInSecondsArgsForCall)
}

func (fake *FakeStarter) SetStartTimeoutInSecondsArgsForCall(i int) int {
	fake.setStartTimeoutInSecondsMutex.RLock()
	defer fake.setStartTimeoutInSecondsMutex.RUnlock()
	return fake.setStartTimeoutInSecondsArgsForCall[i].timeout
}

func (fake *FakeStarter) ApplicationStart(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error) {
	fake.applicationStartMutex.Lock()
	fake.applicationStartArgsForCall = append(fake.applicationStartArgsForCall, struct {
		app       models.Application
		orgName   string
		spaceName string
	}{app, orgName, spaceName})
	fake.recordInvocation("ApplicationStart", []interface{}{app, orgName, spaceName})
	fake.applicationStartMutex.Unlock()
	if fake.ApplicationStartStub != nil {
		return fake.ApplicationStartStub(app, orgName, spaceName)
	} else {
		return fake.applicationStartReturns.result1, fake.applicationStartReturns.result2
	}
}

func (fake *FakeStarter) ApplicationStartCallCount() int {
	fake.applicationStartMutex.RLock()
	defer fake.applicationStartMutex.RUnlock()
	return len(fake.applicationStartArgsForCall)
}

func (fake *FakeStarter) ApplicationStartArgsForCall(i int) (models.Application, string, string) {
	fake.applicationStartMutex.RLock()
	defer fake.applicationStartMutex.RUnlock()
	return fake.applicationStartArgsForCall[i].app, fake.applicationStartArgsForCall[i].orgName, fake.applicationStartArgsForCall[i].spaceName
}

func (fake *FakeStarter) ApplicationStartReturns(result1 models.Application, result2 error) {
	fake.ApplicationStartStub = nil
	fake.applicationStartReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.setStartTimeoutInSecondsMutex.RLock()
	defer fake.setStartTimeoutInSecondsMutex.RUnlock()
	fake.applicationStartMutex.RLock()
	defer fake.applicationStartMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStarter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ application.Starter = new(FakeStarter)
