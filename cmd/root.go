package cmd

import (
	"os"
	"fmt"
	"runtime"
	"path"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	// "github.com/markbates/goth"
	// "github.com/kataras/iris"

	// "github.com/dgrijalva/jwt-go"
)

var VERSION string
var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A forward authentication service",
	Long: `A forward authentication service that provides Google and Github OAuth based login and authentication for Traefik and Nginx reverse proxies. inspired by traefik-forward-auth.
Use case: use your organization Google email or Githhub user to have the option to authenticate/authorize to any application you run in your organization that doesn't support a native authentication and authorization.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Global-Froward-Auth version information",
	Long:  `Print the version information of Global-Froward-Auth`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Global-Froward-Auth v%s (Go version: %s)\n", os.Getenv("VERSION"), runtime.Version())
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Global-Froward-Auth application",
	Long:  `run Global-Froward-Auth`,
	Run: func(cmd *cobra.Command, args []string) {
		level := viper.GetString("log_level")
		setLogLevel(level)
	},
}

func Execute() {
	cobra.OnInitialize(initConfig)

	runCmd.Flags().StringVarP(&configFile, "config", "c", "", "path to config file e.g /path/to/config/config.yaml")
	runCmd.Flags().StringP("log_level", "", "info", "set the log level: debug, info, error, fatal or none")
	// runCmd.Flags().BoolP("prometheus", "", false, "enable Prometheus metrics endpoint")
	runCmd.Flags().BoolP("insecure-cookie", "", false, "Use insecure cookies (default: false")
	runCmd.Flags().BoolP("client-id", "", false, "Google/Github client ID")
	runCmd.Flags().BoolP("client-secret", "", false, "Google/Github client secret")

	// runCmd.Flags().Parse()
	// viper.BindPFlags(runCmd.Flags().CommandLine)
	viper.BindPFlags(runCmd.Flags())

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig(){
	viper.SetEnvPrefix("GFA")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if configFile != "" {
		viper.SetConfigName(configFile)
		fmt.Println(">>> configFile: ", configFile)
		viper.SetConfigFile(configFile)
		configDir := path.Dir(configFile)
		if configDir != "." && configDir != dir {
			viper.AddConfigPath(configDir)
		}

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			fmt.Println(err)
		}
	}

	viper.AutomaticEnv()
}

func setLogLevel(logLevel string) {
	loglevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(loglevel)
	fmt.Println("Using log level:", loglevel)
}
