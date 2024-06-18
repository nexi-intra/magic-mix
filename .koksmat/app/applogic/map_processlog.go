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
	"github.com/magicbutton/magic-mix/services/models/processlogmodel"
)

func MapProcessLogOutgoing(db database.ProcessLog) processlogmodel.ProcessLog {
	return processlogmodel.ProcessLog{
		ID:                db.ID,
		CreatedAt:         db.CreatedAt,
		CreatedBy:         db.CreatedBy,
		UpdatedAt:         db.UpdatedAt,
		UpdatedBy:         db.UpdatedBy,
		Name:              db.Name,
		Description:       db.Description,
		Transformation_id: db.Transformation_id,
		Status:            db.Status,
		Message:           db.Message,
	}
}

func MapProcessLogIncoming(in processlogmodel.ProcessLog) database.ProcessLog {
	return database.ProcessLog{
		ID:                in.ID,
		CreatedAt:         in.CreatedAt,
		CreatedBy:         in.CreatedBy,
		UpdatedAt:         in.UpdatedAt,
		UpdatedBy:         in.UpdatedBy,
		Name:              in.Name,
		Description:       in.Description,
		Transformation_id: in.Transformation_id,
		Status:            in.Status,
		Message:           in.Message,
		Searchindex:       in.Name,
	}
}
