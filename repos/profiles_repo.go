package repos

import (
	"database/sql"
	"fmt"
	"profile_service/app"
	"profile_service/domain"
	"strconv"
	"strings"
)

type ProfilesRepo struct {
	log app.Logger
	db  *sql.DB
}

func NewProfilesRepo(log app.Logger, db *sql.DB) *ProfilesRepo {
	return &ProfilesRepo{log: log, db: db}
}

func (r *ProfilesRepo) GetProfiles(ides []int64) (*[]domain.Profile, error) {
	var profiles []domain.Profile
	var profile domain.Profile
	var str []string
	for _, ide := range ides {
		str = append(str, strconv.Itoa(int(ide)))
	}
	str1 := strings.Join(str, ",")
	query := fmt.Sprintf(`SELECT name, photo, created_at, updated_at FROM profiles WHERE user_id IN (%s)`, str1)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&profile.Name, &profile.Photo, &profile.CreatedAt, &profile.UpdatedAt)
		profiles = append(profiles, profile)
	}
	return &profiles, nil
}

func (r *ProfilesRepo) List(offset, limit int, orderBy string) (*[]domain.Profile, error) {
	var profiles []domain.Profile
	var profile domain.Profile
	query := `SELECT user_id ,name, photo, created_at, updated_at FROM profiles WHERE user_id<=$1 AND user_id>=$2`
	rows, err := r.db.Query(query, offset, orderBy)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&profile.UserId, &profile.Name, &profile.Photo, &profile.CreatedAt, &profile.UpdatedAt)
		profiles = append(profiles, profile)
	}
	return &profiles, nil
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
