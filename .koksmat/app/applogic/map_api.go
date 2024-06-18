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
	"github.com/magicbutton/magic-mix/services/models/apimodel"
)

func MapAPIOutgoing(db database.API) apimodel.API {
	return apimodel.API{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		CreatedBy:   db.CreatedBy,
		UpdatedAt:   db.UpdatedAt,
		UpdatedBy:   db.UpdatedBy,
		Name:        db.Name,
		Description: db.Description,
		Method:      db.Method,
		Source:      db.Source,
		Schema:      db.Schema,
	}
}

func MapAPIIncoming(in apimodel.API) database.API {
	return database.API{
		ID:          in.ID,
		CreatedAt:   in.CreatedAt,
		CreatedBy:   in.CreatedBy,
		UpdatedAt:   in.UpdatedAt,
		UpdatedBy:   in.UpdatedBy,
		Name:        in.Name,
		Description: in.Description,
		Method:      in.Method,
		Source:      in.Source,
		Schema:      in.Schema,
		Searchindex: in.Name,
	}
}
