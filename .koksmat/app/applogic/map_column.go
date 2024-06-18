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
	"github.com/magicbutton/magic-mix/services/models/columnmodel"
)

func MapColumnOutgoing(db database.Column) columnmodel.Column {
	return columnmodel.Column{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		CreatedBy:   db.CreatedBy,
		UpdatedAt:   db.UpdatedAt,
		UpdatedBy:   db.UpdatedBy,
		Name:        db.Name,
		Description: db.Description,
		Datatype:    db.Datatype,
		Sortorder:   db.Sortorder,
	}
}

func MapColumnIncoming(in columnmodel.Column) database.Column {
	return database.Column{
		ID:          in.ID,
		CreatedAt:   in.CreatedAt,
		CreatedBy:   in.CreatedBy,
		UpdatedAt:   in.UpdatedAt,
		UpdatedBy:   in.UpdatedBy,
		Name:        in.Name,
		Description: in.Description,
		Datatype:    in.Datatype,
		Sortorder:   in.Sortorder,
		Searchindex: in.Name,
	}
}
