// Code generated by counterfeiter. DO NOT EDIT.
package cfnetworkingactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cfnetworking-cli-api/cfnetworking/cfnetv1"
	"github.com/liamawhite/cli-with-i18n/actor/cfnetworkingaction"
)

type FakeNetworkingClient struct {
	CreatePoliciesStub        func(policies []cfnetv1.Policy) error
	createPoliciesMutex       sync.RWMutex
	createPoliciesArgsForCall []struct {
		policies []cfnetv1.Policy
	}
	createPoliciesReturns struct {
		result1 error
	}
	createPoliciesReturnsOnCall map[int]struct {
		result1 error
	}
	ListPoliciesStub        func(appNames ...string) ([]cfnetv1.Policy, error)
	listPoliciesMutex       sync.RWMutex
	listPoliciesArgsForCall []struct {
		appNames []string
	}
	listPoliciesReturns struct {
		result1 []cfnetv1.Policy
		result2 error
	}
	listPoliciesReturnsOnCall map[int]struct {
		result1 []cfnetv1.Policy
		result2 error
	}
	RemovePoliciesStub        func(policies []cfnetv1.Policy) error
	removePoliciesMutex       sync.RWMutex
	removePoliciesArgsForCall []struct {
		policies []cfnetv1.Policy
	}
	removePoliciesReturns struct {
		result1 error
	}
	removePoliciesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworkingClient) CreatePolicies(policies []cfnetv1.Policy) error {
	var policiesCopy []cfnetv1.Policy
	if policies != nil {
		policiesCopy = make([]cfnetv1.Policy, len(policies))
		copy(policiesCopy, policies)
	}
	fake.createPoliciesMutex.Lock()
	ret, specificReturn := fake.createPoliciesReturnsOnCall[len(fake.createPoliciesArgsForCall)]
	fake.createPoliciesArgsForCall = append(fake.createPoliciesArgsForCall, struct {
		policies []cfnetv1.Policy
	}{policiesCopy})
	fake.recordInvocation("CreatePolicies", []interface{}{policiesCopy})
	fake.createPoliciesMutex.Unlock()
	if fake.CreatePoliciesStub != nil {
		return fake.CreatePoliciesStub(policies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createPoliciesReturns.result1
}

func (fake *FakeNetworkingClient) CreatePoliciesCallCount() int {
	fake.createPoliciesMutex.RLock()
	defer fake.createPoliciesMutex.RUnlock()
	return len(fake.createPoliciesArgsForCall)
}

func (fake *FakeNetworkingClient) CreatePoliciesArgsForCall(i int) []cfnetv1.Policy {
	fake.createPoliciesMutex.RLock()
	defer fake.createPoliciesMutex.RUnlock()
	return fake.createPoliciesArgsForCall[i].policies
}

func (fake *FakeNetworkingClient) CreatePoliciesReturns(result1 error) {
	fake.CreatePoliciesStub = nil
	fake.createPoliciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworkingClient) CreatePoliciesReturnsOnCall(i int, result1 error) {
	fake.CreatePoliciesStub = nil
	if fake.createPoliciesReturnsOnCall == nil {
		fake.createPoliciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createPoliciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworkingClient) ListPolicies(appNames ...string) ([]cfnetv1.Policy, error) {
	fake.listPoliciesMutex.Lock()
	ret, specificReturn := fake.listPoliciesReturnsOnCall[len(fake.listPoliciesArgsForCall)]
	fake.listPoliciesArgsForCall = append(fake.listPoliciesArgsForCall, struct {
		appNames []string
	}{appNames})
	fake.recordInvocation("ListPolicies", []interface{}{appNames})
	fake.listPoliciesMutex.Unlock()
	if fake.ListPoliciesStub != nil {
		return fake.ListPoliciesStub(appNames...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listPoliciesReturns.result1, fake.listPoliciesReturns.result2
}

func (fake *FakeNetworkingClient) ListPoliciesCallCount() int {
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	return len(fake.listPoliciesArgsForCall)
}

func (fake *FakeNetworkingClient) ListPoliciesArgsForCall(i int) []string {
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	return fake.listPoliciesArgsForCall[i].appNames
}

func (fake *FakeNetworkingClient) ListPoliciesReturns(result1 []cfnetv1.Policy, result2 error) {
	fake.ListPoliciesStub = nil
	fake.listPoliciesReturns = struct {
		result1 []cfnetv1.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworkingClient) ListPoliciesReturnsOnCall(i int, result1 []cfnetv1.Policy, result2 error) {
	fake.ListPoliciesStub = nil
	if fake.listPoliciesReturnsOnCall == nil {
		fake.listPoliciesReturnsOnCall = make(map[int]struct {
			result1 []cfnetv1.Policy
			result2 error
		})
	}
	fake.listPoliciesReturnsOnCall[i] = struct {
		result1 []cfnetv1.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworkingClient) RemovePolicies(policies []cfnetv1.Policy) error {
	var policiesCopy []cfnetv1.Policy
	if policies != nil {
		policiesCopy = make([]cfnetv1.Policy, len(policies))
		copy(policiesCopy, policies)
	}
	fake.removePoliciesMutex.Lock()
	ret, specificReturn := fake.removePoliciesReturnsOnCall[len(fake.removePoliciesArgsForCall)]
	fake.removePoliciesArgsForCall = append(fake.removePoliciesArgsForCall, struct {
		policies []cfnetv1.Policy
	}{policiesCopy})
	fake.recordInvocation("RemovePolicies", []interface{}{policiesCopy})
	fake.removePoliciesMutex.Unlock()
	if fake.RemovePoliciesStub != nil {
		return fake.RemovePoliciesStub(policies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removePoliciesReturns.result1
}

func (fake *FakeNetworkingClient) RemovePoliciesCallCount() int {
	fake.removePoliciesMutex.RLock()
	defer fake.removePoliciesMutex.RUnlock()
	return len(fake.removePoliciesArgsForCall)
}

func (fake *FakeNetworkingClient) RemovePoliciesArgsForCall(i int) []cfnetv1.Policy {
	fake.removePoliciesMutex.RLock()
	defer fake.removePoliciesMutex.RUnlock()
	return fake.removePoliciesArgsForCall[i].policies
}

func (fake *FakeNetworkingClient) RemovePoliciesReturns(result1 error) {
	fake.RemovePoliciesStub = nil
	fake.removePoliciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworkingClient) RemovePoliciesReturnsOnCall(i int, result1 error) {
	fake.RemovePoliciesStub = nil
	if fake.removePoliciesReturnsOnCall == nil {
		fake.removePoliciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removePoliciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworkingClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPoliciesMutex.RLock()
	defer fake.createPoliciesMutex.RUnlock()
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	fake.removePoliciesMutex.RLock()
	defer fake.removePoliciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworkingClient) recordInvocation(key string, args []interface{}) {
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

var _ cfnetworkingaction.NetworkingClient = new(FakeNetworkingClient)
