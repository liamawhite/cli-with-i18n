// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/command/v3"
)

type FakeV3DeleteActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	DeleteApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v3action.Warnings, error)
	deleteApplicationByNameAndSpaceMutex       sync.RWMutex
	deleteApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	deleteApplicationByNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	deleteApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3DeleteActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3DeleteActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3DeleteActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3DeleteActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3DeleteActor) DeleteApplicationByNameAndSpace(name string, spaceGUID string) (v3action.Warnings, error) {
	fake.deleteApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.deleteApplicationByNameAndSpaceReturnsOnCall[len(fake.deleteApplicationByNameAndSpaceArgsForCall)]
	fake.deleteApplicationByNameAndSpaceArgsForCall = append(fake.deleteApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("DeleteApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.deleteApplicationByNameAndSpaceMutex.Unlock()
	if fake.DeleteApplicationByNameAndSpaceStub != nil {
		return fake.DeleteApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteApplicationByNameAndSpaceReturns.result1, fake.deleteApplicationByNameAndSpaceReturns.result2
}

func (fake *FakeV3DeleteActor) DeleteApplicationByNameAndSpaceCallCount() int {
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.deleteApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3DeleteActor) DeleteApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	return fake.deleteApplicationByNameAndSpaceArgsForCall[i].name, fake.deleteApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3DeleteActor) DeleteApplicationByNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.DeleteApplicationByNameAndSpaceStub = nil
	fake.deleteApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3DeleteActor) DeleteApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.DeleteApplicationByNameAndSpaceStub = nil
	if fake.deleteApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.deleteApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.deleteApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3DeleteActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3DeleteActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3DeleteActor = new(FakeV3DeleteActor)
