package magicapp

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect/officegraph"
	"github.com/magicbutton/magic-mix/utils"
	"sigs.k8s.io/yaml"
)

func RegisterDownloadCmd() {
	downloadCmd := &cobra.Command{
		Use: "download",
	}
	downloadBatch := &cobra.Command{
		Use:   "batch [destination folder] [batchfile.json]",
		Short: "Start a downloader batch",
		Example: `
magic-mix download batch groups groupdata.json		

the format of batchfile.json is 
{
	"url": "https://graph.microsoft.com/v1.0/groups",
	"childUrls": [{
		"url": "https://graph.microsoft.com/v1.0/groups/%s/owners",
		"prefix": "owners"
	},
	{
		"url": "https://graph.microsoft.com/v1.0/groups/%s/members",
		"prefix": "members"
	},
	]
}

Alternatively, you can use a YAML file and convert it to JSON on the fly:

magic-mix download batch groups groupdata.yaml

the format of the YAML file is

url: https://graph.microsoft.com/v1.0/groups
childUrls:
	- url: https://graph.microsoft.com/v1.0/groups/%s/owners
	  prefix: owners
	- url: https://graph.microsoft.com/v1.0/groups/%s/members
	  prefix: members
			

		`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			batchID := args[0]
			batchFile := strings.ToLower(args[1])
			batchData, err := os.ReadFile(batchFile)

			if strings.HasSuffix(batchFile, ".yaml") {
				batchData, err = yaml.YAMLToJSON(batchData)
				if err != nil {
					log.Fatal("Error converting YAML", err)
				}
			}
			if err != nil {
				log.Fatal("Reading batch file", err)
			}

			batch := officegraph.DownloadBatchType{}

			json.Unmarshal(batchData, &batch)
			log.Println("Preparing download of batch", batchID)
			officegraph.DownloadBatch(batchID, batch, nil)

		},
	}

	downloadAuditLog := &cobra.Command{
		Use:   "auditlog [destination folder] [[year]] [[month]] [[day]]",
		Short: "Download audit logs for a given day",
		Long:  "Download audit logs for a given day. If year, month and day are not provided, yesterday is used.",
		Example: `

magic-mix download auditlog auditlogs 	

magic-mix download auditlog auditlogs 2024 12 1		

		`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			batchID := args[0]
			var year int
			var month time.Month
			var day int
			if len(args) == 1 {
				yesterday := time.Now().AddDate(0, 0, -1)
				year = yesterday.Year()
				month = yesterday.Month()
				day = yesterday.Day()

				log.Println("Preparing download of audit logs for yesterday")
			} else {
				if len(args) != 4 {
					log.Fatal("You need to specify year, month and day")
				}
				year = utils.StrToInt(args[1])
				month = time.Month(utils.StrToInt(args[2]))
				day = utils.StrToInt(args[3])
				log.Println("Preparing download of audit logs for a given day")

			}
			var date = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

			err := officegraph.GetAuditLogsForADay(batchID, date)
			if err != nil {
				log.Fatal("Error downloading audit logs", err)
			}

		},
	}
	downloadCmd.AddCommand(downloadBatch)
	downloadCmd.AddCommand(downloadAuditLog)
	utils.RootCmd.AddCommand(downloadCmd)
}
