package ports

import "build-doctor/entities"

type Builder interface {
	Build() entities.BuildResult
}
