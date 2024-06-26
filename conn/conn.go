package conn

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


func ConnectToDB() *sql.DB {
    fmt.Println("connecttoDb")

    var (
        host="5.185.3.151"
        port=5432
        user="postgres"
        password= os.Getenv("DB_PASS")
        dbname="goChat"
    )


	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db

}
