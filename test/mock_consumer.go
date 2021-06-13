// Code generated by mockery v2.8.0. DO NOT EDIT.

package test

import (
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	mock "github.com/stretchr/testify/mock"
)

// MockConsumer is an autogenerated mock type for the Consumer type
type MockConsumer struct {
	mock.Mock
}

// Assign provides a mock function with given fields: _a0
func (_m *MockConsumer) Assign(_a0 []kafka.TopicPartition) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]kafka.TopicPartition) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockConsumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommitMessage provides a mock function with given fields: _a0
func (_m *MockConsumer) CommitMessage(_a0 *kafka.Message) ([]kafka.TopicPartition, error) {
	ret := _m.Called(_a0)

	var r0 []kafka.TopicPartition
	if rf, ok := ret.Get(0).(func(*kafka.Message) []kafka.TopicPartition); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.TopicPartition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*kafka.Message) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Events provides a mock function with given fields:
func (_m *MockConsumer) Events() chan kafka.Event {
	ret := _m.Called()

	var r0 chan kafka.Event
	if rf, ok := ret.Get(0).(func() chan kafka.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan kafka.Event)
		}
	}

	return r0
}

// Poll provides a mock function with given fields: _a0
func (_m *MockConsumer) Poll(_a0 int) kafka.Event {
	ret := _m.Called(_a0)

	var r0 kafka.Event
	if rf, ok := ret.Get(0).(func(int) kafka.Event); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kafka.Event)
		}
	}

	return r0
}

// Subscribe provides a mock function with given fields: topic, rebalanceCb
func (_m *MockConsumer) Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error {
	ret := _m.Called(topic, rebalanceCb)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, kafka.RebalanceCb) error); ok {
		r0 = rf(topic, rebalanceCb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeTopics provides a mock function with given fields: topics, rebalanceCb
func (_m *MockConsumer) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error {
	ret := _m.Called(topics, rebalanceCb)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, kafka.RebalanceCb) error); ok {
		r0 = rf(topics, rebalanceCb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unassign provides a mock function with given fields:
func (_m *MockConsumer) Unassign() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
