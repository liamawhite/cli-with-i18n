// Code generated by counterfeiter. DO NOT EDIT.
package commonfakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/command/common"
)

type FakeHelpActor struct {
	CommandInfoByNameStub        func(interface{}, string) (sharedaction.CommandInfo, error)
	commandInfoByNameMutex       sync.RWMutex
	commandInfoByNameArgsForCall []struct {
		arg1 interface{}
		arg2 string
	}
	commandInfoByNameReturns struct {
		result1 sharedaction.CommandInfo
		result2 error
	}
	commandInfoByNameReturnsOnCall map[int]struct {
		result1 sharedaction.CommandInfo
		result2 error
	}
	CommandInfosStub        func(interface{}) map[string]sharedaction.CommandInfo
	commandInfosMutex       sync.RWMutex
	commandInfosArgsForCall []struct {
		arg1 interface{}
	}
	commandInfosReturns struct {
		result1 map[string]sharedaction.CommandInfo
	}
	commandInfosReturnsOnCall map[int]struct {
		result1 map[string]sharedaction.CommandInfo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHelpActor) CommandInfoByName(arg1 interface{}, arg2 string) (sharedaction.CommandInfo, error) {
	fake.commandInfoByNameMutex.Lock()
	ret, specificReturn := fake.commandInfoByNameReturnsOnCall[len(fake.commandInfoByNameArgsForCall)]
	fake.commandInfoByNameArgsForCall = append(fake.commandInfoByNameArgsForCall, struct {
		arg1 interface{}
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CommandInfoByName", []interface{}{arg1, arg2})
	fake.commandInfoByNameMutex.Unlock()
	if fake.CommandInfoByNameStub != nil {
		return fake.CommandInfoByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.commandInfoByNameReturns.result1, fake.commandInfoByNameReturns.result2
}

func (fake *FakeHelpActor) CommandInfoByNameCallCount() int {
	fake.commandInfoByNameMutex.RLock()
	defer fake.commandInfoByNameMutex.RUnlock()
	return len(fake.commandInfoByNameArgsForCall)
}

func (fake *FakeHelpActor) CommandInfoByNameArgsForCall(i int) (interface{}, string) {
	fake.commandInfoByNameMutex.RLock()
	defer fake.commandInfoByNameMutex.RUnlock()
	return fake.commandInfoByNameArgsForCall[i].arg1, fake.commandInfoByNameArgsForCall[i].arg2
}

func (fake *FakeHelpActor) CommandInfoByNameReturns(result1 sharedaction.CommandInfo, result2 error) {
	fake.CommandInfoByNameStub = nil
	fake.commandInfoByNameReturns = struct {
		result1 sharedaction.CommandInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeHelpActor) CommandInfoByNameReturnsOnCall(i int, result1 sharedaction.CommandInfo, result2 error) {
	fake.CommandInfoByNameStub = nil
	if fake.commandInfoByNameReturnsOnCall == nil {
		fake.commandInfoByNameReturnsOnCall = make(map[int]struct {
			result1 sharedaction.CommandInfo
			result2 error
		})
	}
	fake.commandInfoByNameReturnsOnCall[i] = struct {
		result1 sharedaction.CommandInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeHelpActor) CommandInfos(arg1 interface{}) map[string]sharedaction.CommandInfo {
	fake.commandInfosMutex.Lock()
	ret, specificReturn := fake.commandInfosReturnsOnCall[len(fake.commandInfosArgsForCall)]
	fake.commandInfosArgsForCall = append(fake.commandInfosArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("CommandInfos", []interface{}{arg1})
	fake.commandInfosMutex.Unlock()
	if fake.CommandInfosStub != nil {
		return fake.CommandInfosStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.commandInfosReturns.result1
}

func (fake *FakeHelpActor) CommandInfosCallCount() int {
	fake.commandInfosMutex.RLock()
	defer fake.commandInfosMutex.RUnlock()
	return len(fake.commandInfosArgsForCall)
}

func (fake *FakeHelpActor) CommandInfosArgsForCall(i int) interface{} {
	fake.commandInfosMutex.RLock()
	defer fake.commandInfosMutex.RUnlock()
	return fake.commandInfosArgsForCall[i].arg1
}

func (fake *FakeHelpActor) CommandInfosReturns(result1 map[string]sharedaction.CommandInfo) {
	fake.CommandInfosStub = nil
	fake.commandInfosReturns = struct {
		result1 map[string]sharedaction.CommandInfo
	}{result1}
}

func (fake *FakeHelpActor) CommandInfosReturnsOnCall(i int, result1 map[string]sharedaction.CommandInfo) {
	fake.CommandInfosStub = nil
	if fake.commandInfosReturnsOnCall == nil {
		fake.commandInfosReturnsOnCall = make(map[int]struct {
			result1 map[string]sharedaction.CommandInfo
		})
	}
	fake.commandInfosReturnsOnCall[i] = struct {
		result1 map[string]sharedaction.CommandInfo
	}{result1}
}

func (fake *FakeHelpActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandInfoByNameMutex.RLock()
	defer fake.commandInfoByNameMutex.RUnlock()
	fake.commandInfosMutex.RLock()
	defer fake.commandInfosMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHelpActor) recordInvocation(key string, args []interface{}) {
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

var _ common.HelpActor = new(FakeHelpActor)
