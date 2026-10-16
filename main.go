package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.podman.io/buildah/define"
	"go.podman.io/podman/v6/pkg/bindings"
	"go.podman.io/podman/v6/pkg/bindings/images"
	"go.podman.io/podman/v6/pkg/domain/entities/types"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	socket := "unix:///run/user/1000/podman/podman.sock"

	conn, err := bindings.NewConnection(ctx, socket)
	if err != nil {
		return fmt.Errorf("connect to podman service at %q: %w", socket, err)
	}

	containerfile := filepath.Join("/home/mpistacchio/test", "simpleContainerfile")
	imageName := "localhost/build-doctor-demo:latest"

	response, err := images.Build(conn, []string{containerfile}, types.BuildOptions{
		BuildOptions: define.BuildOptions{
			AdditionalTags: []string{imageName},
		},
	})
	if err != nil {
		return fmt.Errorf("build image %q: %w", imageName, err)
	}

	fmt.Printf("built image: %s\n", imageName)

	if response != nil {
		fmt.Printf("build response: %+v\n", response)
	}

	return nil
}
