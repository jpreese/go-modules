package main

import (
	"github.com/spf13/cobra"
	"github.com/jpreese/importme/noimport"
)

var MainUseCommand cobra.Command

func main() {
	noimport.Foo()
}
