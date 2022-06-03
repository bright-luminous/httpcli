/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var queryArr []string
var headerArr []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "root",
	Short: "sent GET request to URL",
	Long:  `sent GET request to URL`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		client := &http.Client{}

		finalUrl := args[0]
		if len(queryArr) > 0 {
			finalUrl = finalUrl + "?"
			for i := range queryArr {
				finalUrl = finalUrl + queryArr[i]
				if i+1 < len(queryArr) {
					finalUrl = finalUrl + "&"
				}
			}
		}

		fmt.Println(finalUrl)

		req, err := http.NewRequest("GET", finalUrl, nil)
		if err != nil {
			log.Fatalln(err)
		}
		if len(headerArr) > 0 {
			for i := range headerArr {
				headerToAdd := strings.Split(headerArr[i], "=")
				if len(headerToAdd) > 1 {
					req.Header.Add(headerToAdd[0], headerToAdd[1])
				} else {
					log.Fatalln("wrong header flag")
				}
			}
		}
		// client.Head()
		fmt.Println(req)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Print(sb)
		fmt.Println("get called root")
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httpcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.PersistentFlags().String("header", "", "add this key and value into header")
	rootCmd.PersistentFlags().StringArrayVarP(&headerArr, "header", "", []string{}, "add this key and value into header")
}
