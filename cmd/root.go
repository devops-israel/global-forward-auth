package cmd

import (
	"github.com/spf13/cobra"
    "github.com/markbates/goth"
    "github.com/kataras/iris"
    "github.com/spf13/viper"
	"github.com/dgrijalva/jwt-go"
)


// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Create the web service with the configuration values",
	Long: `This application is simple program to learn spf13/cobra
spf13/cobra looks really nice CLI framework.
I want to create lovely CLI program with this framework :)
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}