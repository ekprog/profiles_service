package bootstrap

import (
	"database/sql"
	"profile_service/app"
	"profile_service/delivery"
	"profile_service/interactors"
	"profile_service/repos"
)

func injectDependencies(db *sql.DB, logger app.Logger) error {

	// DI Manual
	profilesRepo := repos.NewProfilesRepo(logger, db)
	profileUcase := interactors.NewProfilesInteractor(logger, profilesRepo)

	// Delivery Init
	tasksDelivery := delivery.NewTasksDeliveryService(logger, profileUcase)
	assetsDelivery := delivery.NewAssetsDelivery(logger)

	err := app.InitDelivery(tasksDelivery)
	if err != nil {
		return err
	}

	err = app.InitRestDelivery("/assets", assetsDelivery)
	if err != nil {
		return err
	}

	return nil
}
