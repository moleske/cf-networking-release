// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type OrgCheckerCli struct {
	TargetOrgStub        func(name string) error
	targetOrgMutex       sync.RWMutex
	targetOrgArgsForCall []struct {
		name string
	}
	targetOrgReturns struct {
		result1 error
	}
	targetOrgReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrgCheckerCli) TargetOrg(name string) error {
	fake.targetOrgMutex.Lock()
	ret, specificReturn := fake.targetOrgReturnsOnCall[len(fake.targetOrgArgsForCall)]
	fake.targetOrgArgsForCall = append(fake.targetOrgArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("TargetOrg", []interface{}{name})
	fake.targetOrgMutex.Unlock()
	if fake.TargetOrgStub != nil {
		return fake.TargetOrgStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.targetOrgReturns.result1
}

func (fake *OrgCheckerCli) TargetOrgCallCount() int {
	fake.targetOrgMutex.RLock()
	defer fake.targetOrgMutex.RUnlock()
	return len(fake.targetOrgArgsForCall)
}

func (fake *OrgCheckerCli) TargetOrgArgsForCall(i int) string {
	fake.targetOrgMutex.RLock()
	defer fake.targetOrgMutex.RUnlock()
	return fake.targetOrgArgsForCall[i].name
}

func (fake *OrgCheckerCli) TargetOrgReturns(result1 error) {
	fake.TargetOrgStub = nil
	fake.targetOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *OrgCheckerCli) TargetOrgReturnsOnCall(i int, result1 error) {
	fake.TargetOrgStub = nil
	if fake.targetOrgReturnsOnCall == nil {
		fake.targetOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.targetOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OrgCheckerCli) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.targetOrgMutex.RLock()
	defer fake.targetOrgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrgCheckerCli) recordInvocation(key string, args []interface{}) {
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
