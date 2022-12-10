package main

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestSelection(t *testing.T) {
	t.Parallel()

	tests.Sort(t, sort)
}
