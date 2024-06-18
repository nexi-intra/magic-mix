/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic

import (
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/jobmodel"
)

func MapJobOutgoing(db database.Job) jobmodel.Job {
	return jobmodel.Job{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		CreatedBy:   db.CreatedBy,
		UpdatedAt:   db.UpdatedAt,
		UpdatedBy:   db.UpdatedBy,
		Name:        db.Name,
		Description: db.Description,
		Status:      db.Status,
		Script:      db.Script,
		Data:        db.Data,
	}
}

func MapJobIncoming(in jobmodel.Job) database.Job {
	return database.Job{
		ID:          in.ID,
		CreatedAt:   in.CreatedAt,
		CreatedBy:   in.CreatedBy,
		UpdatedAt:   in.UpdatedAt,
		UpdatedBy:   in.UpdatedBy,
		Name:        in.Name,
		Description: in.Description,
		Status:      in.Status,
		Script:      in.Script,
		Data:        in.Data,
		Searchindex: in.Name,
	}
}
