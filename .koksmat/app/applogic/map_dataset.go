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
	"github.com/magicbutton/magic-mix/services/models/datasetmodel"
)

func MapDatasetOutgoing(db database.Dataset) datasetmodel.Dataset {
	return datasetmodel.Dataset{
		ID:             db.ID,
		CreatedAt:      db.CreatedAt,
		CreatedBy:      db.CreatedBy,
		UpdatedAt:      db.UpdatedAt,
		UpdatedBy:      db.UpdatedBy,
		Name:           db.Name,
		Description:    db.Description,
		Connection_id:  db.Connection_id,
		Transformer_id: db.Transformer_id,
	}
}

func MapDatasetIncoming(in datasetmodel.Dataset) database.Dataset {
	return database.Dataset{
		ID:             in.ID,
		CreatedAt:      in.CreatedAt,
		CreatedBy:      in.CreatedBy,
		UpdatedAt:      in.UpdatedAt,
		UpdatedBy:      in.UpdatedBy,
		Name:           in.Name,
		Description:    in.Description,
		Connection_id:  in.Connection_id,
		Transformer_id: in.Transformer_id,
		Searchindex:    in.Name,
	}
}
