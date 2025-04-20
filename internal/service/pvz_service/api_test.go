package pvz_service

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	"github.com/Turalchik/pvz-service/internal/service/pvz_service/mocks"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func makeRegisterRequest(createdUser *users.User) *desc.RegisterRequest {
	return &desc.RegisterRequest{
		Login:    createdUser.Login,
		Password: createdUser.Password,
		Role:     createdUser.Role,
	}
}

func setupPVZService(ctrl *gomock.Controller) (*mocks.MockRepoInterface, *mocks.MockUUIDInterface, desc.PVZServiceServer, error) {
	mockRepo := mocks.NewMockRepoInterface(ctrl)
	mockUUID := mocks.NewMockUUIDInterface(ctrl)
	pvzService, err := NewPVZServiceServer(mockRepo, mockUUID)
	if err != nil {
		return nil, nil, nil, err
	}
	return mockRepo, mockUUID, pvzService, nil
}

func TestPVZServiceAPI_Register_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, mockUUID, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createdUser := &users.User{ID: "someUUID", Login: "someLogin", Password: "somePassword", Role: "сотрудник ПВЗ"}

	mockRepo.EXPECT().
		CheckUserExisting(gomock.Any(), createdUser.Login).
		Return(false, nil).
		Times(1)

	mockRepo.EXPECT().
		CreateUser(gomock.Any(), createdUser).
		Return(nil).
		Times(1)

	mockUUID.EXPECT().
		NewString().
		Return(createdUser.ID).
		Times(1)

	if _, err = pvzService.Register(context.Background(), makeRegisterRequest(createdUser)); err != nil {
		t.Errorf("unexpexted error: %s", err)
	}
}

func TestPVZServiceAPI_Register_FiledByPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createdUser := &users.User{ID: "someUUID", Login: "someLogin",
		Password: "someVeryVeryVeryVeryVeryLongPassword", Role: "сотрудник ПВЗ"}

	_, err = pvzService.Register(context.Background(), makeRegisterRequest(createdUser))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.InvalidArgument {
		t.Fatalf("expected %v code error, got %v", codes.InvalidArgument, st.Code())
	}
	if st.Message() != "Invalid password length or role" {
		t.Fatalf("expected %s error, got %s", "Invalid password length or role", st.Message())
	}
}

func TestPVZServiceAPI_Register_FiledByInvalidRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createdUser := &users.User{ID: "someUUID", Login: "someLogin",
		Password: "somePassword", Role: "invalidRole"}

	_, err = pvzService.Register(context.Background(), makeRegisterRequest(createdUser))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.InvalidArgument {
		t.Fatalf("expected %v code error, got %v", codes.InvalidArgument, st.Code())
	}
	if st.Message() != "Invalid password length or role" {
		t.Fatalf("expected %s error, got %s", "Invalid password length or role", st.Message())
	}
}

func makeLoginRequest(login string, password string) *desc.LoginRequest {
	return &desc.LoginRequest{
		Login:    login,
		Password: password,
	}
}

func TestPVZServiceAPI_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	user := &users.User{ID: "someUUID", Login: "someLogin", Password: "somePassword", Role: "сотрудник ПВЗ"}

	mockRepo.EXPECT().
		GetUserByLogin(gomock.Any(), user.Login).
		Return(user, nil).
		Times(1)

	loginResponse, err := pvzService.Login(context.Background(), makeLoginRequest(user.Login, user.Password))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if loginResponse.Token == "" {
		t.Fatalf("expected token, got empty string")
	}
}

func TestPVZServiceAPI_Login_FailedByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	user := &users.User{ID: "someUUID", Login: "someLogin", Password: "somePassword", Role: "сотрудник ПВЗ"}

	mockRepo.EXPECT().
		GetUserByLogin(gomock.Any(), user.Login).
		Return(nil, status.Error(codes.NotFound, "User not found")).
		Times(1)

	loginResponse, err := pvzService.Login(context.Background(), makeLoginRequest(user.Login, user.Password))
	if loginResponse != nil {
		t.Fatalf("expected nil responce")
	}
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	st, _ := status.FromError(err)
	if st.Code() != codes.NotFound {
		t.Fatalf("expected %v code, got %v", codes.NotFound, st.Code())
	}

	if st.Message() != "User not exist" {
		t.Fatalf("expected %s message, got %s", "User not exist", st.Message())
	}
}

func TestPVZServiceAPI_Login_FailedByPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	user := &users.User{ID: "someUUID", Login: "someLogin", Password: "somePassword", Role: "сотрудник ПВЗ"}

	mockRepo.EXPECT().
		GetUserByLogin(gomock.Any(), user.Login).
		Return(user, nil).
		Times(1)

	wrongPassword := "wrongPassword"
	loginResponse, err := pvzService.Login(context.Background(), makeLoginRequest(user.Login, wrongPassword))
	if loginResponse != nil {
		t.Fatalf("expected nil responce")
	}
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	st, _ := status.FromError(err)
	if st.Code() != codes.InvalidArgument {
		t.Fatalf("expected %v code, got %v", codes.InvalidArgument, st.Code())
	}

	if st.Message() != "Wrong password" {
		t.Fatalf("expected %s message, got %s", "Wrong password", st.Message())
	}
}
