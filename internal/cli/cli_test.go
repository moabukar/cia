package cli

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	assert.NotPanics(t, func() { Execute("0.0.0", []string{"cia", "--version"}) })
}

func TestNewApp(t *testing.T) {
	app := NewApp("0.0.0", time.Now())
	assert.Equal(t, "cia", app.Name)
	assert.Equal(t, "0.0.0", app.Version)
}
