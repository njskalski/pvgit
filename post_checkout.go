package main

import "github.com/spf13/cobra"

func postCheckoutHookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "hook-post-checkout",
		Short:  "Processes a git post-checkout hook event.",
		Run:    runPostCheckout,
		Hidden: true,
	}

	return cmd
}

func runPostCheckout(cmd *cobra.Command, args []string) {
	if args[2] == "0" {
		// if flag at third arg is zero, it means it's a file checkout; we do nothing
		return
	}

	repo := getRepoOrExit()
	m := newMessage()
	recordHead(m, repo)
	sendMapToPipeviz(m, repo)
}
