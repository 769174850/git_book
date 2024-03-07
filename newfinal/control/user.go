package control

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"newfinal/dao"
	"newfinal/model"
	"newfinal/util"
	"strconv"
	"time"
)

var jwtSecret = []byte("you-bing-bao-dao")

func generateToken(userID int64) (string, string, error) {
	tokenExpire := time.Now().Add(2 * time.Hour).Unix()         //token过期时间
	refreshTokenExpire := time.Now().Add(12 * time.Hour).Unix() //refreshToken过期时间

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     tokenExpire,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     refreshTokenExpire,
	})

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return tokenStr, refreshTokenStr, nil
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	userID, err := dao.VerifyUserAndGetID(loginRequest.Username, loginRequest.Password)
	if err != nil {
		util.RespErr(c)
		return
	}
	if userID == 0 {
		util.RespErr(c)
		return
	}

	token, refreshToken, err := generateToken(userID)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, gin.H{
		"refresh_token": refreshToken,
		"token":         token,
	})
	//users, err := dao.GetUser()
	//if err != nil {
	//	util.RespInternalErr(c) //返回500状态码错误
	//	return
	//}
	//
	//for _, user := range users {
	//	if user.Username == loginRequest.Username && user.Password == loginRequest.Password {
	//		util.RespOKWithData(c, user) //返回200状态吗
	//		return
	//	}
	//}

	//util.RespErr(c)
}

func Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	users, err := dao.GetUser()
	if err != nil {
		log.Println("Error adding user:", err)
		util.RespInternalErr(c)
		return
	}

	for _, user := range users {
		if user.Username == u.Username {
			util.RespAlreadyReported(c)
			return
		}
	}

	err = dao.AddUser(u)
	if err != nil {
		log.Println("Error adding user:", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func RefreshToken(c *gin.Context) {
	var refreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := c.ShouldBindJSON(&refreshTokenRequest)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshTokenRequest.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		util.RespErr(c)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			util.RespErr(c)
			return
		}

		newToken, newRefreshToken, err := generateToken(int64(userID))
		if err != nil {
			util.RespInternalErr(c)
			return
		}

		util.RespOKWithData(c, gin.H{
			"refresh_token": newRefreshToken,
			"token":         newToken,
		})

		return
	}

	util.RespErr(c)
}

func ChangePassword(c *gin.Context) {
	var PasswordRequest struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	err := c.ShouldBindJSON(&PasswordRequest)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	user, err := dao.GetUserByUsername(PasswordRequest.Username)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	if PasswordRequest.OldPassword != user.Password {
		util.RespUnauthorized(c)
		return
	}

	err = dao.ChangeUserPassword(PasswordRequest.NewPassword, PasswordRequest.Username)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GetUserInfo(c *gin.Context) {
	userID := c.Param("id")

	user, err := dao.GetUsersInfo(userID)

	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOKWithData(c, user)
}

func ChangeUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	var request dao.Users
	if err := c.ShouldBindJSON(&request); err != nil {
		util.RespParamErr(c)
		return
	}

	err = dao.ChangeUser(id, request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			util.RespNotFound(c)
			return
		}
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}
