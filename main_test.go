package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectCommandFirstScenarioSuccess(t *testing.T) {
	commandMockup := "LMLMLMLMM"
	posMockup := Position{
		X:         1,
		Y:         2,
		Direction: "N",
	}

	expectedResult := Position{
		X:         1,
		Y:         3,
		Direction: "N",
	}

	_ = SelectCommand(commandMockup, &posMockup)
	assert.Equal(t, expectedResult, posMockup)
}

func TestSelectCommandSecondScenarioSuccess(t *testing.T) {
	commandMockup := "MMRMMRMRRM"
	posMockup := Position{
		X:         3,
		Y:         3,
		Direction: "E",
	}

	expectedResult := Position{
		X:         5,
		Y:         1,
		Direction: "E",
	}

	_ = SelectCommand(commandMockup, &posMockup)
	assert.Equal(t, expectedResult, posMockup)
}
