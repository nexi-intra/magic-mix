/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn2
package database

import (
	"time"
    
	"github.com/uptrace/bun"
)

type Job struct {
	bun.BaseModel `bun:"table:job,alias:job"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Searchindex string `bun:"searchindex"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Status string `bun:"status"`
    Startat time.Time `bun:"startat"`
    Startedat time.Time `bun:"startedAt"`
    Completedat time.Time `bun:"completedAt"`
    Maxduration int `bun:"maxduration"`
    Script string `bun:"script"`
    Data interface{} `bun:"data"`

}

