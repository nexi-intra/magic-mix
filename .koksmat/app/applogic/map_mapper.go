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
	"github.com/magicbutton/magic-mix/services/models/mappermodel"
)

func MapMapperOutgoing(db database.Mapper) mappermodel.Mapper {
	return mappermodel.Mapper{
		ID:                db.ID,
		CreatedAt:         db.CreatedAt,
		CreatedBy:         db.CreatedBy,
		UpdatedAt:         db.UpdatedAt,
		UpdatedBy:         db.UpdatedBy,
		Name:              db.Name,
		Description:       db.Description,
		Source_id:         db.Source_id,
		Transformation_id: db.Transformation_id,
		Target_id:         db.Target_id,
	}
}

func MapMapperIncoming(in mappermodel.Mapper) database.Mapper {
	return database.Mapper{
		ID:                in.ID,
		CreatedAt:         in.CreatedAt,
		CreatedBy:         in.CreatedBy,
		UpdatedAt:         in.UpdatedAt,
		UpdatedBy:         in.UpdatedBy,
		Name:              in.Name,
		Description:       in.Description,
		Source_id:         in.Source_id,
		Transformation_id: in.Transformation_id,
		Target_id:         in.Target_id,
		Searchindex:       in.Name,
	}
}
