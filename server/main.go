package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

const USERNAME = "apipintek"
const PASSWORD = "secret"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Pageviews int32  `json:"pageviews"`
}

var jwtSecret []byte = []byte("apipintekidjwtcredential")

var accountsMock []User = []User{
	User{
		Id:       "1",
		Username: "nraboy",
		Password: "1234",
	},
	User{
		Id:       "2",
		Username: "mraboy",
		Password: "5678",
	},
}

var blogsMock []Blog = []Blog{
	Blog{
		Id:        "1",
		Author:    "nraboy",
		Title:     "Sample Article",
		Content:   "This is a sample article written by Nic Raboy",
		Pageviews: 1000,
	},
}

var accountType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var blogType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Blog",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"pageviews": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, err := ValidateJWT(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				return params.Source.(Blog).Pageviews, nil
			},
		},
	},
})

func ValidateJWT(t string) (interface{}, error) {
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
		var decodedToken interface{}
		mapstructure.Decode(claims, &decodedToken)
		log.Println(claims)
		return decodedToken, nil
	} else {
		return nil, errors.New("Invalid authorization token")
	}
}

func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")

	//AUTH BASIC
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

	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username, //user.Username,
		"password": password, //user.Password,
	})
	tokenString, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}

	response.Write([]byte(`{ "token": "` + tokenString + `" }`))
}

func main() {
	fmt.Println("Starting the application at :12345...")
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"account": &graphql.Field{
				Type: accountType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					_, err := ValidateJWT(params.Context.Value("token").(string))

					if err != nil {
						return nil, err
					}

					//return &User{}, nil

					//fmt.Println(account.(User).Username)
					for _, accountMock := range accountsMock {
						//log.Println(account.(Username).Username)
						//if accountMock.Username == account.(User).Username {
						return accountMock, nil
						//}
					}
					return &User{}, nil
				},
			},
			"blogs": &graphql.Field{
				Type: graphql.NewList(blogType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return blogsMock, nil
				},
			},
		},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	http.HandleFunc("/graphql", func(response http.ResponseWriter, request *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: request.URL.Query().Get("query"),
			Context:       context.WithValue(context.Background(), "token", request.URL.Query().Get("token")),
		})
		json.NewEncoder(response).Encode(result)
	})
	http.HandleFunc("/login", CreateTokenEndpoint)
	http.ListenAndServe(":12345", nil)
}
