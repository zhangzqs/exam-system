package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJwt_ParseToken(t *testing.T) {
	j := NewJwt("Hello", time.Hour)
	assert.NotNil(t, j)
	token := j.GenerateToken(123)
	assert.NotEmpty(t, token)
	cc, err := j.ParseToken(token)
	assert.Equal(t, cc.Uid, 123)
	assert.Nil(t, err)

	j1 := NewJwt("Hell", time.Hour)
	cc, err = j1.ParseToken(token)
	assert.Equal(t, cc.Uid, 123)
	assert.NotNil(t, err)
}
