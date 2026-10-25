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
	return b.Builder.Build()
}
