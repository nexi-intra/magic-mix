package query

import "testing"

func TestMove(t *testing.T) {
	records, err := QueryGetJSON("booking", `SELECT proname
FROM pg_proc
JOIN pg_namespace ON pg_proc.pronamespace = pg_namespace.oid
WHERE pg_namespace.nspname NOT IN ('pg_catalog', 'information_schema')
AND pg_proc.prokind = 'p';
`)
	if err != nil {
		t.Errorf("Move() failed: %v", err)
	}
	if len(*records) == 0 {
		t.Errorf("Move() failed: %v", err)
	}
}
