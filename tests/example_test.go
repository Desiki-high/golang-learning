package tests

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	got := math.Abs(-1)
	assert.Equal(t, got, 1.0)
	t.Logf("TestExample1")
}

func TestExample2(t *testing.T) {
	time.Sleep(100 * time.Millisecond)
	got := math.Abs(-1)
	assert.Equal(t, got, 1.0)
	t.Logf("TestExample2")
}
