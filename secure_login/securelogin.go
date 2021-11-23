package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type userStruct struct {
	uuid     string
	password string
}

func getJsonData(input string) []byte {
	output, err := json.Marshal(input)
	if err != nil {
		return []byte("Error Occured")
	} else {
		return output
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//w.Write(getJsonData("This is the api for the SecureDns Project.Vist /register for registration and /login for login."))

	uuid, err := verifyJwtToken(w, r)
	if err != nil {
		if err.Error() == "token not set" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(200)
			w.Write([]byte("Welcome anonymouse user.<br> Would you like to <a href='/login'>Login</a> or <a href='/register'>Register</a>\n"))
			return
		}
		w.WriteHeader(401)
		w.Write([]byte("Invalid Jwt Token. Please login again"))
		return
	} else {

		type storage struct {
			Uuid       string `json:"uuid"`
			Email      string `json:"email"`
			Created_at string `json:"created_at"`
		}

		details := new(storage)
		stmt, err := db.Prepare("select uuid, email, created_at  from users where uuid=? ;")
		if err != nil {
			log.Println(err)
		}
		defer stmt.Close()

		getDetails, _ := stmt.Query(uuid)

		defer getDetails.Close()
		for getDetails.Next() {
			getDetails.Scan(&details.Uuid, &details.Email, &details.Created_at)
		}

		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf(`<html><h2>Hey, you're %v. <br> You have uuid =%v <br> you registered on %v</h2>`, details.Email, details.Uuid, details.Created_at)))

	}
}

func verifyJwtToken(w http.ResponseWriter, r *http.Request) (string, error) {
	bearToken := r.Header.Get("Cookie")
	if bearToken == "" {
		return "", fmt.Errorf("token not set")
	}
	//log.Println(bearToken)
	tokenString := strings.Split(bearToken, "=")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("10d61631988b57dc19e9c083d76ff45110d61631988b57dc19e9c083d76ff451"), nil
	})

	if !token.Valid && err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uuid, ok := claims["uuid"].(string)
		if !ok {
			return "", err
		}

		if !ok {
			return "", err
		}
		return uuid, nil
	} else {
		return "", fmt.Errorf("something bad happened with token")
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	if _, ok := r.PostForm["email"]; !ok {
		w.Write(getJsonData("email field not present"))
		return
	}
	if _, ok := r.PostForm["password"]; !ok {
		w.Write(getJsonData("password field not present"))
		return
	}

	user := new(userStruct)

	//query := "select uuid from mentor where men_email='" + r.PostForm["email"][0] + "';"
	stmt, err := db.Prepare("select uuid from users where email=?;")

	if err != nil {
		w.Write(getJsonData("Something went wrong"))
		return
	}
	defer stmt.Close()

	existanceCheck := stmt.QueryRow(r.PostForm["email"][0])
	existanceCheck.Scan(&user.uuid)

	if user.uuid == "" {
		hash := getPasswordHash(r.PostForm["password"][0])

		stmt, err = db.Prepare("insert into users (uuid,email,password,created_at) Values(uuid(),?,?,now());")
		if err != nil {
			log.Println(err)
			w.Write(getJsonData("Something went wrong"))
			return
		}
		defer stmt.Close()
		log.Println(stmt)
		_, err = stmt.Exec(r.PostForm["email"][0], hash)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(200)
		w.Write(getJsonData("User Registration Success"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write(getJsonData("User Already in database; login at /login"))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.PostForm)
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	if _, ok := r.PostForm["email"]; !ok {
		w.Write(getJsonData("email field not provided"))
		return
	}
	if _, ok := r.PostForm["password"]; !ok {
		w.Write(getJsonData("password field not provided"))
		return
	}
	user := new(userStruct)

	stmt, err := db.Prepare("select uuid,password from users where email=?;")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getJsonData("Something went wrong"))
		return
	}
	defer stmt.Close()

	checkPassword, _ := stmt.Query(r.PostForm["email"][0])

	for checkPassword.Next() {
		checkPassword.Scan(&user.uuid, &user.password)
	}
	checkPassword.Close()

	if user.uuid == "" {
		w.WriteHeader(401)
		w.Write(getJsonData("Invalid Username or Password"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(r.PostForm["password"][0]))

	if err == nil {
		w.WriteHeader(200)
		jwt := createJWTToken(user.uuid)
		w.Write(getJsonData(jwt))
	} else {
		w.WriteHeader(401)
		w.Write(getJsonData("Invalid Username or Password"))
	}
}

func createJWTToken(uuid string) string {
	secret := []byte("10d61631988b57dc19e9c083d76ff45110d61631988b57dc19e9c083d76ff451")
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	}
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println(err)
		return ""
	}
	return tokenString
}

func getPasswordHash(password string) string {
	bcryptHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(bcryptHashBytes)
}
