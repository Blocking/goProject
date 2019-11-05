package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db := connect()
	// defer the close till after the main function has finished
	defer db.Close()
	// executing
	//sql := "INSERT INTO `customers`.`school` (`id`, `created_by`, `created_date`, `last_modified_by`, `last_modified_date`, `version`, `address`, `name`) " +
	//	"VALUES ('6', '', '2019-04-26 18:40:57', '', '2019-04-26 18:40:57', '0', '湖北省武当山go2', '武当派go2');"
	//insert(sql,db)
	//sql := "CREATE TABLE `tags` (`id` bigint(20) NOT NULL AUTO_INCREMENT,`name` varchar(255) DEFAULT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB  DEFAULT CHARSET=utf8;"
	//exec(sql,db)
	// var sql string = "INSERT INTO `customers`.`tags` (`id`, `name`) VALUES ('3', '唐山大熊')"
	// insert(sql,db)
	rows, e := db.Query("select id,name from tags ")
	if e != nil {
		panic(e.Error())
	}
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
	}
	defer rows.Close()

}

func connect() *sql.DB {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/customers")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}
func insert(sql string, db *sql.DB) {
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
}

func exec(sql string, db *sql.DB) {
	_, e := db.Exec(sql)
	if e != nil {
		log.Fatal(e)
	}

}
