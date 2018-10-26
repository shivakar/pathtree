package pathtree

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestNewPathTree(t *testing.T) {
	p := New()
	require.NotNil(t, p)
	assert.Equal(t, "/", p.Root())

	p = New("/usr/local/bin")
	assert.NotNil(t, p)
	assert.Equal(t, "/", p.Root())
}

func TestAdd(t *testing.T) {
	p := New()
	require.NotNil(t, p)

	err := p.Add("/usr/local/bin/bash")
	assert.NoError(t, err)
}
