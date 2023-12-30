package time

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var JST *time.Location

func TestMain(m *testing.M) {
	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Test_RealClock(t *testing.T) {
	testCases := []struct {
		loc *time.Location
	}{
		{time.UTC},
		{JST},
	}
	for _, tc := range testCases {
		t.Run(tc.loc.String(), func(t *testing.T) {
			c := NewRealClock(*tc.loc)

			got := c.Now().Location()
			assert.Equal(t, tc.loc, got)
		})
	}
}
