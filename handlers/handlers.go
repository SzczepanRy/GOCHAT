package handlers

import (
	"chat/jwt"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	//check if the protocol can become a websocket one
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic("could not create a ws connection")
	}
	clients = append(clients, *wsConn)
	for {
		msgType, msg, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s send: %s\n", wsConn.RemoteAddr(), string(msg))
		//loop if found and sent to browser
		for _, client := range clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				fmt.Println(err)
				return
			}

		}
	}

}

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
			errLog(err, w)
		}
		// if hashged password and login are valid add the user to db
		//os.Getenv("HASH")
		query := fmt.Sprintf(`insert into "users" values ( '%s' , '%s' )`, data.Login, string(hash))
		fmt.Println(query)
		_, err = db.Query(query)
		if err != nil {
			errLog(err, w)
		}

		fmt.Println(data)
	} else {
		errLog(errors.New("wiorn method"), w)
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
			errLog(err, w)
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
				errLog(err, w)
			}
			values = append(values, value)
		}

		if len(values) > 1 {
			errLog(errors.New(" more than one ocurrene"), w)
		}

		login := values[0].login
		hashedPassword := values[0].password
		valid := CheckPasswordHash(data.Password, hashedPassword)
		if !valid {
			errLog(errors.New("password is not valid"), w)
		}
		token, err := jwt.CreateToken(login)

		res := resType{
			Success:     true,
			AccessToken: token,
		}

		jsonRes, err := json.Marshal(res)

		if err != nil {
			errLog(err, w)
		}
		fmt.Println("json sent successfully")
		w.Header().Set("content-type", "application/json")
		w.Write(jsonRes)
	} else {
		errLog(errors.New("wiorn method"), w)
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
			errLog(err, w)
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
		errLog(errors.New("wiorn method"), w)
	}

}

func errLog(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
