package handlers

import (
	"chat/jwt"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)
func HandleFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/login.html")

}

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type loginBody struct {
	Login    string `json:"login"`
	Password string `josn:"password"`
}

func Redgister(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data loginBody
		err := json.NewDecoder(r.Body).Decode(&data)
		//hash password
		hash, err := HashPassword(data.Password)
		if err != nil {
			ErrLog(err, w)
		}
		// if hashged password and login are valid add the user to db
		//os.Getenv("HASH")
		query := fmt.Sprintf(`insert into "users" values ( '%s' , '%s' )`, data.Login, string(hash))
		fmt.Println(query)
		_, err = db.Query(query)
		if err != nil {
			ErrLog(err, w)
		}

		fmt.Println(data)
	} else {
		ErrLog(errors.New("wiorn method"), w)
	}
}

type usersRow struct {
	login    string
	password string
}

type resType struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"accessToken"`
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data loginBody
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			ErrLog(err, w)
		}
		query := fmt.Sprintf(`select * from "users" where login='%s'`, data.Login)
		fmt.Println(query)
		rows, err := db.Query(query)
		defer rows.Close()
		values := []usersRow{}
		for rows.Next() {
			value := usersRow{}
			err := rows.Scan(&value.login, &value.password)
			if err != nil {
				ErrLog(err, w)
			}
			values = append(values, value)
		}

		if len(values) > 1 {
			ErrLog(errors.New(" more than one ocurrene"), w)
		}

		login := values[0].login
		hashedPassword := values[0].password
		valid := CheckPasswordHash(data.Password, hashedPassword)
		if !valid {
			ErrLog(errors.New("password is not valid"), w)
		}
		token, err := jwt.CreateToken(login)

		res := resType{
			Success:     true,
			AccessToken: token,
		}

		jsonRes, err := json.Marshal(res)

		if err != nil {
			ErrLog(err, w)
		}
		fmt.Println("json sent successfully")
		w.Header().Set("content-type", "application/json")
		w.Write(jsonRes)
	} else {
		ErrLog(errors.New("wrong method"), w)
	}

}

type validateBody struct {
	AccessToken string `json:"accessToken"`
}

type validBody struct {
	OK bool `json:"ok"`
}

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data validateBody
		var res validBody
        err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			ErrLog(err, w)
		}
        err = jwt.VerifyToken(data.AccessToken)
		if err != nil {
            res.OK = false
		}else{

            res.OK = true
        }
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else {
		ErrLog(errors.New("wrong method"), w)
	}

}

func ErrLog(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
