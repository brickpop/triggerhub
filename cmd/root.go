package cmd

import (
	"fmt"
	"os"

	"github.com/brickpop/triggerhub/config"
	"github.com/brickpop/triggerhub/services"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "triggerhub",
	Short: "Trigger Hub is a simple service that listens for trigger events on HTTP clients and relays them to subscribed services that may not want to expose a dedicated port.",
}

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start a dispatcher service",
		Long:  `Starts a dispatcher service, listenint for HTTP and relaying to WS registered services`,
		Run: func(cmd *cobra.Command, args []string) {
			services.Serve()
		},
	}
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "Joins a Trigger Hub server",
		Long:  `Registers to a Trigger Hub dispatcher service and waits for triggers to be reported`,
		Run: func(cmd *cobra.Command, args []string) {
			services.Listen()
		},
	}
	rootCmd.AddCommand(serveCmd, listenCmd)

	// Read flags
	rootCmd.PersistentFlags().String("config", "", "the config file to use")

	serveCmd.PersistentFlags().String("cert", "", "the certificate file (TLS only)")
	serveCmd.PersistentFlags().String("key", "", "the TLS encryption key file")
	serveCmd.PersistentFlags().Bool("tls", false, "whether to use TLS encryption (cert and key required)")
	serveCmd.PersistentFlags().IntP("port", "p", 8080, "port to bind to")

	// Get params
	cobra.OnInitialize(func() {
		config.DispatcherInit(rootCmd, serveCmd)
		config.ListenerInit(rootCmd)
	})
}

// Execute runs the cobra commands and parameters
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
