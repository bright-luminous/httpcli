package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "delete",
	Short: "delete request URL",
	Long:  `delete request URL`,
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
		req, err := http.NewRequest("DELETE", finalUrl, nil)
		if err != nil {
			log.Fatalln(err)
		}
		if len(headerArr) > 0 {
			for i := range headerArr {
				headerToAdd := strings.Split(headerArr[i], "=")
				req.Header.Add(headerToAdd[0], headerToAdd[1])
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
		fmt.Println("get called delete")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().String("ID", "", "ID of the task you want to delete.")
}
