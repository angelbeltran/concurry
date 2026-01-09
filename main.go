/*
Copyright © 2026 Angel Beltran beltranbeaverdam@gmail.com
*/
package main

import (
	"context"
	"strings"

	"github.com/angelbeltran/concurry/cmd"
)

func main() {
	cmd.Execute()
}

type someType struct {
	A string
	B bool
	C int
}

func (st *someType) someMethod(ctx context.Context, x string) {
}

func (st *someType) someMethod2(ctx context.Context) (int, error) {
	return 1, nil
}

func (st *someType) someMethod3(ctx context.Context) *strings.Builder {
	return nil
}

func (st *someType) someMethod4(string) error {
	return nil
}
