/*
Copyright © 2026 Harry Sharma harrysharma1066@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	clientPort string // port number for client to connect to; should match launched serverPort(defaults to 6969).
	clientHost string // hostname for client to connect to; should match launched serverHost(defaults to 127.0.0.1).
	clientName string // identifier for client in frontend (duplicates solved through including first 6 digits of uuid) (defaults to anonymous).
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the websocket server",
	Long:  `Connect to the websocket server as a user client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	connectCmd.Flags().StringVarP(&clientPort, "port", "p", "", "server port to connect to")
	connectCmd.Flags().StringVarP(&clientHost, "hostname", "n", "", "server host name to connect to")
	connectCmd.Flags().StringVarP(&clientName, "clientname", "c", "", "client name they will be referred to")
}
