// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/playnet-public/gorcon/pkg/event"
	"github.com/playnet-public/gorcon/pkg/rcon"
)

type RconConnection struct {
	OpenStub        func(context.Context) error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		arg1 context.Context
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func(context.Context) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 context.Context
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func(context.Context, string) (rcon.Transmission, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	writeReturns struct {
		result1 rcon.Transmission
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 rcon.Transmission
		result2 error
	}
	SubscribeStub        func(context.Context, chan<- event.Event)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		arg1 context.Context
		arg2 chan<- event.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RconConnection) Open(arg1 context.Context) error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Open", []interface{}{arg1})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openReturns.result1
}

func (fake *RconConnection) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *RconConnection) OpenArgsForCall(i int) context.Context {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].arg1
}

func (fake *RconConnection) OpenReturns(result1 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *RconConnection) OpenReturnsOnCall(i int, result1 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RconConnection) Close(arg1 context.Context) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Close", []interface{}{arg1})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *RconConnection) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *RconConnection) CloseArgsForCall(i int) context.Context {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.closeArgsForCall[i].arg1
}

func (fake *RconConnection) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *RconConnection) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RconConnection) Write(arg1 context.Context, arg2 string) (rcon.Transmission, error) {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Write", []interface{}{arg1, arg2})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *RconConnection) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *RconConnection) WriteArgsForCall(i int) (context.Context, string) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].arg1, fake.writeArgsForCall[i].arg2
}

func (fake *RconConnection) WriteReturns(result1 rcon.Transmission, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 rcon.Transmission
		result2 error
	}{result1, result2}
}

func (fake *RconConnection) WriteReturnsOnCall(i int, result1 rcon.Transmission, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 rcon.Transmission
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 rcon.Transmission
		result2 error
	}{result1, result2}
}

func (fake *RconConnection) Subscribe(arg1 context.Context, arg2 chan<- event.Event) {
	fake.subscribeMutex.Lock()
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		arg1 context.Context
		arg2 chan<- event.Event
	}{arg1, arg2})
	fake.recordInvocation("Subscribe", []interface{}{arg1, arg2})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		fake.SubscribeStub(arg1, arg2)
	}
}

func (fake *RconConnection) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *RconConnection) SubscribeArgsForCall(i int) (context.Context, chan<- event.Event) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.subscribeArgsForCall[i].arg1, fake.subscribeArgsForCall[i].arg2
}

func (fake *RconConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RconConnection) recordInvocation(key string, args []interface{}) {
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

var _ rcon.Connection = new(RconConnection)
