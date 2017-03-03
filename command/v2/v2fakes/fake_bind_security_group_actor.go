// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeBindSecurityGroupActor struct {
	GetOrganizationByNameStub        func(orgName string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationByNameReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationSpacesStub        func(orgGUID string) ([]v2action.Space, v2action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		orgGUID string
	}
	getOrganizationSpacesReturns struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	GetSpaceByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error)
	getSpaceByOrganizationAndNameMutex       sync.RWMutex
	getSpaceByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceByOrganizationAndNameReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	GetSecurityGroupByNameStub        func(securityGroupName string) (v2action.SecurityGroup, v2action.Warnings, error)
	getSecurityGroupByNameMutex       sync.RWMutex
	getSecurityGroupByNameArgsForCall []struct {
		securityGroupName string
	}
	getSecurityGroupByNameReturns struct {
		result1 v2action.SecurityGroup
		result2 v2action.Warnings
		result3 error
	}
	BindSecurityGroupToSpaceStub        func(securityGroupGUID string, spaceGUID string) (v2action.Warnings, error)
	bindSecurityGroupToSpaceMutex       sync.RWMutex
	bindSecurityGroupToSpaceArgsForCall []struct {
		securityGroupGUID string
		spaceGUID         string
	}
	bindSecurityGroupToSpaceReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByName(orgName string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationByName", []interface{}{orgName})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(orgName)
	} else {
		return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
	}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpaces(orgGUID string) ([]v2action.Space, v2action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{orgGUID})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(orgGUID)
	} else {
		return fake.getOrganizationSpacesReturns.result1, fake.getOrganizationSpacesReturns.result2, fake.getOrganizationSpacesReturns.result3
	}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return fake.getOrganizationSpacesArgsForCall[i].orgGUID
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesReturns(result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByOrganizationAndName(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error) {
	fake.getSpaceByOrganizationAndNameMutex.Lock()
	fake.getSpaceByOrganizationAndNameArgsForCall = append(fake.getSpaceByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceByOrganizationAndNameStub != nil {
		return fake.GetSpaceByOrganizationAndNameStub(orgGUID, spaceName)
	} else {
		return fake.getSpaceByOrganizationAndNameReturns.result1, fake.getSpaceByOrganizationAndNameReturns.result2, fake.getSpaceByOrganizationAndNameReturns.result3
	}
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByOrganizationAndNameCallCount() int {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceByOrganizationAndNameArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByOrganizationAndNameReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	fake.getSpaceByOrganizationAndNameReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupByName(securityGroupName string) (v2action.SecurityGroup, v2action.Warnings, error) {
	fake.getSecurityGroupByNameMutex.Lock()
	fake.getSecurityGroupByNameArgsForCall = append(fake.getSecurityGroupByNameArgsForCall, struct {
		securityGroupName string
	}{securityGroupName})
	fake.recordInvocation("GetSecurityGroupByName", []interface{}{securityGroupName})
	fake.getSecurityGroupByNameMutex.Unlock()
	if fake.GetSecurityGroupByNameStub != nil {
		return fake.GetSecurityGroupByNameStub(securityGroupName)
	} else {
		return fake.getSecurityGroupByNameReturns.result1, fake.getSecurityGroupByNameReturns.result2, fake.getSecurityGroupByNameReturns.result3
	}
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupByNameCallCount() int {
	fake.getSecurityGroupByNameMutex.RLock()
	defer fake.getSecurityGroupByNameMutex.RUnlock()
	return len(fake.getSecurityGroupByNameArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupByNameArgsForCall(i int) string {
	fake.getSecurityGroupByNameMutex.RLock()
	defer fake.getSecurityGroupByNameMutex.RUnlock()
	return fake.getSecurityGroupByNameArgsForCall[i].securityGroupName
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupByNameReturns(result1 v2action.SecurityGroup, result2 v2action.Warnings, result3 error) {
	fake.GetSecurityGroupByNameStub = nil
	fake.getSecurityGroupByNameReturns = struct {
		result1 v2action.SecurityGroup
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpace(securityGroupGUID string, spaceGUID string) (v2action.Warnings, error) {
	fake.bindSecurityGroupToSpaceMutex.Lock()
	fake.bindSecurityGroupToSpaceArgsForCall = append(fake.bindSecurityGroupToSpaceArgsForCall, struct {
		securityGroupGUID string
		spaceGUID         string
	}{securityGroupGUID, spaceGUID})
	fake.recordInvocation("BindSecurityGroupToSpace", []interface{}{securityGroupGUID, spaceGUID})
	fake.bindSecurityGroupToSpaceMutex.Unlock()
	if fake.BindSecurityGroupToSpaceStub != nil {
		return fake.BindSecurityGroupToSpaceStub(securityGroupGUID, spaceGUID)
	} else {
		return fake.bindSecurityGroupToSpaceReturns.result1, fake.bindSecurityGroupToSpaceReturns.result2
	}
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceCallCount() int {
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	return len(fake.bindSecurityGroupToSpaceArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceArgsForCall(i int) (string, string) {
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	return fake.bindSecurityGroupToSpaceArgsForCall[i].securityGroupGUID, fake.bindSecurityGroupToSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceReturns(result1 v2action.Warnings, result2 error) {
	fake.BindSecurityGroupToSpaceStub = nil
	fake.bindSecurityGroupToSpaceReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeBindSecurityGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	fake.getSecurityGroupByNameMutex.RLock()
	defer fake.getSecurityGroupByNameMutex.RUnlock()
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBindSecurityGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.BindSecurityGroupActor = new(FakeBindSecurityGroupActor)