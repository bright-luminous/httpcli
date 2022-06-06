package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().String("ID", "", "ID of the task you want.")
	putCmd.PersistentFlags().String("json", "", "task to be update.")
}

var putCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "put",
	Short: "put request to url",
	Long:  `put request to url`,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := cmd.Flags().GetString("ID")
		if err != nil {
			log.Fatalln(err)
		}
		finalUrl := args[0] + "/todos/" + ID
		jsonData, err := cmd.Flags().GetString("json")
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
		client := &http.Client{
			Timeout: 20 * time.Second,
		}

		fmt.Println(jsonData)
		jsonDataByte := []byte(jsonData)

		finalUrl = urlAddQuery(finalUrl, queryParameters)
		req, err := http.NewRequest("PUT", finalUrl, bytes.NewBuffer(jsonDataByte))
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
