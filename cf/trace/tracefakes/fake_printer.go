// This file was generated by counterfeiter
package tracefakes

import (
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/trace"
)

type FakePrinter struct {
	PrintStub        func(v ...interface{})
	printMutex       sync.RWMutex
	printArgsForCall []struct {
		v []interface{}
	}
	PrintfStub        func(format string, v ...interface{})
	printfMutex       sync.RWMutex
	printfArgsForCall []struct {
		format string
		v      []interface{}
	}
	PrintlnStub        func(v ...interface{})
	printlnMutex       sync.RWMutex
	printlnArgsForCall []struct {
		v []interface{}
	}
	WritesToConsoleStub        func() bool
	writesToConsoleMutex       sync.RWMutex
	writesToConsoleArgsForCall []struct{}
	writesToConsoleReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePrinter) Print(v ...interface{}) {
	fake.printMutex.Lock()
	fake.printArgsForCall = append(fake.printArgsForCall, struct {
		v []interface{}
	}{v})
	fake.recordInvocation("Print", []interface{}{v})
	fake.printMutex.Unlock()
	if fake.PrintStub != nil {
		fake.PrintStub(v...)
	}
}

func (fake *FakePrinter) PrintCallCount() int {
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	return len(fake.printArgsForCall)
}

func (fake *FakePrinter) PrintArgsForCall(i int) []interface{} {
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	return fake.printArgsForCall[i].v
}

func (fake *FakePrinter) Printf(format string, v ...interface{}) {
	fake.printfMutex.Lock()
	fake.printfArgsForCall = append(fake.printfArgsForCall, struct {
		format string
		v      []interface{}
	}{format, v})
	fake.recordInvocation("Printf", []interface{}{format, v})
	fake.printfMutex.Unlock()
	if fake.PrintfStub != nil {
		fake.PrintfStub(format, v...)
	}
}

func (fake *FakePrinter) PrintfCallCount() int {
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	return len(fake.printfArgsForCall)
}

func (fake *FakePrinter) PrintfArgsForCall(i int) (string, []interface{}) {
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	return fake.printfArgsForCall[i].format, fake.printfArgsForCall[i].v
}

func (fake *FakePrinter) Println(v ...interface{}) {
	fake.printlnMutex.Lock()
	fake.printlnArgsForCall = append(fake.printlnArgsForCall, struct {
		v []interface{}
	}{v})
	fake.recordInvocation("Println", []interface{}{v})
	fake.printlnMutex.Unlock()
	if fake.PrintlnStub != nil {
		fake.PrintlnStub(v...)
	}
}

func (fake *FakePrinter) PrintlnCallCount() int {
	fake.printlnMutex.RLock()
	defer fake.printlnMutex.RUnlock()
	return len(fake.printlnArgsForCall)
}

func (fake *FakePrinter) PrintlnArgsForCall(i int) []interface{} {
	fake.printlnMutex.RLock()
	defer fake.printlnMutex.RUnlock()
	return fake.printlnArgsForCall[i].v
}

func (fake *FakePrinter) WritesToConsole() bool {
	fake.writesToConsoleMutex.Lock()
	fake.writesToConsoleArgsForCall = append(fake.writesToConsoleArgsForCall, struct{}{})
	fake.recordInvocation("WritesToConsole", []interface{}{})
	fake.writesToConsoleMutex.Unlock()
	if fake.WritesToConsoleStub != nil {
		return fake.WritesToConsoleStub()
	} else {
		return fake.writesToConsoleReturns.result1
	}
}

func (fake *FakePrinter) WritesToConsoleCallCount() int {
	fake.writesToConsoleMutex.RLock()
	defer fake.writesToConsoleMutex.RUnlock()
	return len(fake.writesToConsoleArgsForCall)
}

func (fake *FakePrinter) WritesToConsoleReturns(result1 bool) {
	fake.WritesToConsoleStub = nil
	fake.writesToConsoleReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePrinter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	fake.printlnMutex.RLock()
	defer fake.printlnMutex.RUnlock()
	fake.writesToConsoleMutex.RLock()
	defer fake.writesToConsoleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePrinter) recordInvocation(key string, args []interface{}) {
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

var _ trace.Printer = new(FakePrinter)
