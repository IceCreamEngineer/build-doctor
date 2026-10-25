package testDoubles

import (
	"build-doctor/entities"
	"errors"
)

type FailedBuilderStub struct{}

func (f FailedBuilderStub) Build() entities.BuildResult {
	return entities.BuildResult{Error: errors.New("build failed")}
}
