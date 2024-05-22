package magicapp

import (
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

	utils.RootCmd.AddCommand(convertCmd)
}
