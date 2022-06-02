/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var postCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "post",
	Short: "post request to url",
	Long:  `post request to url`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		task, _ := cmd.Flags().GetString("json-task")
		description, _ := cmd.Flags().GetString("json-description")

		requestBody, err := json.Marshal(map[string]string{
			"task":        task,
			"description": description,
		})
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post(args[0], "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
		fmt.Println("get called post")
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httpcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(postCmd)
	rootCmd.PersistentFlags().String("json-task", "", "task to be add.")
	rootCmd.PersistentFlags().String("json-description", "", "description for the task.")
}
