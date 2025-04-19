package repo

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"testing"
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

	// Вернём одну строку с true
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
