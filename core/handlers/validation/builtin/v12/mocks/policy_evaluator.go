// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/hyperledger/fabric-gm/protos/common"
import mock "github.com/stretchr/testify/mock"

// PolicyEvaluator is an autogenerated mock type for the PolicyEvaluator type
type PolicyEvaluator struct {
	mock.Mock
}

// Evaluate provides a mock function with given fields: policyBytes, signatureSet
func (_m *PolicyEvaluator) Evaluate(policyBytes []byte, signatureSet []*common.SignedData) error {
	ret := _m.Called(policyBytes, signatureSet)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []*common.SignedData) error); ok {
		r0 = rf(policyBytes, signatureSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}