// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

type FakeServicePlanRepository struct {
	SearchStub        func(searchParameters map[string]string) ([]models.ServicePlanFields, error)
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		searchParameters map[string]string
	}
	searchReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	UpdateStub        func(models.ServicePlanFields, string, bool) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 models.ServicePlanFields
		arg2 string
		arg3 bool
	}
	updateReturns struct {
		result1 error
	}
	ListPlansFromManyServicesStub        func(serviceGUIDs []string) ([]models.ServicePlanFields, error)
	listPlansFromManyServicesMutex       sync.RWMutex
	listPlansFromManyServicesArgsForCall []struct {
		serviceGUIDs []string
	}
	listPlansFromManyServicesReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServicePlanRepository) Search(searchParameters map[string]string) ([]models.ServicePlanFields, error) {
	fake.searchMutex.Lock()
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		searchParameters map[string]string
	}{searchParameters})
	fake.recordInvocation("Search", []interface{}{searchParameters})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(searchParameters)
	} else {
		return fake.searchReturns.result1, fake.searchReturns.result2
	}
}

func (fake *FakeServicePlanRepository) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeServicePlanRepository) SearchArgsForCall(i int) map[string]string {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return fake.searchArgsForCall[i].searchParameters
}

func (fake *FakeServicePlanRepository) SearchReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanRepository) Update(arg1 models.ServicePlanFields, arg2 string, arg3 bool) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 models.ServicePlanFields
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeServicePlanRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServicePlanRepository) UpdateArgsForCall(i int) (models.ServicePlanFields, string, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2, fake.updateArgsForCall[i].arg3
}

func (fake *FakeServicePlanRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServicePlanRepository) ListPlansFromManyServices(serviceGUIDs []string) ([]models.ServicePlanFields, error) {
	var serviceGUIDsCopy []string
	if serviceGUIDs != nil {
		serviceGUIDsCopy = make([]string, len(serviceGUIDs))
		copy(serviceGUIDsCopy, serviceGUIDs)
	}
	fake.listPlansFromManyServicesMutex.Lock()
	fake.listPlansFromManyServicesArgsForCall = append(fake.listPlansFromManyServicesArgsForCall, struct {
		serviceGUIDs []string
	}{serviceGUIDsCopy})
	fake.recordInvocation("ListPlansFromManyServices", []interface{}{serviceGUIDsCopy})
	fake.listPlansFromManyServicesMutex.Unlock()
	if fake.ListPlansFromManyServicesStub != nil {
		return fake.ListPlansFromManyServicesStub(serviceGUIDs)
	} else {
		return fake.listPlansFromManyServicesReturns.result1, fake.listPlansFromManyServicesReturns.result2
	}
}

func (fake *FakeServicePlanRepository) ListPlansFromManyServicesCallCount() int {
	fake.listPlansFromManyServicesMutex.RLock()
	defer fake.listPlansFromManyServicesMutex.RUnlock()
	return len(fake.listPlansFromManyServicesArgsForCall)
}

func (fake *FakeServicePlanRepository) ListPlansFromManyServicesArgsForCall(i int) []string {
	fake.listPlansFromManyServicesMutex.RLock()
	defer fake.listPlansFromManyServicesMutex.RUnlock()
	return fake.listPlansFromManyServicesArgsForCall[i].serviceGUIDs
}

func (fake *FakeServicePlanRepository) ListPlansFromManyServicesReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.ListPlansFromManyServicesStub = nil
	fake.listPlansFromManyServicesReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.listPlansFromManyServicesMutex.RLock()
	defer fake.listPlansFromManyServicesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServicePlanRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServicePlanRepository = new(FakeServicePlanRepository)
