package tests

import (
	"build-doctor/tests/testDoubles"
	"build-doctor/useCases"
	"testing"
)

var presenter = testDoubles.PresenterSpy{}

func TestFailedBuild(t *testing.T) {
	builder := useCases.ArtifactBuilder{Builder: testDoubles.FailedBuilderStub{}, Presenter: presenter}
	buildResult := builder.BuildArtifact()
	if buildResult.Error == nil {
		t.Errorf("expected build error")
	}
}
