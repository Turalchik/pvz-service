package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"testing"
	"time"
)

func setupDataBase(t *testing.T) (*Repo, sqlmock.Sqlmock, func(), error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	db := sqlx.NewDb(sqlDB, "sqlmock")
	repo := NewRepo(db)

	return repo, mock, func() { db.Close() }, nil
}

func TestRepo_CreateUser_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	u := &users.User{
		ID:       "gajfogjadgifjgafjljo",
		Login:    "alice",
		Password: "secret",
		Role:     "user",
	}

	expectQuery := regexp.QuoteMeta(
		`INSERT INTO users VALUES ($1,$2,$3,$4)`,
	)
	mock.
		ExpectQuery(expectQuery).
		WithArgs(u.ID, u.Login, u.Password, u.Role).
		WillReturnRows(sqlmock.NewRows([]string{}))

	if err := repo.CreateUser(ctx, u); err != nil {
		t.Errorf("CreateUser returned unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations didn't met: %s", err)
	}
}

func TestRepo_CreateUser_DBError(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	u := &users.User{ID: "999", Login: "bob", Password: "pwd", Role: "admin"}

	expectQuery := regexp.QuoteMeta(
		`INSERT INTO users VALUES ($1,$2,$3,$4)`,
	)
	mock.
		ExpectQuery(expectQuery).
		WithArgs(u.ID, u.Login, u.Password, u.Role).
		WillReturnError(fmt.Errorf("some db error"))

	err = repo.CreateUser(ctx, u)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.Internal {
		t.Errorf("expected code %v, got %v", codes.Internal, st.Code())
	}
	if st.Message() != "Internal corrupt" {
		t.Errorf("expected message %q, got %q", "Internal corrupt", st.Message())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

const existExpectQuery = `SELECT EXISTS ( SELECT 1 FROM users WHERE login = $1 ) AS user_exists`

func TestCheckUserExisting_WhenExists(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "alice"

	rows := sqlmock.NewRows([]string{"user_exists"}).
		AddRow(true)
	mock.
		ExpectQuery(regexp.QuoteMeta(existExpectQuery)).
		WithArgs(login).
		WillReturnRows(rows)

	exists, err := repo.CheckUserExisting(ctx, login)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("expected exists=true, got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestCheckUserExisting_WhenNotExists(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "bob"

	rows := sqlmock.NewRows([]string{"user_exists"}).
		AddRow(false)
	mock.
		ExpectQuery(regexp.QuoteMeta(existExpectQuery)).
		WithArgs(login).
		WillReturnRows(rows)

	exists, err := repo.CheckUserExisting(ctx, login)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exists {
		t.Errorf("expected exists=false, got true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestCheckUserExisting_WhenDBError(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "charlie"

	mock.
		ExpectQuery(regexp.QuoteMeta(existExpectQuery)).
		WithArgs(login).
		WillReturnError(fmt.Errorf("db timeout"))

	exists, err := repo.CheckUserExisting(ctx, login)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if exists {
		t.Error("expected exists=false on error, got true")
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.Internal {
		t.Errorf("expected code %v, got %v", codes.Internal, st.Code())
	}
	if st.Message() != "Internal error" {
		t.Errorf("expected message %q, got %q", "Internal error", st.Message())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var getUserExpectQuery = `SELECT * FROM users WHERE login = $1`

func TestGetUserByLogin_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "alice"

	rows := sqlmock.
		NewRows([]string{"id", "login", "password", "role"}).
		AddRow("someUUID", "someLogin", "somePassword", "someRole")

	mock.
		ExpectQuery(regexp.QuoteMeta(getUserExpectQuery)).
		WithArgs(login).
		WillReturnRows(rows)

	user, err := repo.GetUserByLogin(ctx, login)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user == nil {
		t.Fatal("expected non-nil user")
	}
	if user.ID != "someUUID" {
		t.Errorf("ID = %q; want %q", user.ID, "123")
	}
	if user.Login != "someLogin" {
		t.Errorf("Login = %q; want %q", user.Login, "alice")
	}
	if user.Password != "somePassword" {
		t.Errorf("Password = %q; want %q", user.Password, "secret")
	}
	if user.Role != "someRole" {
		t.Errorf("Role = %q; want %q", user.Role, "user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGetUserByLogin_NotFound(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "notExistingLogin"

	rows := sqlmock.NewRows([]string{"id", "login", "password", "role"})

	mock.
		ExpectQuery(regexp.QuoteMeta(getUserExpectQuery)).
		WithArgs(login).
		WillReturnRows(rows) // пустой результат

	user, err := repo.GetUserByLogin(ctx, login)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound error, got: %v", st.Code())
	}
}

func TestGetUserByLogin_DBError(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	login := "someLogin"

	mock.
		ExpectQuery(regexp.QuoteMeta(getUserExpectQuery)).
		WithArgs(login).
		WillReturnError(fmt.Errorf("db timeout"))

	user, err := repo.GetUserByLogin(ctx, login)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user on error, got %+v", user)
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.Internal {
		t.Errorf("code = %v; want %v", st.Code(), codes.Internal)
	}
	if st.Message() != "Cannot take user by login" {
		t.Errorf("message = %q; want %q", st.Message(), "Cannot take user by login")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var createPVZExpectQuery = `INSERT INTO pvzs VALUES ($1,$2,$3)`

func TestRepo_CreatePVZ(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	createdPVZ := &pvz.PVZ{ID: "someUUID", RegistrationDate: time.Now(), City: "someCity"}

	mock.
		ExpectQuery(regexp.QuoteMeta(createPVZExpectQuery)).
		WithArgs(createdPVZ.ID, createdPVZ.RegistrationDate, createdPVZ.City).
		WillReturnRows(sqlmock.NewRows([]string{}))

	if err := repo.CreatePVZ(context.Background(), createdPVZ); err != nil {
		t.Errorf("CreateUser returned unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations didn't met: %s", err)
	}
}

var createReceptionExpectQuery = `INSERT INTO receptions (id,start_time,pvz_id,status) VALUES ($1,$2,$3,$4)`

func TestRepo_CreateReception(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	createdReception := &receptions.Reception{
		ID:        "someUUID",
		StartTime: time.Now(),
		PVZID:     "somePVZID",
		Status:    "someStatus",
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(createReceptionExpectQuery)).
		WithArgs(createdReception.ID, createdReception.StartTime, createdReception.PVZID, createdReception.Status).
		WillReturnRows(sqlmock.NewRows([]string{}))

	if err := repo.CreateReception(context.Background(), createdReception); err != nil {
		t.Errorf("CreateReception returned unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations didn't met: %s", err)
	}
}

var checkReceptionActiveExpectQuery = `SELECT EXISTS ( SELECT 1 FROM pvzs WHERE (id = $1 AND active_reception IS NOT NULL) ) AS has_active_reception`

func TestRepo_CheckReceptionActive_ClosedReception(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	pvzID := "somePVZID"

	mock.
		ExpectQuery(regexp.QuoteMeta(checkReceptionActiveExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(sqlmock.NewRows([]string{"has_active_reception"}).AddRow(true))

	result, err := repo.CheckReceptionActive(context.Background(), pvzID)
	if err != nil {
		t.Fatalf("CheckReceptionActive returned unexpected error: %v", err)
	}
	if !result {
		t.Fatalf("CheckReceptionActive expected true, got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations weren't met: %s", err)
	}
}

func TestRepo_CheckReceptionActive_OpenedReception(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	pvzID := "somePVZID"
	mock.
		ExpectQuery(regexp.QuoteMeta(checkReceptionActiveExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(sqlmock.NewRows([]string{"has_active_reception"}).AddRow(false))

	result, err := repo.CheckReceptionActive(context.Background(), pvzID)
	if err != nil {
		t.Fatalf("CheckReceptionActive returned unexpected error: %v", err)
	}
	if result {
		t.Fatalf("CheckReceptionActive expected false, got true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations weren't met: %s", err)
	}
}

var checkPVZExistingExpectQuery = `SELECT EXISTS ( SELECT 1 FROM pvzs WHERE id = $1 ) AS pvz_exists`

func TestCheckPVZExisting_WhenExists(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	pvzID := "someUUID"

	mock.
		ExpectQuery(regexp.QuoteMeta(checkPVZExistingExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(sqlmock.NewRows([]string{"pvz_exists"}).AddRow(true))

	exists, err := repo.CheckPVZExisting(ctx, pvzID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("expected exists=true, got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestCheckPVZExisting_WhenNotExists(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	pvzID := "someUUID"

	mock.
		ExpectQuery(regexp.QuoteMeta(checkPVZExistingExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(sqlmock.NewRows([]string{"pvz_exists"}).AddRow(false))

	exists, err := repo.CheckPVZExisting(ctx, pvzID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exists {
		t.Errorf("expected exists=false, got true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var getReceptionByPVZIDExpectQuery = `SELECT receptions.* FROM receptions JOIN pvzs ON receptions.id = pvzs.active_reception WHERE pvz.id = $1`

func TestRepo_GetReceptionByPVZID_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	pvzID := "test_pvz_id"

	expectedReception := &receptions.Reception{
		ID:            "test_reception_id",
		StartTime:     time.Now(),
		EndTime:       sql.NullTime{Time: time.Now().Add(time.Hour), Valid: true},
		PVZID:         pvzID,
		Status:        "in_progress",
		LastProductID: sql.NullString{String: "last-prod-id", Valid: true},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"start_time",
		"end_time",
		"pvz_id",
		"status",
		"last_product_id",
	}).AddRow(
		expectedReception.ID,
		expectedReception.StartTime,
		expectedReception.EndTime.Time,
		expectedReception.PVZID,
		expectedReception.Status,
		expectedReception.LastProductID,
	)

	mock.
		ExpectQuery(regexp.QuoteMeta(getReceptionByPVZIDExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(rows)

	reception, err := repo.GetReceptionByPVZID(ctx, pvzID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reception == nil {
		t.Fatal("expected non-nil reception")
	}

	if reception.ID != expectedReception.ID {
		t.Errorf("ID = %q; want %q", reception.ID, expectedReception.ID)
	}
	if reception.PVZID != expectedReception.PVZID {
		t.Errorf("PVZID = %q; want %q", reception.PVZID, expectedReception.PVZID)
	}
	if reception.Status != expectedReception.Status {
		t.Errorf("Status = %q; want %q", reception.Status, expectedReception.Status)
	}
	if reception.LastProductID.Valid != expectedReception.LastProductID.Valid ||
		reception.LastProductID.String != expectedReception.LastProductID.String {
		t.Errorf("LastProductID = %v; want %v", reception.LastProductID, expectedReception.LastProductID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestRepo_GetReceptionByPVZID_NotFound(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()
	pvzID := "test_pvz_id"

	mock.
		ExpectQuery(regexp.QuoteMeta(getReceptionByPVZIDExpectQuery)).
		WithArgs(pvzID).
		WillReturnRows(sqlmock.NewRows([]string{}))

	reception, err := repo.GetReceptionByPVZID(ctx, pvzID)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if reception != nil {
		t.Fatalf("expected nil user, got %+v", reception)
	}
	st, _ := status.FromError(err)
	if st.Code() != codes.NotFound {
		t.Fatalf("expected NotFound error, got: %v", st.Code())
	}
}

var updateLastProductExpectQuery = `UPDATE receptions SET last_product_id = $1 WHERE id = $2`

func TestRepo_UpdateReceptionLastProduct_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	receptionID := "test_reception_id"
	productID := sql.NullString{String: "test_product_id", Valid: true}

	mock.
		ExpectExec(regexp.QuoteMeta(updateLastProductExpectQuery)).
		WithArgs(productID, receptionID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateReceptionLastProduct(context.Background(), receptionID, productID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var createProductExpectQuery = `INSERT INTO products VALUES ($1,$2,$3,$4,$5)`

func TestRepo_CreateProduct_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	product := &products.Product{
		ID:                "test_product_id",
		ReceptionTime:     time.Now(),
		Type:              "электроника",
		ReceptionID:       "test_reception_id",
		PreviousProductID: sql.NullString{String: "test_prev_prod_id", Valid: true},
	}

	mock.
		ExpectExec(regexp.QuoteMeta(createProductExpectQuery)).
		WithArgs(
			product.ID,
			product.ReceptionTime,
			product.Type,
			product.ReceptionID,
			product.PreviousProductID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.CreateProduct(context.Background(), product)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var getProductByIDExpectQuery = `SELECT * FROM products WHERE id = $1`

func TestRepo_GetProductByID_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	expectedProduct := &products.Product{
		ID:                "test_product_id",
		ReceptionTime:     time.Now(),
		Type:              "электроника",
		ReceptionID:       "test_reception_id",
		PreviousProductID: sql.NullString{String: "test_previous_id", Valid: true},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"reception_time",
		"type",
		"reception_id",
		"previous_product_id",
	}).AddRow(
		expectedProduct.ID,
		expectedProduct.ReceptionTime,
		expectedProduct.Type,
		expectedProduct.ReceptionID,
		expectedProduct.PreviousProductID,
	)

	mock.
		ExpectQuery(regexp.QuoteMeta(getProductByIDExpectQuery)).
		WithArgs(expectedProduct.ID).
		WillReturnRows(rows)

	product, err := repo.GetProductByID(context.Background(), expectedProduct.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if product == nil {
		t.Fatal("expected non-nil product")
	}

	if product.ID != expectedProduct.ID {
		t.Errorf("ID = %q; want %q", product.ID, expectedProduct.ID)
	}
	if !product.ReceptionTime.Equal(expectedProduct.ReceptionTime) {
		t.Errorf("ReceptionTime = %v; want %v", product.ReceptionTime, expectedProduct.ReceptionTime)
	}
	if product.Type != expectedProduct.Type {
		t.Errorf("Type = %q; want %q", product.Type, expectedProduct.Type)
	}
	if product.ReceptionID != expectedProduct.ReceptionID {
		t.Errorf("ReceptionID = %q; want %q", product.ReceptionID, expectedProduct.ReceptionID)
	}
	if product.PreviousProductID != expectedProduct.PreviousProductID {
		t.Errorf("PreviousProductID = %v; want %v", product.PreviousProductID, expectedProduct.PreviousProductID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestRepo_GetProductByID_NotFound(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	emptyRows := sqlmock.NewRows([]string{
		"id", "reception_time", "type", "reception_id", "previous_product_id",
	})

	productID := "test_not_exist_product_id"
	mock.
		ExpectQuery(regexp.QuoteMeta(getProductByIDExpectQuery)).
		WithArgs(productID).
		WillReturnRows(emptyRows)

	product, err := repo.GetProductByID(context.Background(), productID)
	if product != nil {
		t.Errorf("expected nil product, got %+v", product)
	}
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected a gRPC status error, got %v", err)
	}
	if st.Code() != codes.NotFound {
		t.Errorf("code = %v; want %v", st.Code(), codes.NotFound)
	}
	if st.Message() != "Reception doesn't exist" {
		t.Errorf("message = %q; want %q", st.Message(), "Reception doesn't exist")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var deleteProductByIDExpectQuery = `DELETE FROM products WHERE id = $1`

func TestRepo_DeleteProductByID_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	productID := "test_product_id"
	mock.
		ExpectExec(regexp.QuoteMeta(deleteProductByIDExpectQuery)).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteProductByID(context.Background(), productID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

var closeReceptionExpectQuery = `UPDATE receptions SET end_time = $1, status = $2 WHERE id = $3`

func TestRepo_CloseReception_Success(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	receptionID := "test_reception_id"
	endTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	mock.ExpectExec(regexp.QuoteMeta(closeReceptionExpectQuery)).
		WithArgs(endTime, "close", receptionID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.CloseReception(context.Background(), receptionID, endTime)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFilteredPVZs_WhenExists(t *testing.T) {
	repo, mock, closer, err := setupDataBase(t)
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer closer()

	ctx := context.Background()

	startDate := time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 4, 30, 0, 0, 0, 0, time.UTC)
	limit := uint64(10)
	offset := uint64(0)

	rows := sqlmock.NewRows([]string{"id", "registration_date", "city", "active_reception"}).
		AddRow("test_pvz_id_1", time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC), "Москва", nil).
		AddRow("test_pvz_id_2", time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC), "Питер", sql.NullString{String: "test_reception_id", Valid: true})

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT DISTINCT pvzs.* FROM pvzs JOIN receptions ON receptions.pvz_id = pvzs.id WHERE receptions.start_time <= ? AND (receptions.end_time IS NULL OR receptions.end_time >= ?) ORDER BY pvzs.registration_date DESC LIMIT 10 OFFSET 0`,
	)).
		WithArgs(endDate, startDate).
		WillReturnRows(rows)

	pvzs, err := repo.GetFilteredPVZs(ctx, startDate, endDate, limit, offset)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pvzs) != 2 {
		t.Errorf("expected 2 pvzs, got %d", len(pvzs))
	}

	if pvzs[0].ID != "test_pvz_id_1" || pvzs[0].City != "Москва" {
		t.Errorf("unexpected first PVZ: %+v", pvzs[0])
	}
	if pvzs[1].ID != "test_pvz_id_2" || pvzs[1].City != "Питер" || !pvzs[1].ActiveReception.Valid {
		t.Errorf("unexpected second PVZ: %+v", pvzs[1])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
