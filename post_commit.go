package main

import (
	"log"

	"github.com/pipeviz/pipeviz/ingest"
	"github.com/spf13/cobra"
)

func postCommitHookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "hook-post-commit",
		Short:  "Processes a git post-commit hook event.",
		Run:    runPostCommit,
		Hidden: true,
	}

	return cmd
}

func runPostCommit(cmd *cobra.Command, args []string) {
	repo := getRepoOrExit()

	head, err := repo.Head()
	if err != nil {
		log.Fatalf("Could not get repo HEAD")
	}

	commit, err := repo.LookupCommit(head.Target())
	if err != nil {
		log.Fatalf("Could not find commit pointed at by HEAD")
	}

	ident, err := GetRepoIdent(repo)
	if err != nil {
		log.Fatalln("Failed to retrieve identifier for repository")
	}

	m := new(ingest.Message)
	m.Add(commitToSemanticForm(commit, ident))

	recordHead(m, repo)
	sendMapToPipeviz(m, repo)
}
