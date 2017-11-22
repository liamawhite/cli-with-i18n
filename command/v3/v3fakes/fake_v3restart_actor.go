// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/command/v3"
)

type FakeV3RestartActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	StartApplicationStub        func(appGUID string) (v3action.Application, v3action.Warnings, error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		appGUID string
	}
	startApplicationReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	StopApplicationStub        func(appGUID string) (v3action.Warnings, error)
	stopApplicationMutex       sync.RWMutex
	stopApplicationArgsForCall []struct {
		appGUID string
	}
	stopApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	stopApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3RestartActor) CloudControllerAPIVersion() string {
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

func (fake *FakeV3RestartActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3RestartActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3RestartActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeV3RestartActor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV3RestartActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3RestartActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3RestartActor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3RestartActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3RestartActor) StartApplication(appGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("StartApplication", []interface{}{appGUID})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.startApplicationReturns.result1, fake.startApplicationReturns.result2, fake.startApplicationReturns.result3
}

func (fake *FakeV3RestartActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeV3RestartActor) StartApplicationArgsForCall(i int) string {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return fake.startApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3RestartActor) StartApplicationReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3RestartActor) StartApplicationReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3RestartActor) StopApplication(appGUID string) (v3action.Warnings, error) {
	fake.stopApplicationMutex.Lock()
	ret, specificReturn := fake.stopApplicationReturnsOnCall[len(fake.stopApplicationArgsForCall)]
	fake.stopApplicationArgsForCall = append(fake.stopApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("StopApplication", []interface{}{appGUID})
	fake.stopApplicationMutex.Unlock()
	if fake.StopApplicationStub != nil {
		return fake.StopApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stopApplicationReturns.result1, fake.stopApplicationReturns.result2
}

func (fake *FakeV3RestartActor) StopApplicationCallCount() int {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	return len(fake.stopApplicationArgsForCall)
}

func (fake *FakeV3RestartActor) StopApplicationArgsForCall(i int) string {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	return fake.stopApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3RestartActor) StopApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.StopApplicationStub = nil
	fake.stopApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartActor) StopApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.StopApplicationStub = nil
	if fake.stopApplicationReturnsOnCall == nil {
		fake.stopApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.stopApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3RestartActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3RestartActor = new(FakeV3RestartActor)
