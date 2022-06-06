package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var queryParameters []string
var headerParameters []string

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringSliceVarP(&headerParameters, "header", "", []string{}, "add this key and value into header")
}

func urlAddQuery(finalUrl string) string {
	if len(queryParameters) > 0 {
		finalUrl = finalUrl + "?"
		for i := range queryParameters {
			finalUrl = finalUrl + queryParameters[i]
			if i+1 < len(queryParameters) {
				finalUrl = finalUrl + "&"
			}
		}
	}
	return finalUrl
}

func reqAddHeader(headerParameters []string, req *http.Request) *http.Request {
	if len(headerParameters) > 0 {
		for i := range headerParameters {
			headerToAdd := strings.Split(headerParameters[i], "=")
			if len(headerToAdd) > 1 {
				req.Header.Add(headerToAdd[0], headerToAdd[1])
			} else {
				log.Fatalln("wrong header flag")
			}
		}
	}
	return req
}

var rootCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "root",
	Short: "sent GET request to URL",
	Long:  `sent GET request to URL`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		finalUrl := args[0]
		finalUrl = urlAddQuery(finalUrl)
		req, err := http.NewRequest("GET", finalUrl, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req = reqAddHeader(headerParameters, req)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		strBody := string(body)
		log.Print(strBody)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
