package broccoli_go

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	assert.Equal(t, 201, w.Code)

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

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 409, w.Code)
}

func TestFindVertexByIdHandler_ShouldNotFoundWhenVertexNotExists(t *testing.T) {

	db = connectDatabase()
	autoMigrate()
	r := setUpRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/vertex/0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestFindVertexByIdHandler_ShouldReturnVertex(t *testing.T) {

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

	assert.Equal(t, 201, w.Code)

	var response CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/vertex/%d", response.ID)
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSearchVerticesHandler_ShouldReturnEmptyListWhenNoVertex(t *testing.T) {

	db = connectDatabase()
	autoMigrate()
	r := setUpRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/vertex?q=empty", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response PageResponse[Vertex]
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), response.Total)
}

func TestSearchVerticesHandler_ShouldReturnVertices(t *testing.T) {

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

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/vertex?q="+t.Name(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response PageResponse[Vertex]
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), response.Total)
}

func TestDeleteVertexByIdHandler_ShouldNotFoundWhenVertexNotExists(t *testing.T) {

	db = connectDatabase()
	autoMigrate()
	r := setUpRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/vertex/0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestDeleteVertexByIdHandler_ShouldDeleteVertex(t *testing.T) {

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

	assert.Equal(t, 201, w.Code)

	var response CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/vertex/%d", response.ID)
	req = httptest.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
