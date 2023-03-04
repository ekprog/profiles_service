package delivery

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"profile_service/app"
	"profile_service/domain"
	pb "profile_service/pkg/pb/api"
)

type ProfilesDeliveryService struct {
	pb.ProfilesServiceServer
	log           app.Logger
	profilesUCase domain.ProfilesInteractor
}

func NewTasksDeliveryService(log app.Logger, profilesInteractor domain.ProfilesInteractor) *ProfilesDeliveryService {
	return &ProfilesDeliveryService{
		log:           log,
		profilesUCase: profilesInteractor,
	}
}

func (d *ProfilesDeliveryService) Init() error {
	app.InitGRPCService(pb.RegisterProfilesServiceServer, pb.ProfilesServiceServer(d))
	return nil
}

func (d *ProfilesDeliveryService) List(ctx context.Context, r *pb.ListProfileRequest) (*pb.ListProfileResponse, error) {

	uCaseRes, err := d.profilesUCase.List(int(r.Offset), int(r.Limit), r.OrderBy)
	if err != nil {
		return nil, err
	}

	response := &pb.ListProfileResponse{
		Status: &pb.StatusResponse{
			Code:    uCaseRes.StatusCode,
			Message: uCaseRes.StatusCode,
		},
		Profiles: nil,
	}

	if uCaseRes.StatusCode == domain.Success && uCaseRes.Profiles != nil {
		for _, profile := range uCaseRes.Profiles {
			response.Profiles = append(response.Profiles, &pb.Profile{
				UserId:    profile.UserId,
				Name:      profile.Name,
				Photo:     profile.Photo,
				CreatedAt: timestamppb.New(profile.CreatedAt),
				UpdatedAt: timestamppb.New(profile.UpdatedAt),
			})
		}
	}

	return response, nil
}

func (d *ProfilesDeliveryService) Get(ctx context.Context, r *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {

	uCaseRes, err := d.profilesUCase.Get(r.UserId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetProfileResponse{
		Status: &pb.StatusResponse{
			Code:    uCaseRes.StatusCode,
			Message: uCaseRes.StatusCode,
		},
		Profile: nil,
	}

	if uCaseRes.StatusCode == domain.Success && uCaseRes.Profile != nil {
		response.Profile = &pb.Profile{
			UserId:    uCaseRes.Profile.UserId,
			Name:      uCaseRes.Profile.Name,
			Photo:     uCaseRes.Profile.Photo,
			CreatedAt: timestamppb.New(uCaseRes.Profile.CreatedAt),
			UpdatedAt: timestamppb.New(uCaseRes.Profile.UpdatedAt),
		}
	}

	return response, nil
}

func (d *ProfilesDeliveryService) Update(ctx context.Context, r *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {

	uCaseRes, err := d.profilesUCase.Update(domain.Profile{
		UserId: r.UserId,
		Name:   r.Name,
		Photo:  r.Photo,
	})
	if err != nil {
		return nil, err
	}

	response := &pb.UpdateProfileResponse{
		Status: &pb.StatusResponse{
			Code:    uCaseRes.StatusCode,
			Message: uCaseRes.StatusCode,
		},
	}

	return response, nil
}
