package time_inspector

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReturnTrueInspector(t *testing.T) {
	inspector := New()

	suitTable := []time.Time{
		time.Date(2022, 06, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2022, 4, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2022, 8, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2022, 10, 17, 8, 00, 0, 0, time.UTC),
	}

	for _, event := range suitTable {
		valid, err := inspector.CheckInfoAboutMovingExpires(event)
		assert.Error(t, err)
		assert.False(t, valid)
	}
}

func TestReturnFalseInspector(t *testing.T) {
	inspector := New()

	suitTable := []time.Time{
		time.Date(2023, 06, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2023, 4, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2023, 8, 17, 8, 00, 0, 0, time.UTC),
		time.Date(2023, 10, 17, 8, 00, 0, 0, time.UTC),
	}

	for _, event := range suitTable {
		valid, err := inspector.CheckInfoAboutMovingExpires(event)
		assert.NoError(t, err)
		assert.True(t, valid)
	}
}
