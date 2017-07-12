// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"sync"

	"github.com/andreasf/cf-mysql-plugin/cfmysql"
)

type FakePortWaiter struct {
	WaitUntilOpenStub        func(localPort int)
	waitUntilOpenMutex       sync.RWMutex
	waitUntilOpenArgsForCall []struct {
		localPort int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePortWaiter) WaitUntilOpen(localPort int) {
	fake.waitUntilOpenMutex.Lock()
	fake.waitUntilOpenArgsForCall = append(fake.waitUntilOpenArgsForCall, struct {
		localPort int
	}{localPort})
	fake.recordInvocation("WaitUntilOpen", []interface{}{localPort})
	fake.waitUntilOpenMutex.Unlock()
	if fake.WaitUntilOpenStub != nil {
		fake.WaitUntilOpenStub(localPort)
	}
}

func (fake *FakePortWaiter) WaitUntilOpenCallCount() int {
	fake.waitUntilOpenMutex.RLock()
	defer fake.waitUntilOpenMutex.RUnlock()
	return len(fake.waitUntilOpenArgsForCall)
}

func (fake *FakePortWaiter) WaitUntilOpenArgsForCall(i int) int {
	fake.waitUntilOpenMutex.RLock()
	defer fake.waitUntilOpenMutex.RUnlock()
	return fake.waitUntilOpenArgsForCall[i].localPort
}

func (fake *FakePortWaiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.waitUntilOpenMutex.RLock()
	defer fake.waitUntilOpenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePortWaiter) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.PortWaiter = new(FakePortWaiter)
