package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/magicbutton/magic-mix/officegraph"
	"github.com/magicbutton/magic-mix/officegraph/sites"
	"github.com/magicbutton/magic-mix/shared"
	"github.com/magicbutton/magic-mix/sharepoint/sites/nexiintra_home"
)

type UnitListItem struct {
	shared.Item `bson:",inline"`
	UnitData    *nexiintra_home.SP_Units `json:"fields,inline"`
}

type Unit struct {
	mgm.DefaultModel `bson:",inline"`
	Item             UnitListItem `bson:",inline"`
}

func CreateUnit(sharepoint_unit UnitListItem) (unit *Unit, err error) {

	newRecord := &Unit{}
	newRecord.Item = sharepoint_unit

	err = mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}

func ImportUnits() error {

	_, token, err := officegraph.GetClient()
	if err != nil {
		return err
	}

	got, err := sites.GetListItems[UnitListItem](token, "sites/nexiintra-home", "Units", "")
	if err != nil {
		return err
	}

	for _, channel := range *got {

		_, err := CreateUnit(channel)
		if err != nil {
			return err
		}

	}

	return nil
}
