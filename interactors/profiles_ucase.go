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
	profile, err := i.repo.FetchByUserId(userId)
	if err != nil {
		return domain.GetProfileResponse{}, err
	}
	return domain.GetProfileResponse{StatusCode: domain.Success, Profile: profile}, nil
}

func (i *ProfilesInteractor) Update(profile domain.Profile) (domain.UpdateProfileResponse, error) {
	err := i.repo.Update(&profile)
	if err != nil {
		return domain.UpdateProfileResponse{}, err
	}
	return domain.UpdateProfileResponse{StatusCode: domain.Success}, err
}
