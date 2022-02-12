package controllers

import (
	"context"
	"github.com/techagentng/GOAUTH-JJWT/database"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	helper "github.com/techagentng/GOAUTH-JJWT/helpers"
	"github.com/techagentng/GOAUTH-JJWT/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Signup(){

}
func Login(){

}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context){
		userId := c.Param("user_id")
		helper.
	}
}
func GetUserID()  {

}

func HashPassword(){

}

func VerifyPassword()  {

}
