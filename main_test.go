package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/miguoliang/broccoli-go/persistence"
	"net/http/httptest"
	"testing"
)

func setupTest() {
	gin.SetMode(gin.TestMode)
}

func TestCreateVertexHandler_ShouldBadRequestWhenNoPayload(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateVertexHandler_ShouldCreateVertex(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
		Properties: map[string]string{
			"test": "test",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", response.ID)
}

func TestCreateVertexHandler_ShouldConflictWhenVertexExists(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 409, w.Code)
}

func TestFindVertexByIdHandler_ShouldNotFoundWhenVertexNotExists(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/vertex/0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestFindVertexByIdHandler_ShouldReturnVertex(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/api/vertex/%d", response.ID)
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var vertex persistence.Vertex
	err = json.NewDecoder(w.Body).Decode(&vertex)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response.ID, vertex.ID)
	assert.Equal(t, t.Name(), vertex.Name)
	assert.Equal(t, "test", vertex.Type)
	assert.Equal(t, 1, len(vertex.Properties))
	assert.Equal(t, "test", vertex.Properties[0].Key)
	assert.Equal(t, "test", vertex.Properties[0].Value)
}

func TestSearchVerticesHandler_ShouldReturnEmptyListWhenNoVertex(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/vertex?q=empty", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response dto.PageResponse[persistence.Vertex]
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), response.Total)
}

func TestSearchVerticesHandler_ShouldReturnVertices(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/vertex?q="+t.Name(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response dto.PageResponse[persistence.Vertex]
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), response.Total)
	assert.Equal(t, t.Name(), response.Data[0].Name)
	assert.Equal(t, "test", response.Data[0].Type)
}

func TestDeleteVertexByIdHandler_ShouldNotFoundWhenVertexNotExists(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/api/vertex/0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestDeleteVertexByIdHandler_ShouldDeleteVertex(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/api/vertex/%d", response.ID)
	req = httptest.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestCreateVertexProperty_ShouldSuccess(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateVertexPropertyRequest{
		Key:   "test",
		Value: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/api/vertex/%d/property", response.ID)
	req = httptest.NewRequest("POST", url, bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var property dto.CreateVertexPropertyResponse
	err = json.NewDecoder(w.Body).Decode(&property)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", property.ID)

	w = httptest.NewRecorder()
	url = fmt.Sprintf("/api/vertex/%d", response.ID)
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var vertex persistence.Vertex
	err = json.NewDecoder(w.Body).Decode(&vertex)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(vertex.Properties))
	assert.Equal(t, "test", vertex.Properties[0].Key)
	assert.Equal(t, "test", vertex.Properties[0].Value)
}

func TestCreateEdgeHandler_ShouldBadRequestWhenNoPayload(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/edge", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateEdgeHandler_ShouldCreateEdge(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var from dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&from)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateVertexRequest{
		Name: t.Name() + "2",
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var to dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/edge", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response dto.CreateEdgeResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", response.ID)
}

func TestCreateEdgeHandler_ShouldConflictWhenEdgeExists(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var from dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&from)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateVertexRequest{
		Name: t.Name() + "2",
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var to dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/edge", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/edge", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 409, w.Code)
}

func TestSearchEdgesHandler_ShouldReturnEmptyListWhenNoEdge(t *testing.T) {

	setupTest()

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/edge?from=0&to=0&type=0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response dto.PageResponse[persistence.Edge]
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(0), response.Total)
}

func TestSearchEdgesHandler_ShouldReturnEdges(t *testing.T) {

	setupTest()

	r := setupRouter()

	jsonData, err := json.Marshal(dto.CreateVertexRequest{
		Name: t.Name(),
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var from dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&from)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateVertexRequest{
		Name: t.Name() + "2",
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/vertex", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var to dto.CreateVertexResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err = json.Marshal(dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/edge", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	url := fmt.Sprintf("/api/edge?from=%d&to=%d&type=%s", from.ID, to.ID, "test")
	req = httptest.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response dto.PageResponse[persistence.Edge]
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(1), response.Total)
	assert.Equal(t, from.ID, response.Data[0].From)
	assert.Equal(t, to.ID, response.Data[0].To)
	assert.Equal(t, "test", response.Data[0].Type)
}
