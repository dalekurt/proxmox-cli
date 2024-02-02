// cmd/status.go
package cmd

import (
	"fmt"
	"github.com/dalekurt/proxmox-cli/proxmox"
	"github.com/spf13/cobra"
	"os"
)

// statusCmd
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the Proxmox server",
	Long:  `The status command checks to confirm that the Proxmox server is up and running.`,
	Run: func(cmd *cobra.Command, args []string) {
		baseURL := os.Getenv("PROXMOX_BASE_URL")
		client := proxmox.NewClient(baseURL)
		err := client.CheckStatus()
		if err != nil {
			fmt.Printf("Failed to check Proxmox server status: %v\n", err)
		} else {
			fmt.Println("Proxmox server is up and running.")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
