/*
Copyright © 2026 Harry Sharma harrysharma1066@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	serverPort string // port number for server (defaults to 6969)
	serverHost string // hostname for server (defaults to 127.0.0.1)
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving the websocket server",
	Long:  `Serving the websocket server for clients to connect to.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringVarP(&serverPort, "port", "p", "", "port to host websocket server on")
	serveCmd.Flags().StringVarP(&serverHost, "hostname", "n", "", "hostname to host websocket server on")
}
