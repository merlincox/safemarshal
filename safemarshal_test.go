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

type selfStruct struct {
	Field1 string
	Field2 *selfStruct
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

type unsafeSelfStruct struct {
	Field1 string
	Field2 *unsafeSelfStruct
	Field3 any
}

type testcase struct {
	name         string
	subject      any
	expectChecks bool
}

func TestCheck(t *testing.T) {
	var (
		strVal        = "a string"
		intVal        = 999
		bytesVal      = []byte("bytes")
		safeStructVal = safeStruct{
			Field1: "a string",
		}
		selfStructVal = selfStruct{
			Field1: "a string",
			Field2: &selfStruct{Field1: "another string"},
		}
		unsafeSelfStructVal = unsafeSelfStruct{
			Field1: "a string",
			Field2: &unsafeSelfStruct{Field1: "another string"},
			Field3: 99.9,
		}
		unsafeStructVal1 = unsafeStruct1{
			Field1: "a string",
			Field2: 999,
		}
		unsafeStructVal2 = unsafeStruct2{
			Field1: "a string",
			Field2: func() {},
		}
		unsafeStructVal3 = unsafeStruct3{
			Field1: "a string",
			Field2: make(chan int),
		}

		complexVal = complex(10, 10)
	)
	testcases := []testcase{
		{
			name:         "string",
			subject:      strVal,
			expectChecks: true,
		},
		{
			name:         "string pointer",
			subject:      &strVal,
			expectChecks: true,
		},
		{
			name:         "bytes",
			subject:      bytesVal,
			expectChecks: true,
		},
		{
			name:         "int",
			subject:      intVal,
			expectChecks: true,
		},
		{
			name:         "int pointer",
			subject:      &intVal,
			expectChecks: true,
		},
		{
			name:         "safe struct",
			subject:      safeStructVal,
			expectChecks: true,
		},
		{
			name:         "safe struct pointer",
			subject:      &safeStructVal,
			expectChecks: true,
		},
		{
			name:         "safe struct slice",
			subject:      []safeStruct{safeStructVal},
			expectChecks: true,
		},
		{
			name:         "safe struct map",
			subject:      map[string]safeStruct{"key": safeStructVal},
			expectChecks: true,
		},
		{
			name:         "self struct",
			subject:      selfStructVal,
			expectChecks: true,
		},
		{
			name:         "unsafe self struct",
			subject:      unsafeSelfStructVal,
			expectChecks: false,
		},
		{
			name:         "unsafe struct 1",
			subject:      unsafeStructVal1,
			expectChecks: false,
		},
		{
			name:         "unsafe struct 2",
			subject:      unsafeStructVal2,
			expectChecks: false,
		},
		{
			name:         "unsafe struct 3",
			subject:      unsafeStructVal3,
			expectChecks: false,
		},
		{
			name:         "unsafe struct pointer",
			subject:      &unsafeStructVal1,
			expectChecks: false,
		},
		{
			name:         "complex number",
			subject:      complexVal,
			expectChecks: false,
		},
	}
	for _, tc := range testcases {
		t.Run(
			tc.name, func(t *testing.T) {
				checks := safemarshal.Check(tc.subject)
				assert.Equal(t, tc.expectChecks, checks, "Check for %T should return %v", tc.subject, tc.expectChecks)
				_, err := json.Marshal(tc.subject)
				if tc.expectChecks {
					assert.Nil(t, err, "Marshalling %T should not throw an error", tc.subject)
				}
			},
		)
	}
}
