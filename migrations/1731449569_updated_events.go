package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("uqrkgbsnoo7lbj6")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("eotswooq")

		// add
		new_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gcqogmpx",
			"name": "description",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_description); err != nil {
			return err
		}
		collection.Schema.AddField(new_description)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("uqrkgbsnoo7lbj6")
		if err != nil {
			return err
		}

		// add
		del_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "eotswooq",
			"name": "description",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_description); err != nil {
			return err
		}
		collection.Schema.AddField(del_description)

		// remove
		collection.Schema.RemoveField("gcqogmpx")

		return dao.SaveCollection(collection)
	})
}
