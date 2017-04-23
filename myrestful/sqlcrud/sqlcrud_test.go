package sqlcrud

import (
	"database/sql"
	"testing"
	_ "log"
	_ "github.com/go-sql-driver/mysql"

)

func TestGetUserByID(t *testing.T) {

//	db,err	:= sql.Open("mysql", "admin:123321@tcp(139.199.152.130:3306)/mydb")
//	if err != nil{
//		t.Errorf("sql.Open error inside TestGetUserByID: %v", err.Error())

//	}


	type args struct {
		db     *sql.DB
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
		{	name:"huahuan4",want:"huahuan4",wantErr:false,
			args:args{db:func (db *sql.DB,err error) *sql.DB {return db}(sql.Open("mysql", "root:sees7&chanting@tcp(localhost:3306)/mydb")),userID:"4"},
		},
		{	name:"huahuan5",want:"huahuan5",wantErr:false,
			args:args{
				db:func (db *sql.DB,err error) *sql.DB {if err != nil{t.Errorf("sql.Open error inside TestGetUserByID: %v", err.Error())}; return db}(sql.Open("mysql", "root:sees7&chanting@tcp(localhost:3306)/mydb")),
				userID:"5",
			},
		},
		{	name:"huahuan6",want:"huahuan6",wantErr:false,
			args:args{
				db:func (db *sql.DB,err error) *sql.DB {if err != nil{t.Errorf("sql.Open error inside TestGetUserByID: %v", err.Error())}; return db}(sql.Open("mysql", "root:sees7&chanting@tcp(localhost:3306)/mydb")),
				userID:"6",
			},
		},

	}
	for _, tt := range tests {
		got, err := GetUserByID(tt.args.db, tt.args.userID)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetUserByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. GetUserByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetUsers(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := GetUsers(tt.args.db)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetUsers() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. GetUsers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCreateUserNameID(t *testing.T) {
	type args struct {
		db   *sql.DB
		name string
		id   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := CreateUserNameID(tt.args.db, tt.args.name, tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. CreateUserNameID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. CreateUserNameID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDeleteUserByID(t *testing.T) {
	type args struct {
		db     *sql.DB
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := DeleteUserByID(tt.args.db, tt.args.userID)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. DeleteUserByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. DeleteUserByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestUpdateUserNameByID(t *testing.T) {
	type args struct {
		db       *sql.DB
		userID   string
		userName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := UpdateUserNameByID(tt.args.db, tt.args.userID, tt.args.userName)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. UpdateUserNameByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. UpdateUserNameByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
