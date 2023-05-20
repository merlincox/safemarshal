package safemarshal_test

import (
	"encoding/json"
	"testing"

	"github.com/merlincox/safemarshal"
	"github.com/stretchr/testify/assert"
)

type safeStruct struct {
	Field1 string
}

type safeRecursiveStruct struct {
	Field1 string
	Field2 *safeRecursiveStruct
}

type safeLinkedRecursiveStruct1 struct {
	Field1 string
	Field2 *safeLinkedRecursiveStruct2
}

type safeLinkedRecursiveStruct2 struct {
	Field1 string
	Field2 *safeLinkedRecursiveStruct1
}

type unsafeLinkedRecursiveStruct1 struct {
	Field1 string
	Field2 *unsafeLinkedRecursiveStruct2
}

type unsafeLinkedRecursiveStruct2 struct {
	Field1 string
	Field2 *unsafeLinkedRecursiveStruct1
	Field3 any
}

type unsafeStruct1 struct {
	Field1 string
	Field2 any
}

type unsafeStruct2 struct {
	Field1 string
	Field2 func()
}

type unsafeStruct3 struct {
	Field1 string
	Field2 chan int
}

type unsafeRecursiveStruct struct {
	Field1 string
	Field2 *unsafeRecursiveStruct
	Field3 any
}

type testcase struct {
	name      string
	subject   any
	expectsOK bool
}

func TestOK(t *testing.T) {
	var (
		strVal        = "a string"
		intVal        = 999
		uintVal       = 999
		floatVal      = 999.99
		bytesVal      = []byte("bytes")
		charVal       = 'x'
		safeStructVal = safeStruct{
			Field1: "a string",
		}
		safeRecursiveStructVal = safeRecursiveStruct{
			Field1: "a string",
			Field2: &safeRecursiveStruct{
				Field1: "another string",
			},
		}
		safeLinkedRecursiveStructVal = safeLinkedRecursiveStruct1{
			Field1: "a string",
			Field2: &safeLinkedRecursiveStruct2{
				Field1: "another string",
				Field2: &safeLinkedRecursiveStruct1{
					Field1: "yet another string",
				},
			},
		}
		unsafeLinkedRecursiveStructVal = unsafeLinkedRecursiveStruct1{
			Field1: "a string",
			Field2: &unsafeLinkedRecursiveStruct2{
				Field1: "another string",
				Field2: &unsafeLinkedRecursiveStruct1{
					Field1: "yet another string",
				},
				Field3: 99.99,
			},
		}
		unsafeRecursiveStructVal = unsafeRecursiveStruct{
			Field1: "a string",
			Field2: &unsafeRecursiveStruct{
				Field1: "another string",
			},
			Field3: 99.9,
		}
		unsafeStruct1Val = unsafeStruct1{
			Field1: "a string",
			Field2: 999,
		}
		unsafeStruct2Val = unsafeStruct2{
			Field1: "a string",
			Field2: func() {},
		}
		unsafeStruct3Val = unsafeStruct3{
			Field1: "a string",
			Field2: make(chan int),
		}
		complexVal      = complex(10, 10)
		nilSafeStruct   *safeStruct
		nilUnsafeStruct *unsafeStruct1
	)
	testcases := []testcase{
		{
			name:      "nil",
			subject:   nil,
			expectsOK: true,
		},
		{
			name:      "nil safe struct",
			subject:   nilSafeStruct,
			expectsOK: true,
		},
		{
			name:      "nil unsafe struct",
			subject:   nilUnsafeStruct,
			expectsOK: false,
		},
		{
			name:      "string",
			subject:   strVal,
			expectsOK: true,
		},
		{
			name:      "string pointer",
			subject:   &strVal,
			expectsOK: true,
		},
		{
			name:      "char",
			subject:   charVal,
			expectsOK: true,
		},
		{
			name:      "char pointer",
			subject:   &charVal,
			expectsOK: true,
		},
		{
			name:      "bytes",
			subject:   bytesVal,
			expectsOK: true,
		},
		{
			name:      "int",
			subject:   intVal,
			expectsOK: true,
		},
		{
			name:      "int pointer",
			subject:   &intVal,
			expectsOK: true,
		},
		{
			name:      "uint",
			subject:   uintVal,
			expectsOK: true,
		},
		{
			name:      "uint pointer",
			subject:   &uintVal,
			expectsOK: true,
		},
		{
			name:      "float",
			subject:   floatVal,
			expectsOK: true,
		},
		{
			name:      "float pointer",
			subject:   &floatVal,
			expectsOK: true,
		},
		{
			name: "anonymous safe struct",
			subject: struct {
				Field string
			}{
				Field: "string",
			},
			expectsOK: true,
		},
		{
			name: "anonymous unsafe struct",
			subject: struct {
				Field any
			}{
				Field: "string",
			},
			expectsOK: false,
		},
		{
			name:      "safe struct",
			subject:   safeStructVal,
			expectsOK: true,
		},
		{
			name:      "safe struct pointer",
			subject:   &safeStructVal,
			expectsOK: true,
		},
		{
			name:      "safe struct slice",
			subject:   []safeStruct{safeStructVal},
			expectsOK: true,
		},
		{
			name:      "safe struct array",
			subject:   [1]safeStruct{safeStructVal},
			expectsOK: true,
		},
		{
			name:      "safe struct map",
			subject:   map[string]safeStruct{"key": safeStructVal},
			expectsOK: true,
		},
		{
			name:      "safe recursive struct",
			subject:   safeRecursiveStructVal,
			expectsOK: true,
		},
		{
			name:      "safe linked recursive struct",
			subject:   safeLinkedRecursiveStructVal,
			expectsOK: true,
		},
		{
			name:      "unsafe recursive struct",
			subject:   unsafeRecursiveStructVal,
			expectsOK: false,
		},
		{
			name:      "unsafe linked recursive struct",
			subject:   unsafeLinkedRecursiveStructVal,
			expectsOK: false,
		},
		{
			name:      "unsafe struct 1",
			subject:   unsafeStruct1Val,
			expectsOK: false,
		},
		{
			name:      "unsafe struct 2",
			subject:   unsafeStruct2Val,
			expectsOK: false,
		},
		{
			name:      "unsafe struct 3",
			subject:   unsafeStruct3Val,
			expectsOK: false,
		},
		{
			name:      "unsafe struct pointer",
			subject:   &unsafeStruct1Val,
			expectsOK: false,
		},
		{
			name:      "unsafe struct slice",
			subject:   []unsafeStruct1{unsafeStruct1Val},
			expectsOK: false,
		},
		{
			name:      "unsafe struct array",
			subject:   [1]unsafeStruct1{unsafeStruct1Val},
			expectsOK: false,
		},
		{
			name:      "unsafe struct map",
			subject:   map[string]unsafeStruct1{"key": unsafeStruct1Val},
			expectsOK: false,
		},
		{
			name:      "complex number",
			subject:   complexVal,
			expectsOK: false,
		},
		{
			name:      "unsafe map key",
			subject:   map[complex128]safeStruct{complexVal: safeStructVal},
			expectsOK: false,
		},
	}

	for _, tc := range testcases {
		test := tc
		t.Run(
			test.name, func(t *testing.T) {
				t.Parallel()
				ok := safemarshal.OK(test.subject)
				assert.Equal(t, test.expectsOK, ok, "OK for %T should return %v", test.subject, test.expectsOK)
				_, err := json.Marshal(test.subject)
				if test.expectsOK {
					assert.Nil(t, err, "Marshalling %T should not throw an error", test.subject)
				}
			},
		)
	}
}
