package test

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/resource"
	"github.com/miguoliang/broccoli-go/pkg/str"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
)

type Suite struct {
	suite.Suite
	r *gin.Engine
}

func (s *Suite) SetupSuite() {
	log.Println("Setup suite")
	s.r = resource.SetupRouter()
	gin.SetMode(gin.TestMode)
}

func (s *Suite) Head(url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("HEAD", url, nil)
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Get(url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Post(url string, body interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url, str.StructToJsonReader(body))
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Put(url string, body interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", url, str.StructToJsonReader(body))
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Delete(url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", url, nil)
	s.r.ServeHTTP(w, req)
	return w
}
