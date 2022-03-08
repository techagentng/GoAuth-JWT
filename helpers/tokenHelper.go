package helpers

import (
	"github.com/dgrijalva/jwt-go"
	_"github.com/gin-gonic/gin"
	_"context"
	"github.com/techagentng/GOAUTH-JJWT/database"
	"log"
	"time"
	"os"
	_"fmt"
	_ "github.com/techagentng/GOAUTH-JJWT/database"
	_ "github.com/dgrijalva/jwt-go"
	_"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	 Email string
	 First_name string
	 Last_name string
	 Uid string
	 User_type string
	 jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (SignedToken string, refreshToken string, err error){
	claims := &SignedDetails{
		Email : email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln(err)
		return
	}
	return token, refreshToken, err
}
