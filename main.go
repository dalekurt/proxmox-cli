// main.go
package main

import (
	"fmt"
	"github.com/dalekurt/proxmox-cli/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
