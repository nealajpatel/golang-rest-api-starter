//package for common functionality
package httputil

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Executes an http request based on the provided parameters
func Execute(method string, url string, bearer string) (resp *http.Response, err error) {

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	resp, err = client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	return
}

// Error constructor string
func NewErrorString(ctx *gin.Context, status int, err string) {
	er := HTTPError{
		Code:    status,
		Message: err,
	}
	ctx.JSON(status, er)
}

// Error constructor
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status internal server error"`
}
