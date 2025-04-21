package pvz_service

import (
	"context"
	"database/sql"
	"github.com/Turalchik/pvz-service/internal/app/tokenizer"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	"github.com/Turalchik/pvz-service/internal/service/pvz_service/mocks"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func makeRegisterRequest(createdUser *users.User) *desc.RegisterRequest {
	return &desc.RegisterRequest{
		Login:    createdUser.Login,
		Password: createdUser.Password,
		Role:     createdUser.Role,
	}
}

func setupPVZService(ctrl *gomock.Controller) (*mocks.MockRepoInterface, *mocks.MockUUIDInterface, *mocks.MockTimerInterface, *mocks.MockTokenizerInterface, desc.PVZServiceServer, error) {
	mockRepo := mocks.NewMockRepoInterface(ctrl)
	mockUUID := mocks.NewMockUUIDInterface(ctrl)
	mockTimer := mocks.NewMockTimerInterface(ctrl)
	mockTokenizer := mocks.NewMockTokenizerInterface(ctrl)
	pvzService, err := NewPVZServiceServer(mockRepo, mockUUID, mockTimer, mockTokenizer)

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return mockRepo, mockUUID, mockTimer, mockTokenizer, pvzService, nil
}

func TestPVZServiceAPI_Register_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, mockUUID, _, _, pvzService, err := setupPVZService(ctrl)
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

	_, _, _, _, pvzService, err := setupPVZService(ctrl)
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

	_, _, _, _, pvzService, err := setupPVZService(ctrl)
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

	mockRepo, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	user := &users.User{ID: "someUUID", Login: "someLogin", Password: "somePassword", Role: "сотрудник ПВЗ"}

	mockRepo.EXPECT().
		GetUserByLogin(gomock.Any(), user.Login).
		Return(user, nil).
		Times(1)

	mockTokenizer.EXPECT().
		GenerateToken(user.ID, user.Role).
		Return("someToken", nil).
		Times(1)

	loginResponse, err := pvzService.Login(context.Background(), makeLoginRequest(user.Login, user.Password))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if loginResponse.Token != "someToken" {
		t.Fatalf("expected token, got empty string")
	}
}

func TestPVZServiceAPI_Login_FailedByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, _, _, pvzService, err := setupPVZService(ctrl)
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

	mockRepo, _, _, _, pvzService, err := setupPVZService(ctrl)
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

func makeCreatePVZRequest(pvz *pvz.PVZ) *desc.CreatePVZRequest {
	return &desc.CreatePVZRequest{
		Token: "someToken",
		City:  pvz.City,
	}
}

func TestPVZServiceAPI_CreatedPVZ_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, mockUUID, mockTimer, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createdPVZ := &pvz.PVZ{ID: "someUUID", RegistrationDate: time.Now(), City: "Москва"}
	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "модератор",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	mockTimer.EXPECT().
		Now().
		Return(createdPVZ.RegistrationDate).
		Times(1)

	mockRepo.EXPECT().
		CreatePVZ(gomock.Any(), createdPVZ).
		Return(nil).
		Times(1)

	mockUUID.EXPECT().
		NewString().
		Return(createdPVZ.ID).
		Times(1)

	createPVZResponse, err := pvzService.CreatePVZ(context.Background(), makeCreatePVZRequest(createdPVZ))
	if err != nil {
		t.Fatalf("unexpexted error: %s", err)
	}
	if createdPVZ.ID != createPVZResponse.Id {
		t.Errorf("expected %s uuid, got %s", createdPVZ.ID, createPVZResponse.Id)
	}
}

func TestPVZServiceAPI_CreatedPVZ_FailedByRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createdPVZ := &pvz.PVZ{ID: "someUUID", RegistrationDate: time.Now(), City: "Москва"}
	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "сотрудник ПВЗ",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	createPVZResponse, err := pvzService.CreatePVZ(context.Background(), makeCreatePVZRequest(createdPVZ))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if createPVZResponse != nil {
		t.Fatalf("expected nil, got response")
	}
	st, _ := status.FromError(err)

	if st.Code() != codes.PermissionDenied {
		t.Fatalf("expected %v code, got %v", codes.PermissionDenied, st.Code())
	}
	if st.Message() != "Insufficient permissions" {
		t.Fatalf("expected %s message, got %s", "Insufficient permissions", st.Message())
	}
}

