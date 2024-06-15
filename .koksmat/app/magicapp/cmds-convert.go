package magicapp

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterConvertCmd() {
	convertCmd := &cobra.Command{
		Use:   "from",
		Short: "Source",
	}

	excelCmd := &cobra.Command{
		Use:   "excel",
		Short: "Transform the data from one format to another",
	}

	convertCmd.AddCommand(excelCmd)

	toCmd := &cobra.Command{
		Use:   "to",
		Short: "Target",
	}
	excelCmd.AddCommand(toCmd)

	toSQLcmd := &cobra.Command{
		Use:   "sql excelfilename sheetname namespace [tablename]",
		Short: "Transform the data to SQL",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			excelfilename := args[0]
			sheetname := args[1]
			tablename := args[2]
			batchsize := 1000
			applogic.ConvertExcelToSQL(excelfilename, sheetname, tablename, batchsize)
		},
	}
	toCmd.AddCommand(toSQLcmd)

	toJSONcmd := &cobra.Command{
		Use:   "json excelfilename jsonfilename",
		Short: "Transform the file to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			excelfilename := args[0]
			jsonfilename := args[1]

			err := applogic.ConvertExcelToSQLJson(excelfilename, jsonfilename)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	toCmd.AddCommand(toJSONcmd)
	utils.RootCmd.AddCommand(convertCmd)
}
