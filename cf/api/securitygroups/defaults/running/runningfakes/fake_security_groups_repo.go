// This file was generated by counterfeiter
package runningfakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/api/securitygroups/defaults/running"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

type FakeSecurityGroupsRepo struct {
	BindToRunningSetStub        func(string) error
	bindToRunningSetMutex       sync.RWMutex
	bindToRunningSetArgsForCall []struct {
		arg1 string
	}
	bindToRunningSetReturns struct {
		result1 error
	}
	ListStub        func() ([]models.SecurityGroupFields, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []models.SecurityGroupFields
		result2 error
	}
	UnbindFromRunningSetStub        func(string) error
	unbindFromRunningSetMutex       sync.RWMutex
	unbindFromRunningSetArgsForCall []struct {
		arg1 string
	}
	unbindFromRunningSetReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityGroupsRepo) BindToRunningSet(arg1 string) error {
	fake.bindToRunningSetMutex.Lock()
	fake.bindToRunningSetArgsForCall = append(fake.bindToRunningSetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("BindToRunningSet", []interface{}{arg1})
	fake.bindToRunningSetMutex.Unlock()
	if fake.BindToRunningSetStub != nil {
		return fake.BindToRunningSetStub(arg1)
	} else {
		return fake.bindToRunningSetReturns.result1
	}
}

func (fake *FakeSecurityGroupsRepo) BindToRunningSetCallCount() int {
	fake.bindToRunningSetMutex.RLock()
	defer fake.bindToRunningSetMutex.RUnlock()
	return len(fake.bindToRunningSetArgsForCall)
}

func (fake *FakeSecurityGroupsRepo) BindToRunningSetArgsForCall(i int) string {
	fake.bindToRunningSetMutex.RLock()
	defer fake.bindToRunningSetMutex.RUnlock()
	return fake.bindToRunningSetArgsForCall[i].arg1
}

func (fake *FakeSecurityGroupsRepo) BindToRunningSetReturns(result1 error) {
	fake.BindToRunningSetStub = nil
	fake.bindToRunningSetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupsRepo) List() ([]models.SecurityGroupFields, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeSecurityGroupsRepo) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeSecurityGroupsRepo) ListReturns(result1 []models.SecurityGroupFields, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []models.SecurityGroupFields
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupsRepo) UnbindFromRunningSet(arg1 string) error {
	fake.unbindFromRunningSetMutex.Lock()
	fake.unbindFromRunningSetArgsForCall = append(fake.unbindFromRunningSetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnbindFromRunningSet", []interface{}{arg1})
	fake.unbindFromRunningSetMutex.Unlock()
	if fake.UnbindFromRunningSetStub != nil {
		return fake.UnbindFromRunningSetStub(arg1)
	} else {
		return fake.unbindFromRunningSetReturns.result1
	}
}

func (fake *FakeSecurityGroupsRepo) UnbindFromRunningSetCallCount() int {
	fake.unbindFromRunningSetMutex.RLock()
	defer fake.unbindFromRunningSetMutex.RUnlock()
	return len(fake.unbindFromRunningSetArgsForCall)
}

func (fake *FakeSecurityGroupsRepo) UnbindFromRunningSetArgsForCall(i int) string {
	fake.unbindFromRunningSetMutex.RLock()
	defer fake.unbindFromRunningSetMutex.RUnlock()
	return fake.unbindFromRunningSetArgsForCall[i].arg1
}

func (fake *FakeSecurityGroupsRepo) UnbindFromRunningSetReturns(result1 error) {
	fake.UnbindFromRunningSetStub = nil
	fake.unbindFromRunningSetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupsRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindToRunningSetMutex.RLock()
	defer fake.bindToRunningSetMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.unbindFromRunningSetMutex.RLock()
	defer fake.unbindFromRunningSetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSecurityGroupsRepo) recordInvocation(key string, args []interface{}) {
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

var _ running.SecurityGroupsRepo = new(FakeSecurityGroupsRepo)
