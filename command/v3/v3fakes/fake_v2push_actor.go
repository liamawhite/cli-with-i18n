// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/actor/pushaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/command/v3"
)

type FakeV2PushActor struct {
	CreateAndBindApplicationRoutesStub        func(orgGUID string, spaceGUID string, app v2action.Application) (pushaction.Warnings, error)
	createAndBindApplicationRoutesMutex       sync.RWMutex
	createAndBindApplicationRoutesArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		app       v2action.Application
	}
	createAndBindApplicationRoutesReturns struct {
		result1 pushaction.Warnings
		result2 error
	}
	createAndBindApplicationRoutesReturnsOnCall map[int]struct {
		result1 pushaction.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2PushActor) CreateAndBindApplicationRoutes(orgGUID string, spaceGUID string, app v2action.Application) (pushaction.Warnings, error) {
	fake.createAndBindApplicationRoutesMutex.Lock()
	ret, specificReturn := fake.createAndBindApplicationRoutesReturnsOnCall[len(fake.createAndBindApplicationRoutesArgsForCall)]
	fake.createAndBindApplicationRoutesArgsForCall = append(fake.createAndBindApplicationRoutesArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		app       v2action.Application
	}{orgGUID, spaceGUID, app})
	fake.recordInvocation("CreateAndBindApplicationRoutes", []interface{}{orgGUID, spaceGUID, app})
	fake.createAndBindApplicationRoutesMutex.Unlock()
	if fake.CreateAndBindApplicationRoutesStub != nil {
		return fake.CreateAndBindApplicationRoutesStub(orgGUID, spaceGUID, app)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAndBindApplicationRoutesReturns.result1, fake.createAndBindApplicationRoutesReturns.result2
}

func (fake *FakeV2PushActor) CreateAndBindApplicationRoutesCallCount() int {
	fake.createAndBindApplicationRoutesMutex.RLock()
	defer fake.createAndBindApplicationRoutesMutex.RUnlock()
	return len(fake.createAndBindApplicationRoutesArgsForCall)
}

func (fake *FakeV2PushActor) CreateAndBindApplicationRoutesArgsForCall(i int) (string, string, v2action.Application) {
	fake.createAndBindApplicationRoutesMutex.RLock()
	defer fake.createAndBindApplicationRoutesMutex.RUnlock()
	return fake.createAndBindApplicationRoutesArgsForCall[i].orgGUID, fake.createAndBindApplicationRoutesArgsForCall[i].spaceGUID, fake.createAndBindApplicationRoutesArgsForCall[i].app
}

func (fake *FakeV2PushActor) CreateAndBindApplicationRoutesReturns(result1 pushaction.Warnings, result2 error) {
	fake.CreateAndBindApplicationRoutesStub = nil
	fake.createAndBindApplicationRoutesReturns = struct {
		result1 pushaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) CreateAndBindApplicationRoutesReturnsOnCall(i int, result1 pushaction.Warnings, result2 error) {
	fake.CreateAndBindApplicationRoutesStub = nil
	if fake.createAndBindApplicationRoutesReturnsOnCall == nil {
		fake.createAndBindApplicationRoutesReturnsOnCall = make(map[int]struct {
			result1 pushaction.Warnings
			result2 error
		})
	}
	fake.createAndBindApplicationRoutesReturnsOnCall[i] = struct {
		result1 pushaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAndBindApplicationRoutesMutex.RLock()
	defer fake.createAndBindApplicationRoutesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV2PushActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V2PushActor = new(FakeV2PushActor)
