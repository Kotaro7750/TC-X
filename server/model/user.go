package model

import (
	crand "crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/scrypt"
)

//User is a structure of a user
type User struct {
	Joid       int    `json:"joid" binding:"required"`
	Name       string `json:"name" binding:"required"`
	HashedPass string `json:"hashedPass" binding:"required"`
	Token      string `json:"token" `
}

// UserAll is a function to select all data from table:users
func UserAll(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT joid,name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.Joid, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//UserAdd is a functin to insert user data to dateabase
func UserAdd(db *sql.DB, user User) (*User, error) {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	randomByte := make([]byte, 14)

	_, err := io.ReadFull(crand.Reader, randomByte)
	if err != nil {
		return nil, err
	}
	salt := base64.StdEncoding.EncodeToString(randomByte)
	converted, _ := scrypt.Key([]byte("dkjfkd"), []byte(salt), 16384, 8, 1, 16)
	token := hex.EncodeToString(converted[:])

	_, err = db.Query(fmt.Sprintf("INSERT users (joid,name,pass,token) VALUES (%d,'%s','%s','%s')", user.Joid, user.Name, user.HashedPass, token))
	if err != nil {
		return nil, err
	}

	return &User{
		Joid:       user.Joid,
		Name:       user.Name,
		HashedPass: "deadbeaf",
		Token:      token,
	}, nil
}

//IsAddable is a function to check if adding new user is possible
func IsAddable(db *sql.DB, joid int) (bool, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT joid,name,pass FROM `users` WHERE joid='%d'", joid))
	if err != nil {
		return false, err
	}
	defer rows.Close()

	affectedRows := 0
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Joid, &user.Name, &user.HashedPass); err != nil {
			return false, err
		}
		affectedRows = affectedRows + 1
	}

	if err := rows.Err(); err != nil {
		return false, err
	}

	if affectedRows != 0 {
		return false, nil
	}
	return true, nil
}

//AuthUser is a  function to authenticate user by joid and hashedPass
func AuthUser(db *sql.DB, joid int, hashedPass string) (*User, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT joid,name,token FROM `users` WHERE joid = %d AND pass = '%s'", joid, hashedPass))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	affectedRows := 0
	for rows.Next() {
		if err := rows.Scan(&user.Joid, &user.Name, &user.Token); err != nil {
			return nil, err
		}
		affectedRows = affectedRows + 1
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if affectedRows >= 2 {
		err := errors.New("database unique column does not work")
		return nil, err
	}
	if affectedRows == 0 {
		return nil, nil
	}

	return &User{
		Joid:       user.Joid,
		Name:       user.Name,
		HashedPass: "deadbeaf",
		Token:      user.Token,
	}, nil
}

//UserListString is a function to return [[joid,name]...[]] from users
func UserListString(db *sql.DB) ([][2]string, error) {
	tableName := "users"
	rows, err := db.Query(fmt.Sprintf("SELECT joid,name FROM %s", tableName))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList [][2]string
	for rows.Next() {
		var joid int
		var name string
		if err := rows.Scan(&joid, &name); err != nil {
			return nil, err
		}
		userList = append(userList, [2]string{strconv.Itoa(joid), name})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userList, nil
}
