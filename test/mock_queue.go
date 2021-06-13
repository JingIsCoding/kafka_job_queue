// Code generated by mockery v2.8.0. DO NOT EDIT.

package test

import (
	job "github.com/JingIsCoding/kafka_job_queue/job"
	log "github.com/JingIsCoding/kafka_job_queue/log"

	mock "github.com/stretchr/testify/mock"

	queue "github.com/JingIsCoding/kafka_job_queue/queue"
)

// MockQueue is an autogenerated mock type for the Queue type
type MockQueue struct {
	mock.Mock
}

// Enqueue provides a mock function with given fields: _a0
func (_m *MockQueue) Enqueue(_a0 job.Job) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(job.Job) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnqueueTo provides a mock function with given fields: _a0, _a1
func (_m *MockQueue) EnqueueTo(_a0 job.Job, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(job.Job, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *MockQueue) GetConfig() queue.Config {
	ret := _m.Called()

	var r0 queue.Config
	if rf, ok := ret.Get(0).(func() queue.Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(queue.Config)
	}

	return r0
}

// GetDoneChan provides a mock function with given fields:
func (_m *MockQueue) GetDoneChan() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}

// GetJobDefinition provides a mock function with given fields: jobName
func (_m *MockQueue) GetJobDefinition(jobName string) (*job.JobDefinition, error) {
	ret := _m.Called(jobName)

	var r0 *job.JobDefinition
	if rf, ok := ret.Get(0).(func(string) *job.JobDefinition); ok {
		r0 = rf(jobName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*job.JobDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *MockQueue) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: _a0
func (_m *MockQueue) Register(_a0 job.JobDefinition) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(job.JobDefinition) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogger provides a mock function with given fields: _a0
func (_m *MockQueue) SetLogger(_a0 log.Logger) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields:
func (_m *MockQueue) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *MockQueue) Stop() {
	_m.Called()
}
