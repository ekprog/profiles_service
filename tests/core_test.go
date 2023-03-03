package tests

import (
	"fmt"
	"profile_service/app"
	"profile_service/interactors"
	"profile_service/repos"
	"testing"
)

func TestCore(t *testing.T) {

	err := app.InitApp("..")
	if err != nil {
		panic(err)
	}

	logger, err := app.InitLogs("..")
	if err != nil {
		panic(err)
	}
	db, err := app.InitDatabase()
	if err != nil {
		panic(err)
	}

	// Migrations
	err = app.RunMigrations("..")
	if err != nil {
		panic(err)
	}

	repo := repos.NewProfilesRepo(logger, db)
	res, err := interactors.NewProfilesInteractor(logger, repo).Get(1)
	fmt.Printf("%v", res)
}
