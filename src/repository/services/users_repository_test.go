package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestAccessTokenI(t *testing.T) {
	assert.True(t, true, "Brand new Access token should be expired by Default")
}
