package rsvps

import "github.com/pocketbase/pocketbase"

func NewRSVPController(app *pocketbase.PocketBase){
	app.OnRecordBeforeCreateRequest("rsvps").Add(fillUserID(app))
	app.OnRecordBeforeCreateRequest("rsvps").Add(validateMaxEventCapacity(app))
}