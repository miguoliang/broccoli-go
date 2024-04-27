package test

import (
	"encoding/json"
	"fmt"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/miguoliang/broccoli-go/internal/persistence"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VertexTestSuite struct {
	Suite
}

func (s *VertexTestSuite) TestCreateVertexBadRequest() {

	w := s.Post("/api/vertex", nil)
	s.Equal(400, w.Code)
}

func (s *VertexTestSuite) TestCreateVertexSucceed() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
		Properties: map[string]string{
			"test": "test",
		},
	})
	s.Equal(201, w.Code)

	var response dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.NotEqual("", response.ID)
}

func (s *VertexTestSuite) TestCreateVertexConflict() {

	body := dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
	}

	w := s.Post("/api/vertex", body)
	s.Equal(201, w.Code)

	w = s.Post("/api/vertex", body)
	s.Equal(409, w.Code)
}

func (s *VertexTestSuite) TestFindVertexByIdNotFound() {

	w := s.Get("/api/vertex/0")
	s.Equal(404, w.Code)
}

func (s *VertexTestSuite) TestFindVertexByIdSucceed() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
		Properties: map[string]string{
			"test": "test",
		},
	})
	s.Equal(201, w.Code)

	var response dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)

	w = s.Get(fmt.Sprintf("/api/vertex/%d", response.ID))
	s.Equal(200, w.Code)
	var vertex persistence.Vertex
	err = json.NewDecoder(w.Body).Decode(&vertex)
	s.Nil(err)
	s.Equal(response.ID, vertex.ID)
	s.Equal(s.T().Name(), vertex.Name)
	s.Equal("test", vertex.Type)
	s.Equal(1, len(vertex.Properties))
	s.Equal("test", vertex.Properties[0].Key)
	s.Equal("test", vertex.Properties[0].Value)
}

func (s *VertexTestSuite) TestSearchVerticesEmpty() {

	w := s.Get("/api/vertex?q=not-exists")
	s.Equal(200, w.Code)

	var response dto.PageResponse[persistence.Vertex]
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.Equal(int64(0), response.Total)
}

func (s *VertexTestSuite) TestSearchVerticesNotEmpty() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
	})
	s.Equal(201, w.Code)

	w = s.Get(fmt.Sprintf("/api/vertex?q=%s", s.T().Name()))
	s.Equal(200, w.Code)
	var response dto.SearchVerticesResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.Equal(int64(1), response.Total)
	s.Equal(s.T().Name(), response.Data[0].Name)
	s.Equal("test", response.Data[0].Type)
}

func (s *VertexTestSuite) TestDeleteVertexByIdNotFound() {

	w := s.Delete("/api/vertex/0")
	s.Equal(404, w.Code)
}

func (s *VertexTestSuite) TestDeleteVertexByIdSucceed() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
	})
	s.Equal(201, w.Code)
	var response dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)

	w = s.Delete(fmt.Sprintf("/api/vertex/%d", response.ID))
	s.Equal(204, w.Code)

	w = s.Get(fmt.Sprintf("/api/vertex/%d", response.ID))
	s.Equal(404, w.Code)
}

func (s *VertexTestSuite) TestCreateVertexPropertySucceed() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name(),
		Type: "test",
	})
	s.Equal(201, w.Code)

	var response dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)

	w = s.Post(fmt.Sprintf("/api/vertex/%d/property", response.ID), dto.CreateVertexPropertyRequest{
		Key:   "test",
		Value: "test",
	})
	s.Equal(201, w.Code)

	w = s.Get(fmt.Sprintf("/api/vertex/%d", response.ID))
	s.Equal(200, w.Code)
	var vertex persistence.Vertex
	err = json.NewDecoder(w.Body).Decode(&vertex)
	s.Nil(err)
	s.Equal(response.ID, vertex.ID)
	s.Equal(s.T().Name(), vertex.Name)
	s.Equal("test", vertex.Type)
	s.Equal(1, len(vertex.Properties))
	s.Equal("test", vertex.Properties[0].Key)
	s.Equal("test", vertex.Properties[0].Value)
}

func TestVertexTestSuite(t *testing.T) {
	suite.Run(t, new(VertexTestSuite))
}