func makeOpenReceptionRequest(pvzID string) *desc.OpenReceptionRequest {
	return &desc.OpenReceptionRequest{
		Token: "someToken",
		Id:    pvzID,
	}
}

func TestPVZServiceAPI_OpenReception_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, mockUUID, mockTimer, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	openedReception := &receptions.Reception{
		ID:        "someUUID",
		StartTime: time.Now(),
		PVZID:     "somePVZUUID",
		Status:    "in_progress",
	}
	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "сотрудник ПВЗ",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckPVZExisting(gomock.Any(), openedReception.PVZID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckReceptionActive(gomock.Any(), openedReception.PVZID).
		Return(false, nil).
		Times(1)

	mockUUID.EXPECT().
		NewString().
		Return(openedReception.ID).
		Times(1)

	mockTimer.EXPECT().
		Now().
		Return(openedReception.StartTime).
		Times(1)

	mockRepo.EXPECT().
		UpdateActiveReceptionPVZ(gomock.Any(), openedReception.PVZID, sql.NullString{String: openedReception.ID, Valid: true}).
		Return(nil).
		Times(1)

	mockRepo.EXPECT().
		CreateReception(gomock.Any(), openedReception).
		Return(nil).
		Times(1)

	openReceptionResponse, err := pvzService.OpenReception(context.Background(), makeOpenReceptionRequest(openedReception.PVZID))
	if err != nil {
		t.Fatalf("unexpexted error: %s", err)
	}
	if openReceptionResponse.IdReception != openedReception.ID {
		t.Errorf("expected %s uuid, got %s", openedReception.ID, openReceptionResponse.IdReception)
	}
}

func TestPVZServiceAPI_OpenReception_FailedByRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "модератор",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	openReceptionResponse, err := pvzService.OpenReception(context.Background(), makeOpenReceptionRequest("somePVZUUID"))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if openReceptionResponse != nil {
		t.Fatalf("expected nil, got response")
	}
	st, _ := status.FromError(err)

	if st.Code() != codes.PermissionDenied {
		t.Fatalf("expected %v code, got %v", codes.PermissionDenied, st.Code())
	}
	if st.Message() != "Insufficient permissions" {
		t.Fatalf("expected %s message, got %s", "Insufficient permissions", st.Message())
	}
}

func makeAddProductRequest(token string, pvzID string, productType string) *desc.AddProductRequest {
	return &desc.AddProductRequest{
		Token: token,
		Id:    pvzID,
		Type:  productType,
	}
}

