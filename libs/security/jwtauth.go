package security

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"pintekid/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

const USERNAME = "apipintek"
const PASSWORD = "secret"
const JWTSECRET = "pintekapisecret"

var jwtSecret []byte = []byte(JWTSECRET)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SiginRegistrant struct {
	RegistrantId int64  `json:"registrant_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	username, password, ok := request.BasicAuth()
	if !ok {
		response.Write([]byte(`{ "token": "","msg":"Not Found Auth" }`))
		return
	}
	isValid := (username == USERNAME && password == PASSWORD)
	if !isValid {
		response.Write([]byte(`{ "token": "","msg":"Wrong username or password." }`))
		return
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["password"] = password
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	/*jwt.MapClaims{
		"username": username, //user.Username,
		"password": password, //user.Password,
		//"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})*/
	tokenString, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}

	response.Write([]byte(`{ "token": "` + tokenString + `" }`))
}

func CreateTokenUserRegistrantEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	//log.Println(request.Body)
	var Duser SiginRegistrant
	if request.Method == "POST" {
		if request.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(request.Body)
			if err := decodeJSON.Decode(&Duser); err != nil {
				//log.Fatal(err)
				response.Write([]byte(`{ "token": "","msg":"Paramater Body Not Found" }`))
				return
			}
		} else {
			email := request.PostFormValue("email")
			password := request.PostFormValue("password")
			Duser = SiginRegistrant{
				Email:    email,
				Password: password,
			}
		}
	}
	log.Println(Duser)
	var sha = sha1.New()
	sha.Write([]byte(Duser.Password))
	var passwordEncrypted = sha.Sum(nil)
	var stringpassencrypted = fmt.Sprintf("%x", passwordEncrypted)

	log.Println(stringpassencrypted)

	db, _ := config.GetConnection()
	var Registrant_id int64

	query := `
				SELECT b.registrant_id FROM registrant a 
				inner join registrant_secret b on a.id = b.registrant_id
				WHERE a.email = ? AND b.password = ?
			`

	err2 := db.QueryRow(query, Duser.Email, stringpassencrypted).Scan(&Registrant_id)

	Duser.RegistrantId = Registrant_id

	if err2 != nil {
		log.Println(err2)
		response.Write([]byte(`{ "token_api": "","msg":"Wrong username or password." }`))
		return
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["registrant_id"] = Duser.RegistrantId
	atClaims["username"] = Duser.Email
	atClaims["password"] = stringpassencrypted
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	/*jwt.MapClaims{
		"username": username, //user.Username,
		"password": password, //user.Password,
		//"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})*/
	tokenString, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}
	response.Write([]byte(`{ "token_app": "` + tokenString + `" }`))
}

func ValidateJWT(t string) (*User, error) {
	if t == "" {
		return nil, errors.New("Authorization token must be present")
	}
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//var decodedToken interface{}
		var user User
		mapstructure.Decode(claims, &user)
		/*err := mapstructure.Decode(claims, &decodedToken)
		if err != nil {
			panic(err)
		}*/
		log.Println(user.Username)
		return &user, nil
	} else {
		return nil, errors.New("Invalid authorization token")
	}
}
