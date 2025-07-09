package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFastGOCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "fg-apiserver",
		Short:        "A very lightweight full go project",
		Long:         `A very lightweight full go project, designed to help beginners quickly learn Go project development.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello FastGO!")
			return nil
		},
		Args: cobra.NoArgs,
	}

	return cmd
}
