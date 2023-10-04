package cmd

import (
	"log"
	"os"
	"requestlogger/src/sniffer"
	"strings"

	"github.com/spf13/cobra"
)

var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "requestlogger",
	Short: "Log received requests in the specified ports",
	Long: `Http sniffer for logging all requests made in the specified ports or/and URLs 
	default port is 8080`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		iface, _ := cmd.Flags().GetString("iface")
		port, _ := cmd.Flags().GetString("port")

		if strings.TrimSpace(port) == "" {
			port = "8080"
		}

		if iface != "" {
			sniffer.StartSniffer(sniffer.NewSniffer(iface, port))
		} else {
			log.Fatal("You need to specify an interface, ex: requestlogger --iface=xxxxx")
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.requestlogger.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().String("iface", "", "Network interface where you want to log requests from")
	rootCmd.PersistentFlags().String("port", "", "Port to listen")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
