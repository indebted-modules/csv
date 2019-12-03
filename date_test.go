package csv_test

import (
	"testing"
	"time"

	"github.com/indebted-modules/csv"
	"github.com/jszwec/csvutil"
	"github.com/stretchr/testify/suite"
)

type DateSuite struct {
	suite.Suite
}

type DateEntity struct {
	ID      string
	Created csv.Date
}

func TestDateSuite(t *testing.T) {
	suite.Run(t, new(DateSuite))
}

func (s *DateSuite) TestMarshalCSV() {
	t := time.Date(2019, 3, 15, 0, 0, 0, 0, time.UTC)
	samples := []*DateEntity{
		&DateEntity{
			ID: "f00",
			Created: csv.Date{
				Time: t,
			},
		},
	}
	data, err := csvutil.Marshal(samples)
	s.Equal("ID,Created\nf00,2019-03-15\n", string(data))
	s.Nil(err)
}

func (s *DateSuite) TestUnmarshalCSV() {
	data := []byte("ID,Created\nf00,2019-03-15")
	samples := []*DateEntity{}
	err := csvutil.Unmarshal(data, &samples)
	s.Equal("f00", samples[0].ID)
	s.Equal(time.Date(2019, 3, 15, 0, 0, 0, 0, time.UTC), samples[0].Created.UTC())
	s.Nil(err)
}

func (s *DateSuite) TestUnmarshalEmptyValue() {
	data := []byte("ID,Created\nf00,")
	samples := []*DateEntity{}
	err := csvutil.Unmarshal(data, &samples)
	s.Equal("f00", samples[0].ID)
	s.Equal(csv.Date{}, samples[0].Created)
	s.Nil(err)
}
