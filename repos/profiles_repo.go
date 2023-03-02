package repos

import (
	"database/sql"
	"profile_service/app"
	"profile_service/domain"
)

type ProfilesRepo struct {
	log app.Logger
	db  *sql.DB
}

func NewProfilesRepo(log app.Logger, db *sql.DB) *ProfilesRepo {
	return &ProfilesRepo{log: log, db: db}
}

func (r *ProfilesRepo) FetchByUserId(i int64) (*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProfilesRepo) Update(profile *domain.Profile) error {
	//TODO implement me
	panic("implement me")
}