func TestPVZServiceAPI_AddProduct_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, mockUUID, mockTimer, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	pvzID := "test_pvz_id"
	receptionID := "test_reception_id"
	productID := "test_product_id"
	now := time.Now()

	gotReception := &receptions.Reception{
		ID:            receptionID,
		StartTime:     now.Add(-time.Hour),
		EndTime:       sql.NullTime{},
		PVZID:         pvzID,
		Status:        "открыта",
		LastProductID: sql.NullString{String: "prev_id", Valid: true},
	}

	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "сотрудник ПВЗ",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckPVZExisting(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckReceptionActive(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		GetReceptionByPVZID(gomock.Any(), pvzID).
		Return(gotReception, nil).
		Times(1)

	mockUUID.EXPECT().
		NewString().
		Return(productID).
		Times(1)

	mockTimer.EXPECT().
		Now().
		Return(now).
		Times(1)

	mockRepo.EXPECT().
		UpdateReceptionLastProduct(gomock.Any(), receptionID, sql.NullString{String: productID, Valid: true}).
		Return(nil).
		Times(1)

	mockRepo.EXPECT().
		CreateProduct(gomock.Any(), gomock.AssignableToTypeOf(&products.Product{})).
		DoAndReturn(func(ctx context.Context, p *products.Product) error {
			if p.ID != productID {
				t.Errorf("unexpected product ID: got %s, want %s", p.ID, productID)
			}
			if p.ReceptionID != receptionID {
				t.Errorf("unexpected reception ID: got %s, want %s", p.ReceptionID, receptionID)
			}
			if p.PreviousProductID.String != "prev_id" {
				t.Errorf("unexpected previous product ID: got %v, want %v", p.PreviousProductID.String, "prev_id")
			}
			return nil
		}).
		Times(1)

	resp, err := pvzService.AddProduct(context.Background(), makeAddProductRequest("someToken", pvzID, "электроника"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil || resp.GetIdProduct() != productID {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestPVZServiceAPI_AddProduct_FailedByRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "модератор",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	pvzID := "test_pvz_id"

	addProductResponse, err := pvzService.AddProduct(context.Background(), makeAddProductRequest("someToken", pvzID, "электроника"))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if addProductResponse != nil {
		t.Fatalf("expected nil, got response")
	}
	st, _ := status.FromError(err)

	if st.Code() != codes.PermissionDenied {
		t.Fatalf("expected %v code, got %v", codes.PermissionDenied, st.Code())
	}
	if st.Message() != "Insufficient permissions" {
		t.Fatalf("expected %s message, got %s", "Insufficient permissions", st.Message())
	}
}

func makeRemoveProductRequest(token string, pvzID string) *desc.RemoveProductRequest {
	return &desc.RemoveProductRequest{
		Token: token,
		Id:    pvzID,
	}
}

func TestPVZServiceAPI_RemoveProduct_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	pvzID := "test_pvz_id"
	receptionID := "test_reception_id"
	lastProductID := "last_prod_id"
	prevProductID := "prev_prod_id"

	gotReception := &receptions.Reception{
		ID:            receptionID,
		StartTime:     time.Now(),
		PVZID:         pvzID,
		Status:        "in_progress",
		LastProductID: sql.NullString{String: lastProductID, Valid: true},
	}
	gotProduct := &products.Product{
		ID:                lastProductID,
		ReceptionTime:     time.Now(),
		Type:              "электроника",
		ReceptionID:       receptionID,
		PreviousProductID: sql.NullString{String: prevProductID, Valid: true},
	}
	claims := &tokenizer.Claims{UserID: "someUserUUID", Role: "сотрудник ПВЗ"}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckPVZExisting(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckReceptionActive(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		GetReceptionByPVZID(gomock.Any(), pvzID).
		Return(gotReception, nil).
		Times(1)

	mockRepo.EXPECT().
		GetProductByID(gomock.Any(), lastProductID).
		Return(gotProduct, nil).
		Times(1)

	mockRepo.EXPECT().
		UpdateReceptionLastProduct(gomock.Any(), receptionID, sql.NullString{String: prevProductID, Valid: true}).
		Return(nil).
		Times(1)

	mockRepo.EXPECT().
		DeleteProductByID(gomock.Any(), lastProductID).
		Return(nil).
		Times(1)

	_, err = pvzService.RemoveProduct(context.Background(), makeRemoveProductRequest("someToken", pvzID))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPVZServiceAPI_RemoveProduct_FailedByRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	claims := &tokenizer.Claims{
		UserID: "someUUID",
		Role:   "модератор",
	}

	mockTokenizer.EXPECT().
		VerifyToken("someToken").
		Return(claims, nil).
		Times(1)

	pvzID := "test_pvz_id"

	_, err = pvzService.RemoveProduct(context.Background(), makeRemoveProductRequest("someToken", pvzID))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	st, _ := status.FromError(err)

	if st.Code() != codes.PermissionDenied {
		t.Fatalf("expected %v code, got %v", codes.PermissionDenied, st.Code())
	}
	if st.Message() != "Insufficient permissions" {
		t.Fatalf("expected %s message, got %s", "Insufficient permissions", st.Message())
	}
}

func TestPVZServiceAPI_CloseReception_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, mockTimer, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("setupPVZService failed: %v", err)
	}

	pvzID := "test_pvz_id"
	receptionID := "test_reception_id"
	now := time.Now()

	claims := &tokenizer.Claims{
		UserID: "user123",
		Role:   "сотрудник ПВЗ",
	}

	mockTokenizer.EXPECT().
		VerifyToken("valid_token").
		Return(claims, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckPVZExisting(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		CheckReceptionActive(gomock.Any(), pvzID).
		Return(true, nil).
		Times(1)

	mockRepo.EXPECT().
		UpdateActiveReceptionPVZ(gomock.Any(), pvzID, sql.NullString{}).
		Return(nil).
		Times(1)

	mockRepo.EXPECT().
		GetReceptionByPVZID(gomock.Any(), pvzID).
		Return(&receptions.Reception{
			ID: receptionID,
		}, nil).
		Times(1)

	mockTimer.EXPECT().
		Now().
		Return(now).
		Times(1)

	mockRepo.EXPECT().
		CloseReception(gomock.Any(), receptionID, sql.NullTime{Time: now, Valid: true}).
		Return(nil).
		Times(1)

	req := &desc.CloseReceptionRequest{
		Token: "valid_token",
		Id:    pvzID,
	}

	resp, err := pvzService.CloseReception(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if resp != nil {
		t.Fatalf("expected nil response (Empty), got: %v", resp)
	}
}

func TestPVZServiceAPI_GetFilteredPVZs_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo, _, _, mockTokenizer, pvzService, err := setupPVZService(ctrl)
	if err != nil {
		t.Fatalf("setupPVZService failed: %v", err)
	}

	start := time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)
	finish := time.Date(2024, 4, 30, 23, 59, 59, 0, time.UTC)
	limit := uint64(5)
	offset := uint64(10)
	req := &desc.GetFilteredPVZsRequest{
		Token:  "test_token",
		Start:  timestamppb.New(start),
		Finish: timestamppb.New(finish),
		Limit:  limit,
		Offset: offset,
	}

	claims := &tokenizer.Claims{UserID: "user1", Role: "сотрудник ПВЗ"}
	mockTokenizer.EXPECT().
		VerifyToken("test_token").
		Return(claims, nil).
		Times(1)

	domainPVZs := []*pvz.PVZ{
		{
			ID:               "pvz-1",
			RegistrationDate: time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
			City:             "Москва",
			ActiveReception:  sql.NullString{String: "rec-1", Valid: true},
		},
		{
			ID:               "pvz-2",
			RegistrationDate: time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
			City:             "Питер",
			ActiveReception:  sql.NullString{Valid: false},
		},
	}
	mockRepo.EXPECT().
		GetFilteredPVZs(gomock.Any(), start, finish, limit, offset).
		Return(domainPVZs, nil).
		Times(1)

	resp, err := pvzService.GetFilteredPVZs(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response")
	}

	if len(resp.GetPvzs()) != 2 {
		t.Fatalf("expected 2 pvzs, got %d", len(resp.GetPvzs()))
	}

	first := resp.GetPvzs()[0]
	if first.GetId() != "pvz-1" {
		t.Errorf("first.PVZ.Id = %q; want %q", first.GetId(), "pvz-1")
	}
	if first.GetCity() != "Москва" {
		t.Errorf("first.PVZ.City = %q; want %q", first.GetCity(), "Москва")
	}

	second := resp.GetPvzs()[1]
	if second.GetId() != "pvz-2" {
		t.Errorf("second.PVZ.Id = %q; want %q", second.GetId(), "pvz-2")
	}
	if second.GetCity() != "Питер" {
		t.Errorf("second.PVZ.City = %q; want %q", second.GetCity(), "Питер")
	}
}
