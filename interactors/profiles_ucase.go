package interactors

import (
	"profile_service/app"
	"profile_service/domain"
)

type ProfilesInteractor struct {
	log  app.Logger
	repo domain.ProfilesRepository
}

func NewProfilesInteractor(log app.Logger, repo domain.ProfilesRepository) *ProfilesInteractor {
	return &ProfilesInteractor{log: log, repo: repo}
}

func (i *ProfilesInteractor) Get(userId int64) (domain.GetProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ProfilesInteractor) Update(profile domain.Profile) (domain.UpdateProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}
