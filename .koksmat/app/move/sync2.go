package move

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/lib/pq"
)

// ConnectDB initializes a connection to the database using a connection string
func ConnectDB(connStr string) (*sql.DB, error) {
    return sql.Open("postgres", connStr)
}

// ExtractBatchNames retrieves distinct batch names from the specified table
func ExtractBatchNames(ctx context.Context, db *sql.DB, query string) (map[string]struct{}, error) {
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("failed to query batch names: %w", err)
    }
    defer rows.Close()

    batchNames := make(map[string]struct{})
    for rows.Next() {
        var batchName string
        if err := rows.Scan(&batchName); err != nil {
            return nil, fmt.Errorf("failed to scan batch name: %w", err)
        }
        batchNames[batchName] = struct{}{}
    }
    return batchNames, nil
}

// ExtractDataAsJSON retrieves data for a specific batch name and returns it as JSON
func ExtractDataAsJSON(ctx context.Context, db *sql.DB, batchName string) (string, error) {
    query := "SELECT * FROM sharepoint.pageviews WHERE batchname = $1"
    rows, err := db.QueryContext(ctx, query, batchName)
    if err != nil {
        return "", fmt.Errorf("failed to query pageviews for batchname %s: %w", batchName, err)
    }
    defer rows.Close()

    var results []map[string]interface{}
    cols, err := rows.Columns()
    if err != nil {
        return "", fmt.Errorf("failed to get columns: %w", err)
    }

    for rows.Next() {
        row := make(map[string]interface{})
        vals := make([]interface{}, len(cols))
        for i := range vals {
            vals[i] = new(interface{})
        }
        if err := rows.Scan(vals...); err != nil {
            return "", fmt.Errorf("failed to scan row: %w", err)
        }
        for i, col := range cols {
            row[col] = *(vals[i].(*interface{}))
        }
        results = append(results, row)
    }

    jsonData, err := json.Marshal(results)
    if err != nil {
        return "", fmt.Errorf("failed to marshal JSON: %w", err)
    }
    return string(jsonData), nil
}

// InsertIntoImportTable inserts JSON data into the import table
func InsertIntoImportTable(ctx context.Context, db *sql.DB, batchName, jsonData string) error {
    query := `
        INSERT INTO public.importdata (
            created_by, updated_by, tenant, searchindex, name, data
        ) VALUES (
            '', '', '', '', '', $1
        )`
    _, err := db.ExecContext(ctx, query, jsonData)
    if err != nil {
        return fmt.Errorf("failed to insert into import table: %w", err)
    }
    return nil
}

// SyncPageViews syncs data from sharepoint.pageviews to files.events
func SyncPageViews(ctx context.Context, mixerConnStr, filesConnStr string) error {
    // Connect to mixer database
    mixerDB, err := ConnectDB(mixerConnStr)
    if err != nil {
        return fmt.Errorf("failed to connect to mixer database: %w", err)
    }
    defer mixerDB.Close()

    // Connect to files database
    filesDB, err := ConnectDB(filesConnStr)
    if err != nil {
        return fmt.Errorf("failed to connect to files database: %w", err)
    }
    defer filesDB.Close()

    // Extract batch names
    mixerBatchNames, err := ExtractBatchNames(ctx, mixerDB, "SELECT DISTINCT batchname FROM sharepoint.pageviews")
    if err != nil {
        return err
    }

    filesBatchNames, err := ExtractBatchNames(ctx, filesDB, "SELECT DISTINCT batchname FROM events")
    if err != nil {
        return err
    }

    // Identify batch names only in mixer
    for batchName := range mixerBatchNames {
        if _, exists := filesBatchNames[batchName]; !exists {
            // Extract data as JSON
            jsonData, err := ExtractDataAsJSON(ctx, mixerDB, batchName)
            if err != nil {
                return err
            }

            // Insert into import table
            if err := InsertIntoImportTable(ctx, filesDB, batchName, jsonData); err != nil {
                return err
            }
        }
    }

    // Run stored procedure to process data
    _, err = filesDB.ExecContext(ctx, "CALL process_import_data()")
    if err != nil {
        return fmt.Errorf("failed to call stored procedure: %w", err)
    }

    return nil
}
