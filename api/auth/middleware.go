package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

const (
	Domain   = "AUTH0_DOMAIN"
	Audience = "AUTH0_AUDIENCE"
)

func stripBearerPrefixFromTokenString(tok string) string {
	// Should be a bearer token
	if len(tok) > 6 && strings.EqualFold(tok[0:7], "BEARER ") {
		return tok[7:]
	}
	return tok
}

func handleUnAuthorizedAccess(ginContext *gin.Context) {
	ginContext.AbortWithStatus(http.StatusUnauthorized)
}

func GetUser(ginContext *gin.Context) {
	ginContext.Header("Access-Control-Allow-Origin", "*")
	ginContext.Header("Access-Control-Allow-Headers", "*")

	issuerURL, err := url.Parse("https://" + os.Getenv(Domain) + "/")

	pwd, _ := os.Getwd()
	fmt.Println("---- pwd :", pwd)
	fmt.Println("---- issuerURL :", issuerURL)
	if err != nil {
		// TODO log error
		fmt.Println("error 1")
		handleUnAuthorizedAccess(ginContext)
		return
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv(Audience)},
	)
	// get the token from the request header
	authHeader := stripBearerPrefixFromTokenString(ginContext.Request.Header.Get("Authorization"))

	// Validate the token
	tokenInfo, err := jwtValidator.ValidateToken(ginContext, authHeader)
	if err != nil {
		// TODO log error
		fmt.Println("error 2", err)
		handleUnAuthorizedAccess(ginContext)
		return
	}

	claims, ok := tokenInfo.(*validator.ValidatedClaims)
	if !ok {
		// TODO log error
		fmt.Println("error 3")
		handleUnAuthorizedAccess(ginContext)
		return
	}

	subject := claims.RegisteredClaims.Subject
	fmt.Println("Subject claims :", subject)

	ginContext.Next()
}
