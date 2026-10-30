package useCases

import (
	"build-doctor/entities"
	"build-doctor/ports"
)

type ArtifactBuilder struct {
	Builder   ports.Builder
	Presenter ports.Presenter
}

func (b *ArtifactBuilder) BuildArtifact() entities.BuildResult {
	buildResult := b.Builder.Build()
	b.Presenter.Present("Build failed with error:\n" + buildResult.Error.Error())
	return buildResult
}
