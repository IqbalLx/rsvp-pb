package rsvps

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func fillUserID(app *pocketbase.PocketBase) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		user := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)

		e.Record.Set("user_id", user.Id)

		return nil
	}
}

func validateMaxEventCapacity(app *pocketbase.PocketBase) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		eventID := e.Record.GetString("event_id")

		event, err := app.Dao().FindRecordById("events", eventID)
		if (err != nil) {
			return err
		}

		var eventAttendees struct { Count int }
		err = app.Dao().DB().
			Select("COUNT(*) as count").
			From("rsvps").
			Where(dbx.NewExp("event_id={:event_id}", dbx.Params{"event_id": eventID})).
			Build().One(&eventAttendees)
		if (err != nil) {
			return err
		}
		
		if (eventAttendees.Count + 1 > event.GetInt("max_capacity")) {
			return apis.NewBadRequestError("events full!", nil)
		}

		return nil
	}
}