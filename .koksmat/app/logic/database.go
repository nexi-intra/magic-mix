package logic

import (
	"github.com/magicbutton/magic-mix/database/dbhelpers"
	"github.com/magicbutton/magic-mix/utils"
)

/**
 * Search is a generic function that searches a table for a string in a column
 *
 * @param fieldname string
 * @param query string
 * @param mapper func(DB) DOC
 * @return *utils.Page[DOC]
 * @return error

 Example of mapper function:
 ```go
 func mapGroupsegment(item database.Groupsegment) groupsegmentmodel.Groupsegment {
	return groupsegmentmodel.Groupsegment{
		ID:   fmt.Sprintf("%d", item.ID),
		Name: item.Name,
	}
```
}
*/

type Identifiable struct {
	ID int
}

func Search[DB interface{}, DOC interface{}](fieldname string, query string, mapper func(DB) DOC) (*utils.Page[DOC], error) {

	result, err := dbhelpers.SelectWhereILike[DB](fieldname, query)

	if err != nil {
		return nil, err
	}
	items := []DOC{}
	for _, item := range result {
		mappedItem := mapper(item)
		items = append(items, mappedItem)
	}

	page := utils.Page[DOC]{
		Items:       items,
		TotalPages:  1,
		TotalItems:  len(items),
		CurrentPage: 0,
	}

	return &page, nil
}
func Create[DB interface{}, DOC interface{}](item DOC, mapperIncoming func(DOC) DB, mapperOutgoing func(DB) DOC) (*DOC, error) {

	dbItem := mapperIncoming(item)
	_, err := dbhelpers.Create[DB](dbItem)
	if err != nil {
		return nil, err
	}
	createdItem := mapperOutgoing(dbItem)
	return &createdItem, nil

}
func Read[DB interface{}, DOC interface{}](id int, mapper func(DB) DOC) (*DOC, error) {
	dbItem, err := dbhelpers.SelectById[DB](id)
	if err != nil {
		return nil, err
	}
	item := mapper(*dbItem)
	return &item, nil

}

func Update[DB interface{}, DOC interface{}](id int, item DOC, mapperIncoming func(DOC) DB, mapperOutgoing func(DB) DOC) (*DOC, error) {

	dbItem := mapperIncoming(item)
	err := dbhelpers.Update[DB](dbItem)
	if err != nil {
		return nil, err
	}
	updatedItem, err := dbhelpers.SelectById[DB](id)
	mappedItem := mapperOutgoing(*updatedItem)
	return &mappedItem, nil

}

func Delete[DB interface{}, DOC interface{}](id int) error {
	err := dbhelpers.DeleteById[DB](id)
	if err != nil {
		return err
	}

	return nil

}
