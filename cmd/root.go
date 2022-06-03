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

var rootCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "root",
	Short: "sent GET request to URL",
	Long:  `sent GET request to URL`,
	Run: func(cmd *cobra.Command, args []string) {
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

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Print(sb)
		fmt.Println("get called root")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringArrayVarP(&headerArr, "header", "", []string{}, "add this key and value into header")
}
