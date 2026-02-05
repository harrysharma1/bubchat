package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUUIDGenerator struct {
}

func (m *MockUUIDGenerator) GenerateID() string {
	return "8a189573-e3ab-4b5e-a54d-42a9c4fe4af5"
}

func TestFirstN(t *testing.T) {
	initString := "abcdefghijklmnopqrstuvwxyz"

	expectedN5String := "abcde"
	gotN5String := FirstN(initString, 5)
	expectedN10String := "abcdefghij"
	gotN10String := FirstN(initString, 10)
	expectedN15String := "abcdefghijklmno"
	gotN15String := FirstN(initString, 15)

	assert.Equal(t, expectedN5String, gotN5String)
	assert.Equal(t, expectedN10String, gotN10String)
	assert.Equal(t, expectedN15String, gotN15String)
}

func TestHexFromUUID(t *testing.T) {
	mockUUIDGenerator := MockUUIDGenerator{}
	expected := "#8a1895"
	got := HexFromUUID(FirstN(mockUUIDGenerator.GenerateID(), 6))

	assert.Equal(t, expected, got)
}
