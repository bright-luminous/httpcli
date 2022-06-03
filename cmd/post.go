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

var postCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "post",
	Short: "post request to url",
	Long:  `post request to url`,
	Run: func(cmd *cobra.Command, args []string) {
		finalUrl := args[0] + "/todos/"
		jsonData, _ := cmd.Flags().GetString("json")
		client := &http.Client{}

		fmt.Println(jsonData)

		jsonData1 := strings.Split(jsonData, "'")
		task := jsonData1[3]
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
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().String("json", "", "task to be add.")
}
