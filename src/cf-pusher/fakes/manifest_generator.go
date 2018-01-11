// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type ManifestGenerator struct {
	GenerateStub        func(manifestStruct interface{}) (string, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		manifestStruct interface{}
	}
	generateReturns struct {
		result1 string
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ManifestGenerator) Generate(manifestStruct interface{}) (string, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		manifestStruct interface{}
	}{manifestStruct})
	fake.recordInvocation("Generate", []interface{}{manifestStruct})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub(manifestStruct)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateReturns.result1, fake.generateReturns.result2
}

func (fake *ManifestGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *ManifestGenerator) GenerateArgsForCall(i int) interface{} {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return fake.generateArgsForCall[i].manifestStruct
}

func (fake *ManifestGenerator) GenerateReturns(result1 string, result2 error) {
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ManifestGenerator) GenerateReturnsOnCall(i int, result1 string, result2 error) {
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ManifestGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ManifestGenerator) recordInvocation(key string, args []interface{}) {
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
