package useCases

import (
	"build-doctor/ports"
)

type ArtifactBuilder struct{ Builder ports.Builder }

func (b *ArtifactBuilder) BuildArtifact() {}
