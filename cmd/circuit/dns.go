// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2013 Petar Maymounkov <p@gocircuit.org>

package main

import (
	// "github.com/gocircuit/circuit/client"

	"github.com/gocircuit/circuit/github.com/codegangsta/cli"
)

func mkdns(x *cli.Context) {
	defer func() {
		if r := recover(); r != nil {
			fatalf("error, likely due to missing server or misspelled anchor: %v", r)
		}
	}()
	c := dial(x)
	args := x.Args()
	if len(args) != 1 {
		fatalf("mkdns needs an anchor argument")
	}
	w, _ := parseGlob(args[0])
	_, err := c.Walk(w).MakeNameserver()
	if err != nil {
		fatalf("mkdns error: %s", err)
	}
}