// cmd/root.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "Proxmox CLI interacts with the Proxmox API",
	Long:  `A simple command-line tool built with Go to interact with the Proxmox API for managing VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement the root command
		fmt.Println("Proxmox CLI is running")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// TODO: Initialize any flags or configuration here
}
