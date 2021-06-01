// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package plugin

import (
	"sync"
)

// Ensure, that RPCClientMock does implement RPCClient.
// If this is not the case, regenerate this file with moq.
var _ RPCClient = &RPCClientMock{}

// RPCClientMock is a mock implementation of RPCClient.
//
// 	func TestSomethingThatUsesRPCClient(t *testing.T) {
//
// 		// make and configure a mocked RPCClient
// 		mockedRPCClient := &RPCClientMock{
// 			CallFunc: func(serviceMethod string, args interface{}, reply interface{}) error {
// 				panic("mock out the Call method")
// 			},
// 		}
//
// 		// use mockedRPCClient in code that requires RPCClient
// 		// and then make assertions.
//
// 	}
type RPCClientMock struct {
	// CallFunc mocks the Call method.
	CallFunc func(serviceMethod string, args interface{}, reply interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Call holds details about calls to the Call method.
		Call []struct {
			// ServiceMethod is the serviceMethod argument value.
			ServiceMethod string
			// Args is the args argument value.
			Args interface{}
			// Reply is the reply argument value.
			Reply interface{}
		}
	}
	lockCall sync.RWMutex
}

// Call calls CallFunc.
func (mock *RPCClientMock) Call(serviceMethod string, args interface{}, reply interface{}) error {
	if mock.CallFunc == nil {
		panic("RPCClientMock.CallFunc: method is nil but RPCClient.Call was just called")
	}
	callInfo := struct {
		ServiceMethod string
		Args          interface{}
		Reply         interface{}
	}{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
	}
	mock.lockCall.Lock()
	mock.calls.Call = append(mock.calls.Call, callInfo)
	mock.lockCall.Unlock()
	return mock.CallFunc(serviceMethod, args, reply)
}

// CallCalls gets all the calls that were made to Call.
// Check the length with:
//     len(mockedRPCClient.CallCalls())
func (mock *RPCClientMock) CallCalls() []struct {
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
} {
	var calls []struct {
		ServiceMethod string
		Args          interface{}
		Reply         interface{}
	}
	mock.lockCall.RLock()
	calls = mock.calls.Call
	mock.lockCall.RUnlock()
	return calls
}
