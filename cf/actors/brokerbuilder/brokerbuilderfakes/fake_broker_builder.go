// This file was generated by counterfeiter
package brokerbuilderfakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/actors/brokerbuilder"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

type FakeBrokerBuilder struct {
	AttachBrokersToServicesStub        func([]models.ServiceOffering) ([]models.ServiceBroker, error)
	attachBrokersToServicesMutex       sync.RWMutex
	attachBrokersToServicesArgsForCall []struct {
		arg1 []models.ServiceOffering
	}
	attachBrokersToServicesReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	AttachSpecificBrokerToServicesStub        func(string, []models.ServiceOffering) (models.ServiceBroker, error)
	attachSpecificBrokerToServicesMutex       sync.RWMutex
	attachSpecificBrokerToServicesArgsForCall []struct {
		arg1 string
		arg2 []models.ServiceOffering
	}
	attachSpecificBrokerToServicesReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	GetAllServiceBrokersStub        func() ([]models.ServiceBroker, error)
	getAllServiceBrokersMutex       sync.RWMutex
	getAllServiceBrokersArgsForCall []struct{}
	getAllServiceBrokersReturns     struct {
		result1 []models.ServiceBroker
		result2 error
	}
	GetBrokerWithAllServicesStub        func(brokerName string) (models.ServiceBroker, error)
	getBrokerWithAllServicesMutex       sync.RWMutex
	getBrokerWithAllServicesArgsForCall []struct {
		brokerName string
	}
	getBrokerWithAllServicesReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	GetBrokerWithSpecifiedServiceStub        func(serviceName string) (models.ServiceBroker, error)
	getBrokerWithSpecifiedServiceMutex       sync.RWMutex
	getBrokerWithSpecifiedServiceArgsForCall []struct {
		serviceName string
	}
	getBrokerWithSpecifiedServiceReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBrokerBuilder) AttachBrokersToServices(arg1 []models.ServiceOffering) ([]models.ServiceBroker, error) {
	var arg1Copy []models.ServiceOffering
	if arg1 != nil {
		arg1Copy = make([]models.ServiceOffering, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.attachBrokersToServicesMutex.Lock()
	fake.attachBrokersToServicesArgsForCall = append(fake.attachBrokersToServicesArgsForCall, struct {
		arg1 []models.ServiceOffering
	}{arg1Copy})
	fake.recordInvocation("AttachBrokersToServices", []interface{}{arg1Copy})
	fake.attachBrokersToServicesMutex.Unlock()
	if fake.AttachBrokersToServicesStub != nil {
		return fake.AttachBrokersToServicesStub(arg1)
	} else {
		return fake.attachBrokersToServicesReturns.result1, fake.attachBrokersToServicesReturns.result2
	}
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesCallCount() int {
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	return len(fake.attachBrokersToServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesArgsForCall(i int) []models.ServiceOffering {
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	return fake.attachBrokersToServicesArgsForCall[i].arg1
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesReturns(result1 []models.ServiceBroker, result2 error) {
	fake.AttachBrokersToServicesStub = nil
	fake.attachBrokersToServicesReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServices(arg1 string, arg2 []models.ServiceOffering) (models.ServiceBroker, error) {
	var arg2Copy []models.ServiceOffering
	if arg2 != nil {
		arg2Copy = make([]models.ServiceOffering, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.attachSpecificBrokerToServicesMutex.Lock()
	fake.attachSpecificBrokerToServicesArgsForCall = append(fake.attachSpecificBrokerToServicesArgsForCall, struct {
		arg1 string
		arg2 []models.ServiceOffering
	}{arg1, arg2Copy})
	fake.recordInvocation("AttachSpecificBrokerToServices", []interface{}{arg1, arg2Copy})
	fake.attachSpecificBrokerToServicesMutex.Unlock()
	if fake.AttachSpecificBrokerToServicesStub != nil {
		return fake.AttachSpecificBrokerToServicesStub(arg1, arg2)
	} else {
		return fake.attachSpecificBrokerToServicesReturns.result1, fake.attachSpecificBrokerToServicesReturns.result2
	}
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesCallCount() int {
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	return len(fake.attachSpecificBrokerToServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesArgsForCall(i int) (string, []models.ServiceOffering) {
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	return fake.attachSpecificBrokerToServicesArgsForCall[i].arg1, fake.attachSpecificBrokerToServicesArgsForCall[i].arg2
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesReturns(result1 models.ServiceBroker, result2 error) {
	fake.AttachSpecificBrokerToServicesStub = nil
	fake.attachSpecificBrokerToServicesReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokers() ([]models.ServiceBroker, error) {
	fake.getAllServiceBrokersMutex.Lock()
	fake.getAllServiceBrokersArgsForCall = append(fake.getAllServiceBrokersArgsForCall, struct{}{})
	fake.recordInvocation("GetAllServiceBrokers", []interface{}{})
	fake.getAllServiceBrokersMutex.Unlock()
	if fake.GetAllServiceBrokersStub != nil {
		return fake.GetAllServiceBrokersStub()
	} else {
		return fake.getAllServiceBrokersReturns.result1, fake.getAllServiceBrokersReturns.result2
	}
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersCallCount() int {
	fake.getAllServiceBrokersMutex.RLock()
	defer fake.getAllServiceBrokersMutex.RUnlock()
	return len(fake.getAllServiceBrokersArgsForCall)
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersReturns(result1 []models.ServiceBroker, result2 error) {
	fake.GetAllServiceBrokersStub = nil
	fake.getAllServiceBrokersReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServices(brokerName string) (models.ServiceBroker, error) {
	fake.getBrokerWithAllServicesMutex.Lock()
	fake.getBrokerWithAllServicesArgsForCall = append(fake.getBrokerWithAllServicesArgsForCall, struct {
		brokerName string
	}{brokerName})
	fake.recordInvocation("GetBrokerWithAllServices", []interface{}{brokerName})
	fake.getBrokerWithAllServicesMutex.Unlock()
	if fake.GetBrokerWithAllServicesStub != nil {
		return fake.GetBrokerWithAllServicesStub(brokerName)
	} else {
		return fake.getBrokerWithAllServicesReturns.result1, fake.getBrokerWithAllServicesReturns.result2
	}
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesCallCount() int {
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	return len(fake.getBrokerWithAllServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesArgsForCall(i int) string {
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	return fake.getBrokerWithAllServicesArgsForCall[i].brokerName
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesReturns(result1 models.ServiceBroker, result2 error) {
	fake.GetBrokerWithAllServicesStub = nil
	fake.getBrokerWithAllServicesReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedService(serviceName string) (models.ServiceBroker, error) {
	fake.getBrokerWithSpecifiedServiceMutex.Lock()
	fake.getBrokerWithSpecifiedServiceArgsForCall = append(fake.getBrokerWithSpecifiedServiceArgsForCall, struct {
		serviceName string
	}{serviceName})
	fake.recordInvocation("GetBrokerWithSpecifiedService", []interface{}{serviceName})
	fake.getBrokerWithSpecifiedServiceMutex.Unlock()
	if fake.GetBrokerWithSpecifiedServiceStub != nil {
		return fake.GetBrokerWithSpecifiedServiceStub(serviceName)
	} else {
		return fake.getBrokerWithSpecifiedServiceReturns.result1, fake.getBrokerWithSpecifiedServiceReturns.result2
	}
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceCallCount() int {
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	return len(fake.getBrokerWithSpecifiedServiceArgsForCall)
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceArgsForCall(i int) string {
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	return fake.getBrokerWithSpecifiedServiceArgsForCall[i].serviceName
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceReturns(result1 models.ServiceBroker, result2 error) {
	fake.GetBrokerWithSpecifiedServiceStub = nil
	fake.getBrokerWithSpecifiedServiceReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	fake.getAllServiceBrokersMutex.RLock()
	defer fake.getAllServiceBrokersMutex.RUnlock()
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBrokerBuilder) recordInvocation(key string, args []interface{}) {
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

var _ brokerbuilder.BrokerBuilder = new(FakeBrokerBuilder)
