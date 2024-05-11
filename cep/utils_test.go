package cep_test

import (
	"testing"

	"github.com/agaragon/brutils-go/cep"
)

var tablesValidate = []struct {
	input    string
	expected bool
}{
	{"000111", false},
	{"13723705000189", false},
	{"60.391.947/0001-0", false},
	{"abcdefgh", false},
	{"99999999999999", false},
	{"12345678", true},
	{"12345-678", true},
}

var tablesClean = []struct {
	input    string
	expected string
}{
	{"12345-678", "12345678"},
	{"00000-000", "00000000"},
	{"11111-111", "11111111"},
}

func TestValidate(t *testing.T) {
	for _, table := range tablesValidate {
		if res := cep.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := cep.Generate()
		if !cep.IsValid(res) {
			t.Errorf("An invalid cep was generated: %s", res)
		}
	}
}

func TestClean(t *testing.T) {
	for _, row := range tablesClean {
		if res := cep.Clean(row.input); res != row.expected {
			t.Errorf("Clean failed for %s \t Expected: %s | Received: %s", row.input, row.expected, res)
		}
	}
}