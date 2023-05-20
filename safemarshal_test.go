package safemarshal_test

import (
	"encoding/json"
	"testing"

	"github.com/merlincox/safemarshal"
	"github.com/stretchr/testify/assert"
)

type simple struct {
	Field1 string
}

type complicated struct {
	Field1 string
	Field2 any
}

func TestString(t *testing.T) {
	v := "a string"
	testType(t, v, true)
}

func TestStringPtr(t *testing.T) {
	v := "a string"
	testType(t, &v, true)
}

func TestInt(t *testing.T) {
	v := 9
	testType(t, v, true)
}

func TestIntPtr(t *testing.T) {
	v := "a string"
	testType(t, &v, true)
}

func TestSimple(t *testing.T) {
	v := simple{
		Field1: "field1",
	}
	testType(t, v, true)
}

func TestComplicated(t *testing.T) {
	v := complicated{
		Field1: "field1",
		Field2: 99,
	}
	testType(t, v, false)
}

func TestComplicatedSlice(t *testing.T) {
	v1 := complicated{
		Field1: "field1",
		Field2: 99,
	}
	v := []complicated{v1}
	testType(t, v, false)
}

func TestComplicatedMap(t *testing.T) {
	v1 := complicated{
		Field1: "field1",
		Field2: 99,
	}
	v := map[string]complicated{
		"key": v1,
	}
	testType(t, v, false)
}

func TestSimpleSlice(t *testing.T) {
	v1 := simple{
		Field1: "field1",
	}
	v := []simple{v1}
	testType(t, v, true)
}

func testType(t *testing.T, v any, expectedSafe bool) {
	checked := safemarshal.Check(v)
	assert.Equal(t, expectedSafe, checked, "Check for %T should return %v", v, expectedSafe)
	_, err := json.Marshal(v)
	if expectedSafe {
		assert.Nil(t, err)
	}
}
