// This file was generated by counterfeiter
package sshfakes

import (
	"net"
	"sync"

	"github.com/liamawhite/cli-with-i18n/cf/ssh"
)

type FakeListenerFactory struct {
	ListenStub        func(network, address string) (net.Listener, error)
	listenMutex       sync.RWMutex
	listenArgsForCall []struct {
		network string
		address string
	}
	listenReturns struct {
		result1 net.Listener
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeListenerFactory) Listen(network string, address string) (net.Listener, error) {
	fake.listenMutex.Lock()
	fake.listenArgsForCall = append(fake.listenArgsForCall, struct {
		network string
		address string
	}{network, address})
	fake.recordInvocation("Listen", []interface{}{network, address})
	fake.listenMutex.Unlock()
	if fake.ListenStub != nil {
		return fake.ListenStub(network, address)
	} else {
		return fake.listenReturns.result1, fake.listenReturns.result2
	}
}

func (fake *FakeListenerFactory) ListenCallCount() int {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return len(fake.listenArgsForCall)
}

func (fake *FakeListenerFactory) ListenArgsForCall(i int) (string, string) {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return fake.listenArgsForCall[i].network, fake.listenArgsForCall[i].address
}

func (fake *FakeListenerFactory) ListenReturns(result1 net.Listener, result2 error) {
	fake.ListenStub = nil
	fake.listenReturns = struct {
		result1 net.Listener
		result2 error
	}{result1, result2}
}

func (fake *FakeListenerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeListenerFactory) recordInvocation(key string, args []interface{}) {
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

var _ sshCmd.ListenerFactory = new(FakeListenerFactory)
