// Code generated by counterfeiter. DO NOT EDIT.
package cfnetworkingactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/resources"
)

type FakeCloudControllerClient struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (resources.Application, ccv3.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 resources.Application
		result2 ccv3.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 resources.Application
		result2 ccv3.Warnings
		result3 error
	}
	GetApplicationsStub        func(...ccv3.Query) ([]resources.Application, ccv3.Warnings, error)
	getApplicationsMutex       sync.RWMutex
	getApplicationsArgsForCall []struct {
		arg1 []ccv3.Query
	}
	getApplicationsReturns struct {
		result1 []resources.Application
		result2 ccv3.Warnings
		result3 error
	}
	getApplicationsReturnsOnCall map[int]struct {
		result1 []resources.Application
		result2 ccv3.Warnings
		result3 error
	}
	GetOrganizationsStub        func(...ccv3.Query) ([]resources.Organization, ccv3.Warnings, error)
	getOrganizationsMutex       sync.RWMutex
	getOrganizationsArgsForCall []struct {
		arg1 []ccv3.Query
	}
	getOrganizationsReturns struct {
		result1 []resources.Organization
		result2 ccv3.Warnings
		result3 error
	}
	getOrganizationsReturnsOnCall map[int]struct {
		result1 []resources.Organization
		result2 ccv3.Warnings
		result3 error
	}
	GetSpacesStub        func(...ccv3.Query) ([]ccv3.Space, ccv3.IncludedResources, ccv3.Warnings, error)
	getSpacesMutex       sync.RWMutex
	getSpacesArgsForCall []struct {
		arg1 []ccv3.Query
	}
	getSpacesReturns struct {
		result1 []ccv3.Space
		result2 ccv3.IncludedResources
		result3 ccv3.Warnings
		result4 error
	}
	getSpacesReturnsOnCall map[int]struct {
		result1 []ccv3.Space
		result2 ccv3.IncludedResources
		result3 ccv3.Warnings
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpace(arg1 string, arg2 string) (resources.Application, ccv3.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpaceCalls(stub func(string, string) (resources.Application, ccv3.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpaceReturns(result1 resources.Application, result2 ccv3.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 resources.Application
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 resources.Application, result2 ccv3.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 resources.Application
			result2 ccv3.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 resources.Application
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetApplications(arg1 ...ccv3.Query) ([]resources.Application, ccv3.Warnings, error) {
	fake.getApplicationsMutex.Lock()
	ret, specificReturn := fake.getApplicationsReturnsOnCall[len(fake.getApplicationsArgsForCall)]
	fake.getApplicationsArgsForCall = append(fake.getApplicationsArgsForCall, struct {
		arg1 []ccv3.Query
	}{arg1})
	fake.recordInvocation("GetApplications", []interface{}{arg1})
	fake.getApplicationsMutex.Unlock()
	if fake.GetApplicationsStub != nil {
		return fake.GetApplicationsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCloudControllerClient) GetApplicationsCallCount() int {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return len(fake.getApplicationsArgsForCall)
}

func (fake *FakeCloudControllerClient) GetApplicationsCalls(stub func(...ccv3.Query) ([]resources.Application, ccv3.Warnings, error)) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = stub
}

func (fake *FakeCloudControllerClient) GetApplicationsArgsForCall(i int) []ccv3.Query {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	argsForCall := fake.getApplicationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudControllerClient) GetApplicationsReturns(result1 []resources.Application, result2 ccv3.Warnings, result3 error) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = nil
	fake.getApplicationsReturns = struct {
		result1 []resources.Application
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetApplicationsReturnsOnCall(i int, result1 []resources.Application, result2 ccv3.Warnings, result3 error) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = nil
	if fake.getApplicationsReturnsOnCall == nil {
		fake.getApplicationsReturnsOnCall = make(map[int]struct {
			result1 []resources.Application
			result2 ccv3.Warnings
			result3 error
		})
	}
	fake.getApplicationsReturnsOnCall[i] = struct {
		result1 []resources.Application
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetOrganizations(arg1 ...ccv3.Query) ([]resources.Organization, ccv3.Warnings, error) {
	fake.getOrganizationsMutex.Lock()
	ret, specificReturn := fake.getOrganizationsReturnsOnCall[len(fake.getOrganizationsArgsForCall)]
	fake.getOrganizationsArgsForCall = append(fake.getOrganizationsArgsForCall, struct {
		arg1 []ccv3.Query
	}{arg1})
	fake.recordInvocation("GetOrganizations", []interface{}{arg1})
	fake.getOrganizationsMutex.Unlock()
	if fake.GetOrganizationsStub != nil {
		return fake.GetOrganizationsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCloudControllerClient) GetOrganizationsCallCount() int {
	fake.getOrganizationsMutex.RLock()
	defer fake.getOrganizationsMutex.RUnlock()
	return len(fake.getOrganizationsArgsForCall)
}

func (fake *FakeCloudControllerClient) GetOrganizationsCalls(stub func(...ccv3.Query) ([]resources.Organization, ccv3.Warnings, error)) {
	fake.getOrganizationsMutex.Lock()
	defer fake.getOrganizationsMutex.Unlock()
	fake.GetOrganizationsStub = stub
}

func (fake *FakeCloudControllerClient) GetOrganizationsArgsForCall(i int) []ccv3.Query {
	fake.getOrganizationsMutex.RLock()
	defer fake.getOrganizationsMutex.RUnlock()
	argsForCall := fake.getOrganizationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudControllerClient) GetOrganizationsReturns(result1 []resources.Organization, result2 ccv3.Warnings, result3 error) {
	fake.getOrganizationsMutex.Lock()
	defer fake.getOrganizationsMutex.Unlock()
	fake.GetOrganizationsStub = nil
	fake.getOrganizationsReturns = struct {
		result1 []resources.Organization
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetOrganizationsReturnsOnCall(i int, result1 []resources.Organization, result2 ccv3.Warnings, result3 error) {
	fake.getOrganizationsMutex.Lock()
	defer fake.getOrganizationsMutex.Unlock()
	fake.GetOrganizationsStub = nil
	if fake.getOrganizationsReturnsOnCall == nil {
		fake.getOrganizationsReturnsOnCall = make(map[int]struct {
			result1 []resources.Organization
			result2 ccv3.Warnings
			result3 error
		})
	}
	fake.getOrganizationsReturnsOnCall[i] = struct {
		result1 []resources.Organization
		result2 ccv3.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCloudControllerClient) GetSpaces(arg1 ...ccv3.Query) ([]ccv3.Space, ccv3.IncludedResources, ccv3.Warnings, error) {
	fake.getSpacesMutex.Lock()
	ret, specificReturn := fake.getSpacesReturnsOnCall[len(fake.getSpacesArgsForCall)]
	fake.getSpacesArgsForCall = append(fake.getSpacesArgsForCall, struct {
		arg1 []ccv3.Query
	}{arg1})
	fake.recordInvocation("GetSpaces", []interface{}{arg1})
	fake.getSpacesMutex.Unlock()
	if fake.GetSpacesStub != nil {
		return fake.GetSpacesStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getSpacesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeCloudControllerClient) GetSpacesCallCount() int {
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	return len(fake.getSpacesArgsForCall)
}

func (fake *FakeCloudControllerClient) GetSpacesCalls(stub func(...ccv3.Query) ([]ccv3.Space, ccv3.IncludedResources, ccv3.Warnings, error)) {
	fake.getSpacesMutex.Lock()
	defer fake.getSpacesMutex.Unlock()
	fake.GetSpacesStub = stub
}

func (fake *FakeCloudControllerClient) GetSpacesArgsForCall(i int) []ccv3.Query {
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	argsForCall := fake.getSpacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudControllerClient) GetSpacesReturns(result1 []ccv3.Space, result2 ccv3.IncludedResources, result3 ccv3.Warnings, result4 error) {
	fake.getSpacesMutex.Lock()
	defer fake.getSpacesMutex.Unlock()
	fake.GetSpacesStub = nil
	fake.getSpacesReturns = struct {
		result1 []ccv3.Space
		result2 ccv3.IncludedResources
		result3 ccv3.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeCloudControllerClient) GetSpacesReturnsOnCall(i int, result1 []ccv3.Space, result2 ccv3.IncludedResources, result3 ccv3.Warnings, result4 error) {
	fake.getSpacesMutex.Lock()
	defer fake.getSpacesMutex.Unlock()
	fake.GetSpacesStub = nil
	if fake.getSpacesReturnsOnCall == nil {
		fake.getSpacesReturnsOnCall = make(map[int]struct {
			result1 []ccv3.Space
			result2 ccv3.IncludedResources
			result3 ccv3.Warnings
			result4 error
		})
	}
	fake.getSpacesReturnsOnCall[i] = struct {
		result1 []ccv3.Space
		result2 ccv3.IncludedResources
		result3 ccv3.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeCloudControllerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	fake.getOrganizationsMutex.RLock()
	defer fake.getOrganizationsMutex.RUnlock()
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudControllerClient) recordInvocation(key string, args []interface{}) {
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

var _ cfnetworkingaction.CloudControllerClient = new(FakeCloudControllerClient)
