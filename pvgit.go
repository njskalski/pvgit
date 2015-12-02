package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	root := &cobra.Command{Use: "pvgit"}
	root.AddCommand(postCommitHookCommand())
	root.AddCommand(postCheckoutHookCommand())
	root.AddCommand(syncCommand())
	root.AddCommand(instrumentCommand())

	var vflag bool
	root.PersistentFlags().BoolVarP(&vflag, "version", "v", false, "Print version")
	root.ParseFlags(os.Args)
	if vflag {
		fmt.Println("pvgit", "version", version)
		os.Exit(0)
	}

	root.Execute()
}
