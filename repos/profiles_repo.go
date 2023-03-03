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
	var profile = &domain.Profile{}
	query := `SELECT id,name,photo FROM profiles WHERE id=$1`
	err := r.db.QueryRow(query, i).Scan(&profile.UserId, &profile.Name, &profile.Photo, &profile.CreatedAt, &profile.UpdatedAt) //TODO implement me
	switch err {
	case nil:
		return profile, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *ProfilesRepo) Update(profile *domain.Profile) error {
	query := `INSERT INTO profiles (name,photo,created_at,updated_at) VALUES ($1,$2)`
	_, err := r.db.Exec(query, profile.Name, profile.Photo, profile.CreatedAt, profile.UpdatedAt)
	return err
}
