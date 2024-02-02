// cmd/auth.go
package cmd

import (
	"bufio"
	"fmt"
	"github.com/dalekurt/proxmox-cli/proxmox"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the Proxmox server",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter Password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)
		fmt.Println()

		client := proxmox.NewClient(os.Getenv("PROXMOX_BASE_URL"))
		if err := client.Authenticate(username, password, ""); err != nil {
			return fmt.Errorf("authentication failed: %v", err)
		}

		fmt.Println("Authenticated successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
