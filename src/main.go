package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/IqbalLx/rsvp-pb/migrations"
	rsvps "github.com/IqbalLx/rsvp-pb/src/features/rsvp"
)

func main() {
    app := pocketbase.New()

    // loosely check if it was executed using "go run"
    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        // enable auto creation of migration files when making collection changes in the Admin UI
        // (the isGoRun check is to enable it only during development)
        Automigrate: isGoRun,
    })

	rsvps.NewRSVPController(app)

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}