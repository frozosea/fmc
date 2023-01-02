package captcha_resolver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomStringGenerator(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	gen := NewRandomStringGenerator()
	randStr, err := gen.Generate()
	assert.NoError(t, err)
	assert.Equal(t, len(randStr), 17)
}
