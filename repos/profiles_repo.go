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

func (r *ProfilesRepo) List(offset, limit int, orderBy string) ([]domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProfilesRepo) FetchByUserId(id int64) (*domain.Profile, error) {
	var profile = &domain.Profile{
		UserId: id,
	}
	query := `SELECT name, photo, created_at, updated_at FROM profiles WHERE user_id=$1`
	err := r.db.QueryRow(query, id).Scan(&profile.Name, &profile.Photo, &profile.CreatedAt, &profile.UpdatedAt)
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
	query := `UPDATE profiles SET name=$2, photo=$3, updated_at=NOW() WHERE user_id=$1`
	_, err := r.db.Exec(query, profile.UserId, profile.Name, profile.Photo)
	if err != nil {
		return err
	}
	return nil
}
