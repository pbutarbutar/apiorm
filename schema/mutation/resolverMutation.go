package mutation

import (
	"crypto/sha1"
	"fmt"
	"log"
	"pintekid/config"
	"pintekid/libs/security"
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
	"gopkg.in/go-playground/validator.v9"
)

/*func CreateProductMutation(param graphql.ResolveParams) (interface{}, error) {
	idpro := param.Args["ID_PRO"].(string)
	nama := param.Args["PRO_NAME"].(string)
	qty := param.Args["QTE_IN_STOCK"].(int)
	img := param.Args["PRO_IMG"].(string)
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query("INSERT INTO Products values (?,?,?,?)", idpro, nama, qty, img)
	if err != nil {
		panic(err.Error())
	}

	return nil, err
}*/

func SignupRegistrantMutation(param graphql.ResolveParams) (interface{}, error) {
	_, errsecure := security.ValidateJWT(param.Context.Value("token").(string))
	if errsecure != nil {
		return nil, errsecure
	}

	var password = param.Args["password"].(string)
	var sha = sha1.New()
	sha.Write([]byte(password))
	var passwordEncrypted = sha.Sum(nil)
	var stringpassencrypted = fmt.Sprintf("%x", passwordEncrypted)
	user := types.UserRegistrant{
		RegistrantId: 1,
		Email:        param.Args["email"].(string),
		Fullname:     param.Args["fullname"].(string),
		PhoneNo:      param.Args["phone_no"].(string),
		Password:     stringpassencrypted,
	}

	log.Println(user)

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		panic(err.Error())
	}

	//Input to table Registrant
	registrantid, err2 := InsertRegistrant(user)
	if err2 != nil {
		panic(err2.Error())
	}
	log.Println(registrantid)
	user.RegistrantId = registrantid
	user.Success = true
	return user, err
}

func InsertRegistrant(user types.UserRegistrant) (int64, error) {

	db, err := config.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	streginstrant, err := tx.Prepare("INSERT registrant(fullname,email,phone) VALUES(?,?,?) ")
	if err != nil {
		return -1, err
	}
	defer streginstrant.Close()

	res, err := streginstrant.Exec(user.Fullname, user.Email, user.PhoneNo)
	if err != nil {
		return -1, err
	}

	lastid, _ := res.LastInsertId()

	//Set Password
	streginstrantsecret, err2 := tx.Prepare("INSERT registrant_secret(registrant_id,password,updated_at) VALUES(?,?,Now()) ")
	if err2 != nil {
		return -1, err2
	}
	defer streginstrantsecret.Close()

	_, err3 := streginstrantsecret.Exec(lastid, user.Password)
	if err3 != nil {
		return -1, err3
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return lastid, nil
}
