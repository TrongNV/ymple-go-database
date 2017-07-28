package package_mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

//db, err := sql.Open("mysql", "user:password@/dbname")

const (
	DB_HOST = "tcp(db.penteres.com:3306)"
	DB_NAME = ""
	DB_USER ="" /*"root"*/
	//	DB_PASS = /*""*/
)

func Read() {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var str string
	//q := "select username from bankadmin.accounts where id = 1"

	q := "SELECT count(*) FROM democli_timenet_it.adoc";
	err = db.QueryRow(q).Scan(&str)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(str)
}
