package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/magicbutton/magic-mix/officegraph"
	"github.com/magicbutton/magic-mix/officegraph/sites"
	"github.com/magicbutton/magic-mix/sharepoint/sites/nexiintra_home"
)

type Newschannel struct {
	mgm.DefaultModel `bson:",inline"`
	Item             *nexiintra_home.NewsChannels `bson:"inline"`
}

func CreateNewsChannel(channel sites.NewsChannelsListItem) (newsChannel *Newschannel, err error) {

	newRecord := &Newschannel{}
	newRecord.Item = &nexiintra_home.NewsChannels{
		Title:         channel.NewsChannel.Title,
		RelevantUnits: []nexiintra_home.Units{},
		Mandatory:     channel.NewsChannel.Mandatory,
		//NewsCategory: channel.NewsChannel.Newscategory,

	}

	err = mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}

func ImportNewsChannels() error {

	_, token, err := officegraph.GetClient()
	if err != nil {
		return err
	}

	got, err := sites.GetListItems[sites.NewsChannelsListItem](token, "sites/nexiintra-home", "News Channels", "")
	if err != nil {
		return err
	}

	for _, channel := range *got {

		_, err := CreateNewsChannel(channel)
		if err != nil {
			return err
		}

	}

	return nil
}
