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
	"github.com/fsnotify/fsnotify"
	// "github.com/markbates/goth"
	// "github.com/kataras/iris"

	// "github.com/dgrijalva/jwt-go"
)

var VERSION string
var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "global-forward-auth",
	Short: "A forward authentication service",
	Long: `A forward authentication service that provides Google and Github OAuth based login and authentication for Traefik and Nginx reverse proxies. inspired by traefik-forward-auth.
Use case: use your organization Google email or Githhub user to have the option to authenticate/authorize to any application you run in your organization that doesn't support a native authentication and authorization.`,

	Run: func(cmd *cobra.Command, args []string) {
		level := cmd.Flag("log_level").Value.String()
		setLogLevel(level)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Global-Froward-Auth version information",
	Long:  `Print the version information of Global-Froward-Auth`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Global-Froward-Auth v%s (Go version: %s)\n", os.Getenv("VERSION"), runtime.Version())
	},
}

func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&configFile, "config", "c", "global-forward-auth", "path to config file")
	// rootCmd.Flags().StringP("engine", "e", "memory", "engine to use: memory or redis")
	rootCmd.Flags().StringP("log_level", "", "info", "set the log level: debug, info, error, fatal or none")
	// rootCmd.Flags().BoolP("prometheus", "", false, "enable Prometheus metrics endpoint")
	rootCmd.Flags().BoolP("insecure-cookie", "", false, "Use insecure cookies (default: false")
	rootCmd.Flags().BoolP("client-id", "", false, "Google/Github client ID")
	rootCmd.Flags().BoolP("client-secret", "", false, "Google/Github client secret")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("global-forward-auth")
	if configFile != "" {
		fmt.Println(">>> configFile: ", configFile)
		viper.SetConfigFile(configFile)
		configDir := path.Dir(configFile)
		if configDir != "." && configDir != dir {
			viper.AddConfigPath(configDir)
		}
	}

	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func setLogLevel(logLevel string) {
	loglevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(loglevel)
}
