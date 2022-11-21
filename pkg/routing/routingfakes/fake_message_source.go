// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FakeMessageSource struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ReadChanStub        func() <-chan protoreflect.ProtoMessage
	readChanMutex       sync.RWMutex
	readChanArgsForCall []struct {
	}
	readChanReturns struct {
		result1 <-chan protoreflect.ProtoMessage
	}
	readChanReturnsOnCall map[int]struct {
		result1 <-chan protoreflect.ProtoMessage
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMessageSource) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeMessageSource) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeMessageSource) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeMessageSource) ReadChan() <-chan protoreflect.ProtoMessage {
	fake.readChanMutex.Lock()
	ret, specificReturn := fake.readChanReturnsOnCall[len(fake.readChanArgsForCall)]
	fake.readChanArgsForCall = append(fake.readChanArgsForCall, struct {
	}{})
	stub := fake.ReadChanStub
	fakeReturns := fake.readChanReturns
	fake.recordInvocation("ReadChan", []interface{}{})
	fake.readChanMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMessageSource) ReadChanCallCount() int {
	fake.readChanMutex.RLock()
	defer fake.readChanMutex.RUnlock()
	return len(fake.readChanArgsForCall)
}

func (fake *FakeMessageSource) ReadChanCalls(stub func() <-chan protoreflect.ProtoMessage) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = stub
}

func (fake *FakeMessageSource) ReadChanReturns(result1 <-chan protoreflect.ProtoMessage) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = nil
	fake.readChanReturns = struct {
		result1 <-chan protoreflect.ProtoMessage
	}{result1}
}

func (fake *FakeMessageSource) ReadChanReturnsOnCall(i int, result1 <-chan protoreflect.ProtoMessage) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = nil
	if fake.readChanReturnsOnCall == nil {
		fake.readChanReturnsOnCall = make(map[int]struct {
			result1 <-chan protoreflect.ProtoMessage
		})
	}
	fake.readChanReturnsOnCall[i] = struct {
		result1 <-chan protoreflect.ProtoMessage
	}{result1}
}

func (fake *FakeMessageSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.readChanMutex.RLock()
	defer fake.readChanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMessageSource) recordInvocation(key string, args []interface{}) {
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

var _ routing.MessageSource = new(FakeMessageSource)
