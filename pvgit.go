package main

import "github.com/spf13/cobra"

func main() {
	root := &cobra.Command{Use: "pvgit"}
	root.AddCommand(postCommitHookCommand())
	root.AddCommand(postCheckoutHookCommand())
	root.AddCommand(syncCommand())
	root.AddCommand(instrumentCommand())

	root.Execute()
}
