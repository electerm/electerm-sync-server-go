package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/electerm/electerm-sync-server-go/src/store"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	router  *gin.Engine
	testDir string
}

func (suite *MainTestSuite) SetupSuite() {
	// Set Gin to release mode for tests
	gin.SetMode(gin.ReleaseMode)

	// Create test data directory
	suite.testDir = "test-data.db"

	// Set environment variables
	os.Setenv("PORT", "7837")
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_USERS", "testuser,anotheruser")
	os.Setenv("DB_PATH", suite.testDir)

	// Initialize SQLite store
	if err := store.SQLiteStore.Init(); err != nil {
		panic("Failed to initialize SQLite store: " + err.Error())
	}

	suite.router = setupRouter()
}

func (suite *MainTestSuite) TearDownSuite() {
	// Close SQLite store
	store.SQLiteStore.Close()
	// Cleanup
	os.Remove(suite.testDir)
}

func (suite *MainTestSuite) generateTestToken(userId string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString
}

func (suite *MainTestSuite) TestTestEndpoint() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("ok", w.Body.String())
}

func (suite *MainTestSuite) TestUnauthorizedAccess() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sync", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(401, w.Code)
}

func (suite *MainTestSuite) TestInvalidToken() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	suite.router.ServeHTTP(w, req)

	suite.Equal(401, w.Code)
}

func (suite *MainTestSuite) TestUnauthorizedUser() {
	token := suite.generateTestToken("unauthorized-user")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(401, w.Code)
}

func (suite *MainTestSuite) TestSyncWorkflow() {
	token := suite.generateTestToken("testuser")

	// Test PUT with valid data
	testData := map[string]interface{}{
		"test": "data",
		"nested": map[string]interface{}{
			"key": "value",
		},
	}
	jsonData, _ := json.Marshal(testData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/sync", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("ok", w.Body.String())

	// Test GET
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	suite.NoError(err)
	suite.Equal("data", responseData["test"])
	suite.Equal("value", responseData["nested"].(map[string]interface{})["key"])
}

func (suite *MainTestSuite) TestSyncPost() {
	token := suite.generateTestToken("testuser")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("test ok", w.Body.String())
}

func (suite *MainTestSuite) TestSyncGetNotFound() {
	token := suite.generateTestToken("testuser")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(404, w.Code)
	suite.Equal("File not found", w.Body.String())
}

func (suite *MainTestSuite) TestSyncPutInvalidJSON() {
	token := suite.generateTestToken("testuser")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/sync", bytes.NewBufferString("invalid json"))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(400, w.Code)
}

func (suite *MainTestSuite) TestSyncMethodNotAllowed() {
	token := suite.generateTestToken("testuser")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/sync", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(405, w.Code)
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
