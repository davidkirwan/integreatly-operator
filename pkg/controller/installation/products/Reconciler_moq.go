// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package products

import (
	"context"
	"sync"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	lockInterfaceMockGetPreflightObject sync.RWMutex
	lockInterfaceMockReconcile          sync.RWMutex
)

// Ensure, that InterfaceMock does implement Interface.
// If this is not the case, regenerate this file with moq.
var _ Interface = &InterfaceMock{}

// InterfaceMock is a mock implementation of Interface.
//
//     func TestSomethingThatUsesInterface(t *testing.T) {
//
//         // make and configure a mocked Interface
//         mockedInterface := &InterfaceMock{
//             GetPreflightObjectFunc: func(ns string) runtime.Object {
// 	               panic("mock out the GetPreflightObject method")
//             },
//             ReconcileFunc: func(ctx context.Context, inst *v1alpha1.Installation, product *v1alpha1.InstallationProductStatus, serverClient client.Client) (v1alpha1.StatusPhase, error) {
// 	               panic("mock out the Reconcile method")
//             },
//         }
//
//         // use mockedInterface in code that requires Interface
//         // and then make assertions.
//
//     }
type InterfaceMock struct {
	// GetPreflightObjectFunc mocks the GetPreflightObject method.
	GetPreflightObjectFunc func(ns string) runtime.Object

	// ReconcileFunc mocks the Reconcile method.
	ReconcileFunc func(ctx context.Context, installation *integreatlyv1alpha1.Installation, product *integreatlyv1alpha1.InstallationProductStatus, serverClient k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetPreflightObject holds details about calls to the GetPreflightObject method.
		GetPreflightObject []struct {
			// Ns is the ns argument value.
			Ns string
		}
		// Reconcile holds details about calls to the Reconcile method.
		Reconcile []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Inst is the inst argument value.
			Inst *integreatlyv1alpha1.Installation
			// Product is the product argument value.
			Product *integreatlyv1alpha1.InstallationProductStatus
			// ServerClient is the serverClient argument value.
			ServerClient k8sclient.Client
		}
	}
}

// GetPreflightObject calls GetPreflightObjectFunc.
func (mock *InterfaceMock) GetPreflightObject(ns string) runtime.Object {
	if mock.GetPreflightObjectFunc == nil {
		panic("InterfaceMock.GetPreflightObjectFunc: method is nil but Interface.GetPreflightObject was just called")
	}
	callInfo := struct {
		Ns string
	}{
		Ns: ns,
	}
	lockInterfaceMockGetPreflightObject.Lock()
	mock.calls.GetPreflightObject = append(mock.calls.GetPreflightObject, callInfo)
	lockInterfaceMockGetPreflightObject.Unlock()
	return mock.GetPreflightObjectFunc(ns)
}

// GetPreflightObjectCalls gets all the calls that were made to GetPreflightObject.
// Check the length with:
//     len(mockedInterface.GetPreflightObjectCalls())
func (mock *InterfaceMock) GetPreflightObjectCalls() []struct {
	Ns string
} {
	var calls []struct {
		Ns string
	}
	lockInterfaceMockGetPreflightObject.RLock()
	calls = mock.calls.GetPreflightObject
	lockInterfaceMockGetPreflightObject.RUnlock()
	return calls
}

// Reconcile calls ReconcileFunc.
func (mock *InterfaceMock) Reconcile(ctx context.Context, installation *integreatlyv1alpha1.Installation, product *integreatlyv1alpha1.InstallationProductStatus, serverClient k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	if mock.ReconcileFunc == nil {
		panic("InterfaceMock.ReconcileFunc: method is nil but Interface.Reconcile was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Inst         *integreatlyv1alpha1.Installation
		Product      *integreatlyv1alpha1.InstallationProductStatus
		ServerClient k8sclient.Client
	}{
		Ctx:          ctx,
		Inst:         installation,
		Product:      product,
		ServerClient: serverClient,
	}
	lockInterfaceMockReconcile.Lock()
	mock.calls.Reconcile = append(mock.calls.Reconcile, callInfo)
	lockInterfaceMockReconcile.Unlock()
	return mock.ReconcileFunc(ctx, installation, product, serverClient)
}

// ReconcileCalls gets all the calls that were made to Reconcile.
// Check the length with:
//     len(mockedInterface.ReconcileCalls())
func (mock *InterfaceMock) ReconcileCalls() []struct {
	Ctx          context.Context
	Inst         *integreatlyv1alpha1.Installation
	Product      *integreatlyv1alpha1.InstallationProductStatus
	ServerClient k8sclient.Client
} {
	var calls []struct {
		Ctx          context.Context
		Inst         *integreatlyv1alpha1.Installation
		Product      *integreatlyv1alpha1.InstallationProductStatus
		ServerClient k8sclient.Client
	}
	lockInterfaceMockReconcile.RLock()
	calls = mock.calls.Reconcile
	lockInterfaceMockReconcile.RUnlock()
	return calls
}
