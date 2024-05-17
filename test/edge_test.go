package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/stretchr/testify/suite"
)

const ENDPOINT_EDGE = "/api/edge"

type EdgeTestSuite struct {
	Suite
}

func (s *EdgeTestSuite) TestCreateEdgeBadRequest() {

	w := s.Post(ENDPOINT_EDGE, nil)
	s.Equal(400, w.Code)
}

func (s *EdgeTestSuite) TestCreateEdgeSucceed() {

	w := s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	from, err := decodeBody(w)
	s.Nil(err)

	w = s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	to, err := decodeBody(w)
	s.Nil(err)

	w = s.Post(ENDPOINT_EDGE, dto.CreateEdgeRequest{
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

func decodeBody(w *httptest.ResponseRecorder) (dto.CreatedResponse, error) {
	var to dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&to)
	return to, err
}

func (s *EdgeTestSuite) TestCreateEdgeConflict() {

	w := s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	var from dto.CreatedResponse
	err := json.NewDecoder(w.Body).Decode(&from)
	s.Nil(err)

	w = s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	to, err := decodeBody(w)
	s.Nil(err)

	w = s.Post(ENDPOINT_EDGE, dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(201, w.Code)

	w = s.Post(ENDPOINT_EDGE, dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(409, w.Code)
}

func (s *EdgeTestSuite) TestSearchEdgesEmpty() {

	w := s.Get(ENDPOINT_EDGE + "?from=0&to=0&type=test")
	s.Equal(200, w.Code)

	var response dto.SearchEdgesResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	s.Nil(err)
	s.Equal(int64(0), response.Total)
}

func (s *EdgeTestSuite) TestSearchEdgesNotEmpty() {

	w := s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_1",
		Type: "test",
	})
	s.Equal(201, w.Code)

	from, err := decodeBody(w)
	s.Nil(err)

	w = s.Post(ENDPOINT_VERTEX, dto.CreateVertexRequest{
		Name: s.T().Name() + "_2",
		Type: "test",
	})
	s.Equal(201, w.Code)

	to, err := decodeBody(w)
	s.Nil(err)

	w = s.Post(ENDPOINT_EDGE, dto.CreateEdgeRequest{
		From: from.ID,
		To:   to.ID,
		Type: "test",
	})
	s.Equal(201, w.Code)
}

func TestEdgeTestSuite(t *testing.T) {
	suite.Run(t, new(EdgeTestSuite))
}
