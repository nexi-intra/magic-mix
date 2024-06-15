package magicapp

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect/officegraph"
	"github.com/magicbutton/magic-mix/utils"
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


		`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			batchID := args[0]
			batchFile := args[1]
			batchData, err := os.ReadFile(batchFile)
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
		Use:   "auditlog [destination folder] [year] [month] [day]",
		Short: "Download audit logs for a given day, a file will be created for each hour",
		Example: `
magic-mix download auditlog auditlogs 2024 12 1		



		`,
		Args: cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			batchID := args[0]
			year := utils.StrToInt(args[1])
			month := time.Month(utils.StrToInt(args[2]))
			day := utils.StrToInt(args[3])
			var date = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			log.Println("Preparing download of audit logs", batchID)
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
