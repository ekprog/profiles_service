package interactors

import (
	"github.com/pkg/errors"
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

func (i *ProfilesInteractor) List(offset, limit int, orderBy string) (domain.ListProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ProfilesInteractor) Get(userId int64) (domain.GetProfileResponse, error) {
	profile, err := i.repo.FetchByUserId(userId)
	if err != nil {
		return domain.GetProfileResponse{}, errors.Wrap(err, "error while fetching profile by user_id")
	}

	profile.Photo = "http://127.0.0.1:8089/assets/" + profile.Photo
	return domain.GetProfileResponse{StatusCode: domain.Success, Profile: profile}, nil
}

func (i *ProfilesInteractor) Update(profile domain.Profile) (domain.UpdateProfileResponse, error) {
	err := i.repo.Update(&profile)
	if err != nil {
		return domain.UpdateProfileResponse{}, errors.Wrap(err, "error while updating profile")
	}
	return domain.UpdateProfileResponse{StatusCode: domain.Success}, err
}
