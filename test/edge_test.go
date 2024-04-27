package test

import (
	"encoding/json"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EdgeTestSuite struct {
	Suite
}

func (s *EdgeTestSuite) TestCreateEdgeBadRequest() {

	w := s.Post("/api/edge", nil)
	s.Equal(400, w.Code)
}

func (s *EdgeTestSuite) TestCreateEdgeSucceed() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var from dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&from)
	s.Nil(err)

	w = s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var to dto.CreatedResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	s.Nil(err)

	w = s.Post("/api/edge", dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(201, w.Code)

	var response dto.CreatedResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.NotEmpty(response.ID)
}

func (s *EdgeTestSuite) TestCreateEdgeConflict() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var from dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&from)
	s.Nil(err)

	w = s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var to dto.CreatedResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	s.Nil(err)

	w = s.Post("/api/edge", dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(201, w.Code)

	w = s.Post("/api/edge", dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(409, w.Code)
}

func (s *EdgeTestSuite) TestSearchEdgesEmpty() {

	w := s.Get("/api/edge?from=0&to=0&type=test")
	s.Equal(200, w.Code)

	var response dto.SearchEdgesResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.Equal(int64(0), response.Total)
}

func (s *EdgeTestSuite) TestSearchEdgesNotEmpty() {

	w := s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var from dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&from)
	s.Nil(err)

	w = s.Post("/api/vertex", dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var to dto.CreatedResponse
	err = json.NewDecoder(w.Body).Decode(&to)
	s.Nil(err)

	w = s.Post("/api/edge", dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(201, w.Code)
}

func TestEdgeTestSuite(t *testing.T) {
	suite.Run(t, new(EdgeTestSuite))
}
