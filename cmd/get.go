package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "get",
	Short: "get a data from Url",
	Long:  "get a data from Url",
	Run: func(cmd *cobra.Command, args []string) {
		ID, _ := cmd.Flags().GetString("ID")
		finalUrl := args[0] + "/todos/" + ID
		client := &http.Client{}

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
		fmt.Println("get called get")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().String("ID", "", "ID of the task you want.")
	rootCmd.PersistentFlags().StringArrayVarP(&queryArr, "query", "", []string{}, "query to be ask.")
}
