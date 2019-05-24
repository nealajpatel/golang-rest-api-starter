// Boilerplate test file leveraging testify.
// Creates a suite with initial setup for the server, and then executes defined API tests.
package tests

import (
	"testing"

	"encoding/json"
	"golang-rest-api-starter/config"
	"golang-rest-api-starter/server"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

// suite struct defines necessary components to execute tests
type HealthSuite struct {
	suite.Suite
	Config   *viper.Viper
	Router   *gin.Engine
	Response *httptest.ResponseRecorder
}

// sets up suite with dev configs and starts sets up the router
// SetupSuite runs once before executing any tests
func (suite *HealthSuite) SetupSuite() {
	config.Init("dev")
	suite.Config = config.GetConfig()
	suite.Router = server.SetupRouter()
}

// sets up env for each test
// SetupTest runes before each test (after SetupSuite)
func (suite *HealthSuite) SetupTest() {
	suite.Response = httptest.NewRecorder()
}

// test to verify /health functionality
func (suite *HealthSuite) TestHealthRoute() {
	req, _ := http.NewRequest("GET", "/health/health", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal("Working!", suite.Response.Body.String())
}

// test to verify /ping functionality
func (suite *HealthSuite) TestPingRoute() {
	req, _ := http.NewRequest("GET", "/health/ping", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	// generate expected response from request
	responseMap := map[string]string{"message": "pong"}
	responseJson, _ := json.Marshal(responseMap)

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(string(responseJson), suite.Response.Body.String())
}

// test to verify swagger functionality
func (suite *HealthSuite) TestSwaggerRoute() {
	req, _ := http.NewRequest("GET", "/", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	suite.Equal(http.StatusMovedPermanently, suite.Response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestHealthSuite(t *testing.T) {
	suite.Run(t, new(HealthSuite))
}
