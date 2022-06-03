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
	"strings"

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
		finalUrl := args[0] + "/todos/"
		client := &http.Client{}
		jsonData, _ := cmd.Flags().GetString("json")

		fmt.Println(jsonData)

		jsonData1 := strings.Split(jsonData, "'")

		fmt.Println(jsonData1)
		fmt.Println(jsonData1[3])
		task := jsonData1[3]
		fmt.Println(jsonData1[7])
		description := jsonData1[7]

		requestBody, err := json.Marshal(map[string]string{
			"task":        task,
			"description": description,
		})
		if err != nil {
			log.Fatalln(err)
		}

		if len(queryArr) > 0 {
			finalUrl = finalUrl + "?"
			for i := range queryArr {
				finalUrl = finalUrl + queryArr[i]
				if i+1 < len(queryArr) {
					finalUrl = finalUrl + "&"
				}
			}
		}
		req, err := http.NewRequest("POST", finalUrl, bytes.NewBuffer(requestBody))
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

		resp, err := client.Do(req)
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
	postCmd.PersistentFlags().String("json", "", "task to be add.")
}
