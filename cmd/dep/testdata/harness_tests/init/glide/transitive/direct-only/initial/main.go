// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/sdboyer/deptest"
)

func main() {
	var x deptest.Foo
	fmt.Println(x)
}