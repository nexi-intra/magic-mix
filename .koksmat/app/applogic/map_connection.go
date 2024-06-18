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
	"github.com/magicbutton/magic-mix/services/models/connectionmodel"
)

func MapConnectionOutgoing(db database.Connection) connectionmodel.Connection {
	return connectionmodel.Connection{
		ID:               db.ID,
		CreatedAt:        db.CreatedAt,
		CreatedBy:        db.CreatedBy,
		UpdatedAt:        db.UpdatedAt,
		UpdatedBy:        db.UpdatedBy,
		Name:             db.Name,
		Description:      db.Description,
		Connectionstring: db.Connectionstring,
	}
}

func MapConnectionIncoming(in connectionmodel.Connection) database.Connection {
	return database.Connection{
		ID:               in.ID,
		CreatedAt:        in.CreatedAt,
		CreatedBy:        in.CreatedBy,
		UpdatedAt:        in.UpdatedAt,
		UpdatedBy:        in.UpdatedBy,
		Name:             in.Name,
		Description:      in.Description,
		Connectionstring: in.Connectionstring,
		Searchindex:      in.Name,
	}
}
