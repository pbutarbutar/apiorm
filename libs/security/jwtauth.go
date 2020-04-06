package security

import (
	"errors"
	"fmt"
	"log"
	"net/http"
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
