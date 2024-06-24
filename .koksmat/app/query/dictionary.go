package query

import (
	"encoding/json"
	"fmt"
)

func GetStoredProcedures(database string) (*json.RawMessage, error) {
	return QueryGetJSON(database, `SELECT 
	pg_proc.oid as id,
    pg_namespace.nspname AS schema_name, 
    pg_proc.proname AS procedure_name
FROM 
    pg_proc
JOIN 
    pg_namespace ON pg_proc.pronamespace = pg_namespace.oid
WHERE 
    pg_namespace.nspname NOT IN ('pg_catalog', 'information_schema')
AND 
    pg_proc.prokind = 'p'`)
}

func GetStoredProcedure(database string, id int) (*json.RawMessage, error) {
	return QueryGetJSON(database, fmt.Sprintf(`SELECT 

    pg_namespace.nspname AS schema_name, 
    pg_proc.*
FROM 
    pg_proc
JOIN 
    pg_namespace ON pg_proc.pronamespace = pg_namespace.oid
WHERE 
    pg_proc = %d`, id))
}
