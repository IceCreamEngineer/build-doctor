package tests

import (
	"build-doctor/useCases"
	"testing"
)

func TestFailedBuild(t *testing.T) {
	builder := useCases.ArtifactBuilder{Builder: nil}
	builder.BuildArtifact()
}
