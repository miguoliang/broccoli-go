package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func autoMigrate() {
	err := db.AutoMigrate(&Vertex{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&VertexProperty{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Edge{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&EdgeProperty{})
	if err != nil {
		return
	}
}

func TestCreateVertexHandler_ShouldBadRequestWhenNoPayload(t *testing.T) {

	r := setUpRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/vertex", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateVertexHandler_ShouldCreateVertex(t *testing.T) {

	db = connectDatabase()
	autoMigrate()
	r := setUpRouter()

	jsonData, err := json.Marshal(CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", response.ID)
}

func TestCreateVertexHandler_ShouldConflictWhenVertexExists(t *testing.T) {

	db = connectDatabase()
	autoMigrate()
	r := setUpRouter()

	jsonData, err := json.Marshal(CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 409, w.Code)
}
