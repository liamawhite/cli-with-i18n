// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

type FakeServiceAuthTokenRepository struct {
	FindAllStub        func() (authTokens []models.ServiceAuthTokenFields, apiErr error)
	findAllMutex       sync.RWMutex
	findAllArgsForCall []struct{}
	findAllReturns     struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}
	FindByLabelAndProviderStub        func(label, provider string) (authToken models.ServiceAuthTokenFields, apiErr error)
	findByLabelAndProviderMutex       sync.RWMutex
	findByLabelAndProviderArgsForCall []struct {
		label    string
		provider string
	}
	findByLabelAndProviderReturns struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}
	CreateStub        func(authToken models.ServiceAuthTokenFields) (apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		authToken models.ServiceAuthTokenFields
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(authToken models.ServiceAuthTokenFields) (apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		authToken models.ServiceAuthTokenFields
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(authToken models.ServiceAuthTokenFields) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		authToken models.ServiceAuthTokenFields
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAuthTokenRepository) FindAll() (authTokens []models.ServiceAuthTokenFields, apiErr error) {
	fake.findAllMutex.Lock()
	fake.findAllArgsForCall = append(fake.findAllArgsForCall, struct{}{})
	fake.recordInvocation("FindAll", []interface{}{})
	fake.findAllMutex.Unlock()
	if fake.FindAllStub != nil {
		return fake.FindAllStub()
	} else {
		return fake.findAllReturns.result1, fake.findAllReturns.result2
	}
}

func (fake *FakeServiceAuthTokenRepository) FindAllCallCount() int {
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	return len(fake.findAllArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) FindAllReturns(result1 []models.ServiceAuthTokenFields, result2 error) {
	fake.FindAllStub = nil
	fake.findAllReturns = struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProvider(label string, provider string) (authToken models.ServiceAuthTokenFields, apiErr error) {
	fake.findByLabelAndProviderMutex.Lock()
	fake.findByLabelAndProviderArgsForCall = append(fake.findByLabelAndProviderArgsForCall, struct {
		label    string
		provider string
	}{label, provider})
	fake.recordInvocation("FindByLabelAndProvider", []interface{}{label, provider})
	fake.findByLabelAndProviderMutex.Unlock()
	if fake.FindByLabelAndProviderStub != nil {
		return fake.FindByLabelAndProviderStub(label, provider)
	} else {
		return fake.findByLabelAndProviderReturns.result1, fake.findByLabelAndProviderReturns.result2
	}
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderCallCount() int {
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	return len(fake.findByLabelAndProviderArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderArgsForCall(i int) (string, string) {
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	return fake.findByLabelAndProviderArgsForCall[i].label, fake.findByLabelAndProviderArgsForCall[i].provider
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderReturns(result1 models.ServiceAuthTokenFields, result2 error) {
	fake.FindByLabelAndProviderStub = nil
	fake.findByLabelAndProviderReturns = struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) Create(authToken models.ServiceAuthTokenFields) (apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		authToken models.ServiceAuthTokenFields
	}{authToken})
	fake.recordInvocation("Create", []interface{}{authToken})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(authToken)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeServiceAuthTokenRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) CreateArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].authToken
}

func (fake *FakeServiceAuthTokenRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) Update(authToken models.ServiceAuthTokenFields) (apiErr error) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		authToken models.ServiceAuthTokenFields
	}{authToken})
	fake.recordInvocation("Update", []interface{}{authToken})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(authToken)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeServiceAuthTokenRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) UpdateArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].authToken
}

func (fake *FakeServiceAuthTokenRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) Delete(authToken models.ServiceAuthTokenFields) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		authToken models.ServiceAuthTokenFields
	}{authToken})
	fake.recordInvocation("Delete", []interface{}{authToken})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(authToken)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeServiceAuthTokenRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) DeleteArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].authToken
}

func (fake *FakeServiceAuthTokenRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServiceAuthTokenRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServiceAuthTokenRepository = new(FakeServiceAuthTokenRepository)
