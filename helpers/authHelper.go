package helpers

import (
	"errors"
	_ "fmt"
	"github.com/gin-gonic/gin"
)

func checkUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err =  errors.New("Unauthorised access to resources")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	//.......................uid = userId means the user is requesting his own data
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorised to access this resource")
		return err
	}
	err = checkUserType(c, userType)
	return err
}