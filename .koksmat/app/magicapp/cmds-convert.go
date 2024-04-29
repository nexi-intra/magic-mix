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
		Use:   "sql [excelfilename] [sheetname] [namespace]",
		Short: "Transform the data to SQL",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			excelfilename := args[0]
			sheetname := args[1]
			namespace := args[2]
			tablename := args[3]
			batchsize := 1000
			applogic.ConvertExcelToSQL(excelfilename, sheetname, namespace, tablename, batchsize)
		},
	}
	toCmd.AddCommand(toSQLcmd)

	utils.RootCmd.AddCommand(convertCmd)
}
