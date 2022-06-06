package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().String("json", "", "task to be add.")
}

var postCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "post",
	Short: "post request to url",
	Long:  `post request to url`,
	Run: func(cmd *cobra.Command, args []string) {
		finalUrl := args[0] + "/todos/"
		jsonData, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatalln(err)
		}
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

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

		finalUrl = urlAddQuery(finalUrl)
		req, err := http.NewRequest("POST", finalUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalln(err)
		}
		req = reqAddHeader(headerParameters, req)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		strBody := string(body)
		log.Print(strBody)
	},
}
