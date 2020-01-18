package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SetupMockup() Position {
	var p Position
	p.X = 1
	p.Y = 1
	p.Direction = "N"
	return p
}

func TestRepo_ForwardSuccessWithNCondition(t *testing.T) {
	mockupData := SetupMockup()
	testService, _ := New(mockupData)
	_ = testService.Forward(&mockupData)
	assert.Equal(t, 2, mockupData.Y)
}

func TestRepo_ForwardSuccessWithWCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "W"
	testService, _ := New(mockupData)
	_ = testService.Forward(&mockupData)
	assert.Equal(t, 0, mockupData.X)
}

func TestRepo_ForwardSuccessWithSCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "S"
	testService, _ := New(mockupData)
	_ = testService.Forward(&mockupData)
	assert.Equal(t, 0, mockupData.Y)
}

func TestRepo_ForwardSuccessWithECondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "E"
	testService, _ := New(mockupData)
	_ = testService.Forward(&mockupData)
	assert.Equal(t, 2, mockupData.X)
}

func TestRepo_ForwardWrongFormat(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "e"
	testService, _ := New(mockupData)
	err := testService.Forward(&mockupData)
	assert.Errorf(t, err, "Wrong Format")
}

func TestRepo_LeftSuccessWithNCondition(t *testing.T) {
	mockupData := SetupMockup()
	testService, _ := New(mockupData)
	_ = testService.Left(&mockupData)
	assert.Equal(t, "W", mockupData.Direction)
}

func TestRepo_LeftSuccessWithWCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "W"
	testService, _ := New(mockupData)
	_ = testService.Left(&mockupData)
	assert.Equal(t, "S", mockupData.Direction)
}

func TestRepo_LeftSuccessWithSCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "S"
	testService, _ := New(mockupData)
	_ = testService.Left(&mockupData)
	assert.Equal(t, "E", mockupData.Direction)
}

func TestRepo_LeftSuccessWithECondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "E"
	testService, _ := New(mockupData)
	_ = testService.Left(&mockupData)
	assert.Equal(t, "N", mockupData.Direction)
}

func TestRepo_LeftWrongFormat(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "e"
	testService, _ := New(mockupData)
	err := testService.Left(&mockupData)
	assert.Errorf(t, err, "Wrong Format")
}

func TestRepo_RightSuccessNCondition(t *testing.T) {
	mockupData := SetupMockup()
	testService, _ := New(mockupData)
	_ = testService.Right(&mockupData)
	assert.Equal(t, "E", mockupData.Direction)
}

func TestRepo_RightSuccessWithWCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "W"
	testService, _ := New(mockupData)
	_ = testService.Right(&mockupData)
	assert.Equal(t, "N", mockupData.Direction)
}

func TestRepo_RightSuccessWithSCondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "S"
	testService, _ := New(mockupData)
	_ = testService.Right(&mockupData)
	assert.Equal(t, "W", mockupData.Direction)
}

func TestRepo_RightSuccessWithECondition(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "E"
	testService, _ := New(mockupData)
	_ = testService.Right(&mockupData)
	assert.Equal(t, "S", mockupData.Direction)
}

func TestRepo_RightWrongFormat(t *testing.T) {
	mockupData := SetupMockup()
	mockupData.Direction = "e"
	testService, _ := New(mockupData)
	err := testService.Right(&mockupData)
	assert.Errorf(t, err, "Wrong Format")
}
