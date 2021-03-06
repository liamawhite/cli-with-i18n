// Code generated by counterfeiter. DO NOT EDIT.
package pushactionfakes

import (
	"io"
	"sync"

	"github.com/liamawhite/cli-with-i18n/actor/pushaction"
)

type FakeProgressBar struct {
	NewProgressBarWrapperStub        func(reader io.Reader, sizeOfFile int64) io.Reader
	newProgressBarWrapperMutex       sync.RWMutex
	newProgressBarWrapperArgsForCall []struct {
		reader     io.Reader
		sizeOfFile int64
	}
	newProgressBarWrapperReturns struct {
		result1 io.Reader
	}
	newProgressBarWrapperReturnsOnCall map[int]struct {
		result1 io.Reader
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProgressBar) NewProgressBarWrapper(reader io.Reader, sizeOfFile int64) io.Reader {
	fake.newProgressBarWrapperMutex.Lock()
	ret, specificReturn := fake.newProgressBarWrapperReturnsOnCall[len(fake.newProgressBarWrapperArgsForCall)]
	fake.newProgressBarWrapperArgsForCall = append(fake.newProgressBarWrapperArgsForCall, struct {
		reader     io.Reader
		sizeOfFile int64
	}{reader, sizeOfFile})
	fake.recordInvocation("NewProgressBarWrapper", []interface{}{reader, sizeOfFile})
	fake.newProgressBarWrapperMutex.Unlock()
	if fake.NewProgressBarWrapperStub != nil {
		return fake.NewProgressBarWrapperStub(reader, sizeOfFile)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newProgressBarWrapperReturns.result1
}

func (fake *FakeProgressBar) NewProgressBarWrapperCallCount() int {
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	return len(fake.newProgressBarWrapperArgsForCall)
}

func (fake *FakeProgressBar) NewProgressBarWrapperArgsForCall(i int) (io.Reader, int64) {
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	return fake.newProgressBarWrapperArgsForCall[i].reader, fake.newProgressBarWrapperArgsForCall[i].sizeOfFile
}

func (fake *FakeProgressBar) NewProgressBarWrapperReturns(result1 io.Reader) {
	fake.NewProgressBarWrapperStub = nil
	fake.newProgressBarWrapperReturns = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeProgressBar) NewProgressBarWrapperReturnsOnCall(i int, result1 io.Reader) {
	fake.NewProgressBarWrapperStub = nil
	if fake.newProgressBarWrapperReturnsOnCall == nil {
		fake.newProgressBarWrapperReturnsOnCall = make(map[int]struct {
			result1 io.Reader
		})
	}
	fake.newProgressBarWrapperReturnsOnCall[i] = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeProgressBar) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProgressBar) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.ProgressBar = new(FakeProgressBar)
