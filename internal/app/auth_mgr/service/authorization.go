package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"topcoder.com/space-traffic-control/internal/app/auth_mgr/client"
	"topcoder.com/space-traffic-control/internal/app/auth_mgr/models"
	models2 "topcoder.com/space-traffic-control/internal/pkg/models"
	"topcoder.com/space-traffic-control/internal/pkg/utils"
)

type AuthorizationSvc interface {
	GenerateToken(creds models2.Credentials)(models.Token, error)
	IsAuthorized(accessToken string, role string) bool
	SignUp(user models2.User) error
}

type Authorization struct {
	dbClient client.DBService
}

func NewAuthSvc(dbClient client.DBService) Authorization {
	return Authorization{dbClient: dbClient}
}

var (
	SecretKey  = utils.GetEnvOrDefault("SECRET_KEY", "secret")
	AuthorizationHeader = http.CanonicalHeaderKey("Authorization")
)

// GenerateToken generates bearer token for a given credentials
func (authSvc *Authorization) GenerateToken(creds models2.Credentials)(models.Token, error){
	user, err := authSvc.dbClient.Login(context.Background())
	if err != nil {

	}
	expirationTime := time.Now().Add(5*time.Minute)

	claims := &models.Claims{
		Username:       creds.Username,
		UserID: user.UserID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(SecretKey)

	tokenString, err := token.SignedString(key)
	if err != nil {
		logger.Errorf("error in generating token string %v", err)
		return models.Token{}, err
	}

	tokenDetails := models.Token{
		Type:      "Bearer Token",
		Value:     tokenString,
		ExpiresAt: expirationTime,
	}
	return tokenDetails, nil
}

func (authSvc *Authorization) SignUp(user models2.User) error {
	err := authSvc.dbClient.CreateUser(user)
	if err != nil {

	}
	return err
}


func(authSvc *Authorization) IsAuthorized(accessToken string, role string) bool {
	key := []byte(SecretKey)
	claims := &models.Claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false
	}

	if claims.Role == role {
		return true
	}
	return false
}


func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(AuthorizationHeader)

		if accessToken == "" {
			logger.Errorf("couldn't find Authorization Header")
			c.JSON(http.StatusUnauthorized, "No Authorization Header passed")
			c.Abort()
			return
		}

		key := []byte(SecretKey)
		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				logger.Errorf("signature invalid")
				c.JSON(http.StatusUnauthorized, "ErrSignatureInvalid")
				c.Abort()
				return
			}
			logger.Errorf("bad token passed. err %v", err)
			c.JSON(http.StatusBadRequest, "Invalid/Expired Token Passed. Regenerate Token")
			c.Abort()
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, "Token is not valid")
			c.Abort()
			return
		}
	}
}