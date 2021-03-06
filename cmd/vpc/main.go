package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sean-/conswriter"
	"github.com/sean-/vpc/cli"
	"github.com/sean-/vpc/internal/buildtime"
)

var (
	// Variables populated by govvv(1).
	Version    = "dev"
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
)

func main() {
	exportBuildtimeConsts()

	defer func() {
		p := conswriter.GetTerminal()
		p.Wait()
	}()

	if err := cli.Execute(); err != nil {
		log.Error().Err(err).Msg("unable to run")
		os.Exit(1)
	}
}

func exportBuildtimeConsts() {
	buildtime.GitCommit = GitCommit
	buildtime.GitBranch = GitBranch
	buildtime.GitState = GitState
	buildtime.GitSummary = GitSummary
	buildtime.BuildDate = BuildDate
	buildtime.Version = Version
}
