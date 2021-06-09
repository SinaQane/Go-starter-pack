package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbUser     = ""
	dbPassword = ""
	dbAddress  = "127.0.0.1"
	dbName     = ""
)

const data = `[{"name": "first", "age": 18},
			   {"name": "second", "age": 20},
			   {"name": "third", "age": 25}]`

type User struct {
	Id   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func unmarshalData() (users []User) {
	err := json.Unmarshal([]byte(data), &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func main() {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", dbUser, dbPassword, dbAddress, dbName)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	// driver code
}

func insertData(db *sql.DB) error {
	users := unmarshalData()

	for _, user := range users {
		query := "INSERT INTO `user` (name, age) VALUES (?, ?)"

		insertion, err := db.Prepare(query)

		if err != nil {
			return err
		}

		_, err = insertion.Exec(user.Name, user.Age)

		if err != nil {
			return err
		}

		err = insertion.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func deleteData(db *sql.DB, id int) error {
	query := "DELETE FROM `user` WHERE `id` = ?"

	deletion, err := db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = deletion.Exec(id)

	if err != nil {
		return err
	}

	err = deletion.Close()

	if err != nil {
		return err
	}

	return nil
}

func updateData(db *sql.DB, id int, name string, age int) error {
	query := "UPDATE `user` SET `name` = ?, `age` = ? WHERE `id` = ?"

	update, err := db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = update.Exec(name, age, id)

	if err != nil {
		return err
	}

	err = update.Close()

	if err != nil {
		return err
	}

	return nil
}

func getData(db *sql.DB) (users []User, err error) {
	resp, err := db.Query("SELECT * FROM `user`")

	defer func(resp *sql.Rows) {
		err := resp.Close()
		if err != nil {}
	}(resp)

	if err != nil {
		return users, err
	}

	for resp.Next() {
		var pUser User
		err = resp.Scan(&pUser.Id, &pUser.Name, &pUser.Age)
		if err != nil {
			return users, err
		}
		users = append(users, pUser)
	}

	return users, nil
}
