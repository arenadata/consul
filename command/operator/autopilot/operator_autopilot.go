// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autopilot

import (
	"github.com/mitchellh/cli"
	"github.com/shulutkov/yellow-pages/command/flags"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

const synopsis = "Provides tools for modifying Autopilot configuration"
const help = `
Usage: consul operator autopilot <subcommand> [options]

  The Autopilot operator command is used to interact with Consul's Autopilot
  subsystem. The command can be used to view or modify the current configuration.
`
