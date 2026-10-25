package tests

import (
	"build-doctor/useCases"
	"testing"
)

func TestBuild(t *testing.T) {
	builder := useCases.Builder{}
	builder.Build()
}
