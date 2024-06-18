package magicapp

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/move"
	"github.com/magicbutton/magic-mix/utils"
)

type SQLResult struct {
	Response     interface{} `json:"response"`
	RowsAffected int64       `json:"rows_affected"`
}

func SelectSQL(connectionName string, sqlstatement string) (*string, error) {

	OpenDatabase()
	defer utils.Db.Close()
	toConnection, err := move.GetConnectionString(connectionName)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", *toConnection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	jsonsql := fmt.Sprintf(`
	SELECT json_agg(json_data) AS result
	FROM (
		%s
	) AS json_data;
		
	`, sqlstatement)

	rows, err := db.Query(jsonsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	returnValue := ""
	// Iterate over the result set and append each distinct value to the slice
	for rows.Next() {
		var value string

		err := rows.Scan(&value)
		if err != nil {
			return nil, err
		}
		returnValue = value
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &returnValue, nil
}

// func SelectSQLFromFile(filepath string) (*string, error) {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {

// 		return nil, fmt.Errorf("could not read file: %w", err)
// 	}

// 	text := string(data)
// 	return SelectSQL(text)

// }

func ExecuteSQL(connectionName string, sqlstatement string) (int64, error) {

	OpenDatabase()
	defer utils.Db.Close()
	toConnection, err := move.GetConnectionString(connectionName)
	if err != nil {
		return 0, err
	}
	db, err := sql.Open("postgres", *toConnection)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	ctx := context.Background()

	raw, err := db.ExecContext(ctx, sqlstatement)
	if err != nil {
		return -1, err
	}

	rowsAffected, _ := raw.RowsAffected()
	return rowsAffected, nil
}

// func ExecuteSQLFromFile(filepath string) (int64, error) {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {

// 		return -1, fmt.Errorf("could not read file: %w", err)
// 	}

// 	text := string(data)
// 	return ExecuteSQL(text)

// }

func RegisterSQLCmd() {
	sqlCmd := &cobra.Command{
		Use:   "sql",
		Short: "Execute SQL",
	}
	selectCmd := &cobra.Command{
		Use:     "query database sqlstatement",
		Short:   "Select SQL",
		Example: `magic-devices sql query files "SELECT 1 as col1"`,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			result, err := SelectSQL(args[0], args[1])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(*result)

		},
	}
	executeCmd := &cobra.Command{
		Use:     "exec database sqlstatement",
		Short:   "Execute SQL",
		Example: `magic-devices sql exec files "DELETE from table where id = 1"`,
		Args:    cobra.MinimumNArgs(2),

		Run: func(cmd *cobra.Command, args []string) {

			result, err := ExecuteSQL(args[0], args[1])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(result)

		},
	}
	sqlCmd.AddCommand(selectCmd)
	sqlCmd.AddCommand(executeCmd)
	utils.RootCmd.AddCommand(sqlCmd)
}
