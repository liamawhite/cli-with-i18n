// This file was generated by counterfeiter
package logsfakes

import (
	"sync"
	"time"

	"github.com/liamawhite/cli-with-i18n/cf/api/logs"
)

type FakeLoggable struct {
	ToLogStub        func(loc *time.Location) string
	toLogMutex       sync.RWMutex
	toLogArgsForCall []struct {
		loc *time.Location
	}
	toLogReturns struct {
		result1 string
	}
	ToSimpleLogStub        func() string
	toSimpleLogMutex       sync.RWMutex
	toSimpleLogArgsForCall []struct{}
	toSimpleLogReturns     struct {
		result1 string
	}
	GetSourceNameStub        func() string
	getSourceNameMutex       sync.RWMutex
	getSourceNameArgsForCall []struct{}
	getSourceNameReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoggable) ToLog(loc *time.Location) string {
	fake.toLogMutex.Lock()
	fake.toLogArgsForCall = append(fake.toLogArgsForCall, struct {
		loc *time.Location
	}{loc})
	fake.recordInvocation("ToLog", []interface{}{loc})
	fake.toLogMutex.Unlock()
	if fake.ToLogStub != nil {
		return fake.ToLogStub(loc)
	} else {
		return fake.toLogReturns.result1
	}
}

func (fake *FakeLoggable) ToLogCallCount() int {
	fake.toLogMutex.RLock()
	defer fake.toLogMutex.RUnlock()
	return len(fake.toLogArgsForCall)
}

func (fake *FakeLoggable) ToLogArgsForCall(i int) *time.Location {
	fake.toLogMutex.RLock()
	defer fake.toLogMutex.RUnlock()
	return fake.toLogArgsForCall[i].loc
}

func (fake *FakeLoggable) ToLogReturns(result1 string) {
	fake.ToLogStub = nil
	fake.toLogReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLoggable) ToSimpleLog() string {
	fake.toSimpleLogMutex.Lock()
	fake.toSimpleLogArgsForCall = append(fake.toSimpleLogArgsForCall, struct{}{})
	fake.recordInvocation("ToSimpleLog", []interface{}{})
	fake.toSimpleLogMutex.Unlock()
	if fake.ToSimpleLogStub != nil {
		return fake.ToSimpleLogStub()
	} else {
		return fake.toSimpleLogReturns.result1
	}
}

func (fake *FakeLoggable) ToSimpleLogCallCount() int {
	fake.toSimpleLogMutex.RLock()
	defer fake.toSimpleLogMutex.RUnlock()
	return len(fake.toSimpleLogArgsForCall)
}

func (fake *FakeLoggable) ToSimpleLogReturns(result1 string) {
	fake.ToSimpleLogStub = nil
	fake.toSimpleLogReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLoggable) GetSourceName() string {
	fake.getSourceNameMutex.Lock()
	fake.getSourceNameArgsForCall = append(fake.getSourceNameArgsForCall, struct{}{})
	fake.recordInvocation("GetSourceName", []interface{}{})
	fake.getSourceNameMutex.Unlock()
	if fake.GetSourceNameStub != nil {
		return fake.GetSourceNameStub()
	} else {
		return fake.getSourceNameReturns.result1
	}
}

func (fake *FakeLoggable) GetSourceNameCallCount() int {
	fake.getSourceNameMutex.RLock()
	defer fake.getSourceNameMutex.RUnlock()
	return len(fake.getSourceNameArgsForCall)
}

func (fake *FakeLoggable) GetSourceNameReturns(result1 string) {
	fake.GetSourceNameStub = nil
	fake.getSourceNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLoggable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.toLogMutex.RLock()
	defer fake.toLogMutex.RUnlock()
	fake.toSimpleLogMutex.RLock()
	defer fake.toSimpleLogMutex.RUnlock()
	fake.getSourceNameMutex.RLock()
	defer fake.getSourceNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLoggable) recordInvocation(key string, args []interface{}) {
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

var _ logs.Loggable = new(FakeLoggable)
