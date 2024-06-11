package officegraph

import (
	"testing"
	"time"
)

func TestGetAuditLogs(t *testing.T) {

	var day1 = time.Date(2024, 5, 17, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 30; i++ {

		logDay := day1.AddDate(0, 0, i)
		for h := 0; h < 24; h++ {
			from := logDay.Add(time.Duration(h) * time.Hour)
			to := logDay.Add(time.Duration(h+1)*time.Hour - 1*time.Second)
			GetAuditLogs("temp-auditlog", from, to)
		}

	}
}
