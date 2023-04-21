/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my test cli",
	Short: "this is just a test cli",
	Long:  `this is just a test clik long long long `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		ahch, err := cmd.Flags().GetBool("amir")
		marat, err := cmd.Flags().GetBool("marat")
		manas, err := cmd.Flags().GetBool("manas")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if ahch {
			fmt.Println("react native is my only love")
			return
		}
		if marat {
			fmt.Println("ping pong")
			return
		}
		if manas {
			fmt.Println("(:")
			return
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("ahch", "a", false, "for ahch")
	rootCmd.Flags().BoolP("marat", "m", false, "for ahch")
	rootCmd.Flags().BoolP("manas", "k", false, "for ahch")
}
