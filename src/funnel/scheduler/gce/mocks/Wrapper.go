package mocks

import compute "google.golang.org/api/compute/v1"

import mock "github.com/stretchr/testify/mock"

// Wrapper is an autogenerated mock type for the Wrapper type
type Wrapper struct {
	mock.Mock
}

// InsertInstance provides a mock function with given fields: project, zone, instance
func (_m *Wrapper) InsertInstance(project string, zone string, instance *compute.Instance) (*compute.Operation, error) {
	ret := _m.Called(project, zone, instance)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, string, *compute.Instance) *compute.Operation); ok {
		r0 = rf(project, zone, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *compute.Instance) error); ok {
		r1 = rf(project, zone, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInstanceTemplates provides a mock function with given fields: project
func (_m *Wrapper) ListInstanceTemplates(project string) (*compute.InstanceTemplateList, error) {
	ret := _m.Called(project)

	var r0 *compute.InstanceTemplateList
	if rf, ok := ret.Get(0).(func(string) *compute.InstanceTemplateList); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.InstanceTemplateList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMachineTypes provides a mock function with given fields: project, zone
func (_m *Wrapper) ListMachineTypes(project string, zone string) (*compute.MachineTypeList, error) {
	ret := _m.Called(project, zone)

	var r0 *compute.MachineTypeList
	if rf, ok := ret.Get(0).(func(string, string) *compute.MachineTypeList); ok {
		r0 = rf(project, zone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.MachineTypeList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(project, zone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}