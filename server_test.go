package main

import (
	"bytes"
	"fmt"
	"github.com/onlinehead/simple-rest/pkg/birthday"
	"github.com/onlinehead/simple-rest/pkg/routes"
	"github.com/onlinehead/simple-rest/pkg/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	tests.SkipHTTPTest(t)
	router := SetupGin()
	username, password, db, host, port := tests.GetPostgresTestParams()
	repo, err := initPostgresDB(fmt.Sprintf("%v:%v", host, port), username, password, db, "postgres_migrations")
	if err != nil {
		t.Error("Cannot initialize connnection to a test database", err)
	}
	routes.Repo = repo
	w := httptest.NewRecorder()

	payload := []byte(`{"dateOfBirth": "2016-02-29"}`)
	req, _ := http.NewRequest("PUT", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/hello/testuserzz",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2017-02-29"}`)
	req, _ = http.NewRequest("PUT", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2030-02-29"}`)
	req, _ = http.NewRequest("PUT", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2010-02-23"}`)
	req, _ = http.NewRequest("PUT", "/hello/testuser123",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "zzzzzz"}`)
	req, _ = http.NewRequest("PUT", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/hello/testuser",  nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2016-02-29"}`)
	req, _ = http.NewRequest("PUT", "/hello/testuser/asd",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "zzzzzz"}`)
	req, _ = http.NewRequest("PUT", "/hello/",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func TestUserBirthday(t *testing.T) {
	tests.SkipHTTPTest(t)
	router := SetupGin()
	username, password, db, host, port := tests.GetPostgresTestParams()
	repo, err := initPostgresDB(fmt.Sprintf("%v:%v", host, port), username, password, db, "postgres_migrations")
	if err != nil {
		t.Error("Cannot initialize connnection to a test database", err)
	}
	routes.Repo = repo
	w := httptest.NewRecorder()

	payload := []byte(`{"dateOfBirth": "2016-02-29"}`)
	req, _ := http.NewRequest("PUT", "/hello/testuserx",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2015-01-29"}`)
	req, _ = http.NewRequest("PUT", "/hello/testusers",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	bdatenow := fmt.Sprintf(`{"dateOfBirth": "%v-%02d-%02d"}`, time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	println(bdatenow)
	payload = []byte(bdatenow)
	req, _ = http.NewRequest("PUT", "/hello/testusernow",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	birthdayObj,_ := birthday.GetTimeObjFromVals(2016, 02, 29)
	daysBeforeBirthday := birthday.GetDaysBeforeBirthday(birthdayObj, time.Now())
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/testuserx",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"message\":\"Hello, testuserx! Your birthday in %v day(s)\"}",  daysBeforeBirthday), w.Body.String())

	birthdayObj,_ = birthday.GetTimeObjFromVals(2015, 01, 29)
	daysBeforeBirthday = birthday.GetDaysBeforeBirthday(birthdayObj, time.Now())
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/testusers",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"message\":\"Hello, testusers! Your birthday in %v day(s)\"}",  daysBeforeBirthday), w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/testusernow",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Hello, testusernow! Happy birthday!\"}", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/hello/testuser",  nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	payload = []byte(`{"dateOfBirth": "2016-02-29"}`)
	req, _ = http.NewRequest("GET", "/hello/testuser",  bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/",  nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ping",  nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}