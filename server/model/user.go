package model

import (
	"database/sql"
	"fmt"
	"encoding/base64"
	crand "crypto/rand"
    "math"
    "math/big"
	"math/rand"
	"io"
	"encoding/hex"
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

//Insert is a functin to insert user data to dateabase
func  Insert(db *sql.DB,user User) (*User, error) {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
    rand.Seed(seed.Int64())
	randomByte := make([]byte, 14)
	
	_, err := io.ReadFull(crand.Reader, randomByte)
	if err != nil {
		return nil,err
	}
	 salt := base64.StdEncoding.EncodeToString(randomByte)
	 converted, _ := scrypt.Key([]byte("dkjfkd"), []byte(salt), 16384, 8, 1, 16)
	 token := hex.EncodeToString(converted[:])
	_, err = db.Query(fmt.Sprintf("INSERT users (joid,name,pass) VALUES (%d,'%s','%s')", user.Joid, user.Name,user.HashedPass))
	if err != nil {
		return nil, err
	}

	return &User{
		Joid: user.Joid,
		Name: user.Name,
		HashedPass: "deadbeaf",
		Token: token,
	}, nil
}

//IsAddable is a function to check if adding new user is possible
func  IsAddable(db *sql.DB, joid int) (bool, error) {
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
		return false,nil
	}
	return true, nil
}
