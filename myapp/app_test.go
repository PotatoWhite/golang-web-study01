package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_withoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_withName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=potato", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello potato", string(data))
}

func TestFooHandler_withoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_withJson(t *testing.T) {

	data := new(User)
	data.FirstName = "potato"
	data.LastName = "white"
	data.Email = "bravopotato@gmail.com"

	body, err := json.Marshal(data)
	if err != nil {

		return
	}

	// call
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", strings.NewReader(string(body)))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// check response code
	assert.Equal(http.StatusCreated, res.Code)

	// check create at
	resUser := new(User)
	err = json.NewDecoder(res.Body).Decode(resUser)
	if err != nil {
		return
	}
	assert.Nil(err)
	assert.Equal(data.FirstName, resUser.FirstName)
	assert.Equal(data.LastName, resUser.LastName)
	assert.Equal(data.Email, resUser.Email)
	assert.NotNil(resUser.CreatedAt)
}
