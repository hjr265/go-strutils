// Copyright 2014 The Go-strutils Authors. All rights reserved.

package z_test

import (
	"fmt"

	"github.com/hjr265/go-strutils/z"
)

func ExampleZ() {
	fmt.Printf("%v", z.New("aabaabaaaabbaabaabaaabbaabaabb"))

	// Output: [0 1 0 5 1 0 2 2 3 1 0 0 9 1 0 5 1 0 2 3 1 0 0 6 1 0 3 1 0 0]
}

func ExampleContains() {
	fmt.Println(z.Contains("seafood", "foo"))
	fmt.Println(z.Contains("seafood", "bar"))
	fmt.Println(z.Contains("seafood", ""))
	fmt.Println(z.Contains("", ""))

	// Output:
	// true
	// false
	// true
	// true
}

func ExampleIndex() {
	fmt.Println(z.Index("chicken", "ken"))
	fmt.Println(z.Index("chicken", "dmr"))

	// Output:
	// 4
	// -1
}
