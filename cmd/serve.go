/*
Copyright © 2026 Harry Sharma harrysharma1066@gmail.com
*/
package cmd

import (
	"bubchat/server"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if versionCheck {
			fmt.Printf("%s (%s)", VERSION_NAME, VERSION_NUMBER)
			return nil
		}
		if serverHost == "" {
			serverHost = "127.0.0.1"
		}

		if serverPort == "" {
			serverPort = "6969"
		}
		url := fmt.Sprintf("%s:%s", serverHost, serverPort)

		hub := server.NewHub()
		go hub.Run()

		mux := http.NewServeMux()
		mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			server.ServeWS(hub, w, r)
		})

		server := &http.Server{
			Addr:    url,
			Handler: mux,
		}

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT)
		go func() {
			log.Infof("Websocket server launched on ws::/%s/ws", url)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
		<-stop
		log.Info("Shutting down server gracefully...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Error("Server shutdown failed:", err)
		}
		log.Info("Server exited cleanly 👋")
		return nil
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
