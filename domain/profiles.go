package domain

import "time"

type Profile struct {
	UserId    int64     `json:"id"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ProfilesRepository interface {
	List(offset, limit int, orderBy string) (*[]Profile, error)
	FetchByUserId(int64) (*Profile, error)
	Update(*Profile) error
	GetProfiles(ides []int64) (*[]Profile, error)
}

type ProfilesInteractor interface {
	List(offset, limit int, orderBy string) (ListProfileResponse, error)
	Get(userId int64) (GetProfileResponse, error)
	Update(Profile) (UpdateProfileResponse, error)
	GetProfiles(ides []int64) (ListProfileResponse, error)
}

// Responses (only for UseCase layer)

type ListProfileResponse struct {
	StatusCode string
	Profiles   []Profile
}

type GetProfileResponse struct {
	StatusCode string
	Profile    *Profile
}

type UpdateProfileResponse struct {
	StatusCode string
}
