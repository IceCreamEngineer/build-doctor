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
	assertPresented(t, "Build failed with error:\nbuild failed")
}

func assertPresented(t *testing.T, presentedMessage string) {
	if presenter.GetPresented() != presentedMessage {
		t.Errorf("Expected presented message '%s' but got '%s'", presentedMessage, presenter.GetPresented())
	}
}
