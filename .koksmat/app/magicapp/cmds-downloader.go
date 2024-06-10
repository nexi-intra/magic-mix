package magicapp

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect/officegraph"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterDownloadCmd() {
	natsCmd := &cobra.Command{
		Use:   "download [destination folder] [batchfile.json]",
		Short: "Start a downloader batch",
		Example: `
magic-mix download groups groupdata.json		

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
	utils.RootCmd.AddCommand(natsCmd)
}
