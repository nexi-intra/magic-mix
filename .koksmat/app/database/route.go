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

type Route struct {
	bun.BaseModel `bun:"table:route,alias:route"`

	ID          int       `bun:"id,pk,autoincrement"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy   string    `bun:"created_by,"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy   string    `bun:"updated_by,"`
	DeletedAt   time.Time `bun:",soft_delete,nullzero"`
	Tenant      string    `bun:"tenant"`
	Searchindex string    `bun:"searchindex"`
	Name        string    `bun:"name"`
	Description string    `bun:"description"`
	Method      string    `bun:"method"`
	Slug        string    `bun:"slug"`
	Api_id      int       `bun:"api_id"`
}
