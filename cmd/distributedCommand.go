package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Send a new message from CLI to UI",
	Long:  `This command sends a new message from the CLI to the UI submodule.`,
	Run: func(cmd *cobra.Command, args []string) {
		sendDataToUI()
	},
}

func sendDataToUI() {
	// Sending a POST request to the UI server
	url := "http://localhost:8080/data"
	var jsonStr = []byte(`{"message":"Hello from CLI"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Sent data to UI")
}

func init() {
	rootCmd.AddCommand(newCmd)
}
