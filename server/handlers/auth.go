package handlers

import (
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"fmt"
	"log"
	"github.com/GerardSoleCa/PubKeyManager/server/responses"
	"github.com/gorilla/context"
	"github.com/GerardSoleCa/PubKeyManager/database"
	"github.com/GerardSoleCa/PubKeyManager/utils"
)

const SIGNING_KEY string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6InRlY2hAYXNiYXJjZWxvbmEuY29tIiwicm9sZSI6ImFkbWluIiwiaWF0IjoxNDU2OTIwOTYwLCJleHAiOjE0NTcwMDczNjB9.-9TfK4HWcfWJbSJmT6b2NzMBcc2EFH_KXgXOM2D3Jxk"

func ConfigureAuthHandler(router *mux.Router) {
	router.Path("/auth/login").HandlerFunc(login).Methods("POST")
	router.Path("/auth/register").HandlerFunc(register).Methods("POST")
}

func login(rw http.ResponseWriter, q *http.Request) {

}

func register(rw http.ResponseWriter, q *http.Request) {
	userCount := database.CountUsers()
	if userCount != 0 {
		responses.ErrorResponse(rw, &responses.ApiError{Code:401, Err:"User already registered"})
		return
	}
	user := &database.User{}
	if utils.ParseBody(q.Body, user) != nil {
		responses.BadRequest(rw)
		return
	}
	if err := user.Save(); err != nil {
		responses.ErrorResponse(rw, &responses.ApiError{Code: 500, Err: err.Error()})
	} else {
		responses.Created(rw)
	}

}

func TokenExistsMiddleware(rw http.ResponseWriter, q *http.Request, next http.HandlerFunc) {
	authHeader := q.Header.Get("authorization")
	token := strings.TrimPrefix(authHeader, "BEARER ")

	access, user := checkToken(token)

	if access {
		context.Set(q, "USER", user)
		next(rw, q)
	} else {
		responses.Unauthorized(rw)
		return
	}
}

func checkToken(t string) (bool, bson.M) {
	if t != "" {
		token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SIGNING_KEY), nil
		})

		if err == nil || token.Valid {
			tokenUser := bson.M{}
			claims := token.Claims.(jwt.MapClaims)
			tokenUser["role"] = claims["role"]
			tokenUser["_id"] = claims["_id"]
			tokenUser["username"] = claims["username"]
			tokenUser["email"] = claims["email"]

			return true, tokenUser
		} else {
			log.Println("Error validating token", err)
			return false, nil
		}
	} else {
		return false, nil
	}
}