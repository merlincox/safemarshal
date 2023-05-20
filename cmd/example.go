package main

import (
	"encoding/json"
	"fmt"

	"github.com/merlincox/safemarshal"
)

type Safe struct {
	Field string
}

type Unsafe struct {
	Field any
}

func main() {

	ok := safemarshal.OK(Safe{})

	// This will print 'true'.
	fmt.Println(ok)

	ok = safemarshal.OK(Unsafe{})

	// This will print 'false'.
	fmt.Println(ok)

	// Note that OK returning false does *not* necessarily indicate that JSON marshalling will fail.
	// That may depend on runtime values.

	if _, err := json.Marshal(Unsafe{Field: "safe"}); err != nil {
		// Nothing will print.
		fmt.Println(err)
	}

	if _, err := json.Marshal(Unsafe{Field: func() {}}); err != nil {
		// This will print 'json: unsupported type: func()'
		fmt.Println(err)
	}
}
