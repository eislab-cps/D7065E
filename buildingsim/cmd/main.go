package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	buildingsim "github.com/eislab-cps/buildingsim"
	"github.com/eislab-cps/buildingsim/pkg/server"
)

var rootCmd = &cobra.Command{
	Use:   "buildsim",
	Short: "Building simulation server",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the BuildSim server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetInt("port")
		edit, _ := cmd.Flags().GetBool("edit")
		srv := server.New(port, buildingsim.DataFS, buildingsim.WebFS, edit)
		return srv.Start()
	},
}

func init() {
	startCmd.Flags().IntP("port", "p", 9090, "Port to listen on")
	startCmd.Flags().Bool("edit", false, "Enable floor plan editing tools")
	rootCmd.AddCommand(startCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
