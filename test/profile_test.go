package test

import (
	"encoding/json"
	"fmt"
	"github.com/miguoliang/broccoli-go/internal/persistence"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ProfileTestSuite struct {
	Suite
}

func (s *ProfileTestSuite) TestUsageSucceed() {

	usage := &persistence.Usage{
		Date:              time.Now().AddDate(0, 0, -7),
		VertexNum:         10,
		EdgeNum:           15,
		VertexPropertyNum: 100,
	}
	db := persistence.GetDatabaseConnection()
	result := db.Create(&usage)
	s.Require().NoError(result.Error)
	s.Require().Equal(uint(10), usage.VertexNum)
	s.Require().Equal(uint(15), usage.EdgeNum)
	s.Require().Equal(uint(100), usage.VertexPropertyNum)

	w := s.Get(fmt.Sprintf("/api/profile/usage?start_date=%s&end_date=%s",
		time.Now().AddDate(0, 0, -15).Format("2006-01-02"),
		time.Now().Format("2006-01-02")))
	s.Require().Equal(200, w.Code)
	var usages []persistence.Usage
	err := json.NewDecoder(w.Body).Decode(&usages)
	s.Require().NoError(err)
	s.Require().Len(usages, 1)
	s.Require().Equal(uint(10), usages[0].VertexNum)
	s.Require().Equal(uint(15), usages[0].EdgeNum)
	s.Require().Equal(uint(100), usages[0].VertexPropertyNum)
}

func (s *ProfileTestSuite) TestUsageBadRequestWhenEndDateBeforeStartDate() {

	w := s.Get("/api/profile/usage?start_date=2021-01-01&end_date=2020-01-01")
	s.Require().Equal(400, w.Code)
}

func TestProfileTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileTestSuite))
}
