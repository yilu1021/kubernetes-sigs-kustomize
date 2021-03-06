// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package openapi

import (
	"io"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/cmd/config/configcobra"
	"sigs.k8s.io/kustomize/kustomize/v3/internal/commands/openapi/info"
)

// NewCmdOpenAPI makes a new openapi command.
func NewCmdOpenAPI(w io.Writer) *cobra.Command {

	openApiCmd := &cobra.Command{
		Use:     "openapi",
		Short:   "Commands for interacting with the OpenAPI data",
		Example: `kustomize openapi info`,
	}

	openApiCmd.AddCommand(info.NewCmdInfo(w))
	configcobra.AddCommands(openApiCmd, "openapi")

	return openApiCmd
}
