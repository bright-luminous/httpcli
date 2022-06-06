package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().String("ID", "", "ID of the task you want.")
}

var getCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "get",
	Short: "get a data from Url",
	Long:  "get a data from Url",
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := cmd.Flags().GetString("ID")
		if err != nil {
			log.Fatalln(err)
		}
		headerParameters, err := cmd.Flags().GetStringSlice(flagHeader)
		if err != nil {
			log.Fatalln(err)
		}
		queryParameters, err := cmd.Flags().GetStringSlice(flagQuery)
		if err != nil {
			log.Fatalln(err)
		}
		finalUrl := args[0] + "/todos/" + ID
		client := &http.Client{
			Timeout: 20 * time.Second,
		}

		finalUrl = urlAddQuery(finalUrl, queryParameters)
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
