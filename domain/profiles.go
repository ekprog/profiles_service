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
	FetchByUserId(int64) (*Profile, error)
	Update(*Profile) error
}

type ProfilesInteractor interface {
	Get(userId int64) (GetProfileResponse, error)
	Update(Profile) (UpdateProfileResponse, error)
}

// Responses (only for UseCase layer)

type GetProfileResponse struct {
	StatusCode string
	Profile    *Profile
}

type UpdateProfileResponse struct {
	StatusCode string
}
