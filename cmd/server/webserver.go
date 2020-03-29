package server

import (
	"github.com/spf13/cobra"
	s "../../pkg/server"
)

func ServerCmd() *cobra.Command {
	return &cobra.Command {
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.Serve()
		},
	}
}
