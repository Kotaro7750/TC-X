package model

import(
  "database/sql"
)

//User is a structure of a user
type User struct {
	Joid string `json:"joid" `
	Name string `json:"name" `
}

// UserAll is a function to select all data from table:users
func UserAll(db *sql.DB) ([]*User,error) {
	rows,err := db.Query("SELECT joid,name FROM users")
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var users []*User
	for rows.Next(){
		user := &User{}
		if err := rows.Scan(&user.Joid,&user.Name); err != nil{
			return nil,err
		}
		users = append(users,user)
	}

	if err := rows.Err();err != nil {
		return nil,err
	}

	return users,nil
}

//Insert is a functin to insert user data to dateabase
func (u* User) Insert (db *sql.DB) (*User,error){
	_,err := db.Exec("INSERT users (joid,name) VALUES (?,?)",u.Joid,u.Name)
	if err != nil {
		return nil,err
	}

	return &User{
		Joid: u.Joid,
		Name: u.Name,
	},nil
} 
