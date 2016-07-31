package user

import (
	mysql "app/backend/common/util/mysql"
	localtime "app/backend/common/util/time"
	"encoding/json"
	"fmt"
	"log"
)

const (
	USER_SELECT = "SELECT id, name, password, orgId, createdAt, modifiedAt, modifiedOp FROM user WHERE id=? "

	USER_INSERT = "INSERT INTO " +
		"user(name, password, orgId, status, createdAt, modifiedAt, modifiedOp, comment) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	USER_UPDATE = "UPDATE user SET password=?, orgId=?, modifiedAt=?, modifiedOp=? WHERE id=?"
	USER_DELETE = "UPDATE user SET status=?, modifiedAt=?, modifiedOp=? WHERE id=?"
	VALID       = 1
	INVALID     = 0
)

type User struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	OrgId      int32  `json:"orgId"`
	Password   string `json:"password"`
	Status     int32  `json:"status"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedOp int32  `json:"modifiedOp"`
	Comment    string `json:"comment"`
}

func NewUser(name, password, comment string, orgId, status, modifiedOp int32) *User {

	return &User{
		Name:       name,
		Password:   password,
		OrgId:      orgId,
		Status:     status,
		Comment:    comment,
		ModifiedAt: localtime.NewLocalTime().String(),
		CreatedAt:  localtime.NewLocalTime().String(),
		ModifiedOp: modifiedOp,
	}
}

func (u *User) QueryUserById(id int32) {
	db := mysql.MysqlInstance().Conn()

	// Prepare select-statement
	stmt, err := db.Prepare(USER_SELECT)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	// Query user by id
	err = stmt.QueryRow(id).Scan(&u.Id, &u.Name, &u.Password, &u.OrgId,
		&u.CreatedAt, &u.ModifiedAt, &u.ModifiedOp)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	fmt.Printf("%v\n", u)
}

func (u *User) InsertUser(op int32) {
	db := mysql.MysqlInstance().Conn()

	// Prepare insert-statement
	stmt, err := db.Prepare(USER_INSERT)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	// Update createdAt, modifiedAt, modifiedOp
	u.CreatedAt = localtime.NewLocalTime().String()
	u.ModifiedAt = localtime.NewLocalTime().String()
	u.ModifiedOp = op

	// Insert a user
	_, err = stmt.Exec(u.Name, u.Password, u.OrgId, u.Status,
		u.CreatedAt, u.ModifiedAt, u.ModifiedOp, u.Comment)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())

	}
}

func (u *User) UpdateUser(op int32) {

	db := mysql.MysqlInstance().Conn()

	// Prepare update-statement
	stmt, err := db.Prepare(USER_UPDATE)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer stmt.Close()

	// Update modifiedAt, modifiedOp
	u.ModifiedAt = localtime.NewLocalTime().String()
	u.ModifiedOp = op

	// Update a user: password or orgId
	_, err = stmt.Exec(u.Password, u.OrgId, u.ModifiedAt, u.ModifiedOp, u.Id)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func (u *User) DeleteUser(op int32) {
	db := mysql.MysqlInstance().Conn()

	// Prepare delete-statement
	stmt, err := db.Prepare(USER_DELETE)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	defer stmt.Close()

	// Update modifiedAt and modifiedOp
	u.ModifiedAt = localtime.NewLocalTime().String()
	u.ModifiedOp = op

	// Set user status  INVALID
	u.Status = INVALID
	_, err = stmt.Exec(u.Status, u.ModifiedAt, u.ModifiedOp, u.Id)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func (u *User) DecodeJson(data string) {
	err := json.Unmarshal([]byte(data), u)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func (u *User) EncodeJson() string {
	data, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	return string(data)
}
