package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"go.podman.io/buildah/define"
	"go.podman.io/podman/v6/pkg/bindings"
	"go.podman.io/podman/v6/pkg/bindings/images"
	"go.podman.io/podman/v6/pkg/domain/entities/types"
)

func main() {
	if err := run(); err != nil {
		_, printError := fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
		if printError != nil {
		}
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	containerRuntime, runtimeError := connectToLocalContainerRuntime(ctx)
	if runtimeError != nil {
		return runtimeError
	}
	return buildImage(containerRuntime)
}

func connectToLocalContainerRuntime(ctx context.Context) (context.Context, error) {
	containerRuntimeSocketPath := "/tmp/podman.sock"
	containerRuntimeSocket := "unix://" + containerRuntimeSocketPath

	serviceStartError := startContainerRuntimeService(ctx, containerRuntimeSocket)
	if serviceStartError != nil {
		return ctx, serviceStartError
	}

	socketConnection, socketConnectionError := bindings.NewConnection(ctx, containerRuntimeSocket)
	if socketConnectionError != nil {
		return nil, fmt.Errorf("failed to connect to local container runtime service at %q: %w",
			containerRuntimeSocket, socketConnectionError)
	}
	return socketConnection, nil
}

func startContainerRuntimeService(ctx context.Context, containerRuntimeSocket string) error {
	containerRuntimeService := exec.CommandContext(ctx, "podman", "system", "service", "--time=0", containerRuntimeSocket)
	if err := containerRuntimeService.Start(); err != nil {
		return fmt.Errorf("failed to start embedded container runtime service: %w", err)
	}
	return nil
}

func buildImage(containerRuntime context.Context) error {
	containerfile := filepath.Join("/workspace", "Containerfile")
	imageName := "localhost/build-doctor-demo:latest"
	buildReport, buildError := build(containerfile, imageName, containerRuntime)
	if buildError != nil {
		return fmt.Errorf("build image %q: %w", imageName, buildError)
	}
	presentBuiltImage(imageName, buildReport)
	return nil
}

func build(containerfile string, imageName string, containerRuntime context.Context) (*types.BuildReport, error) {
	return images.Build(containerRuntime, []string{containerfile}, types.BuildOptions{
		BuildOptions: define.BuildOptions{
			AdditionalTags:   []string{imageName},
			ContextDirectory: "/workspace", // Explicitly bind context to avoid working directory mismatches
		},
	})
}

func presentBuiltImage(imageName string, buildReport *types.BuildReport) {
	fmt.Printf("built image: %s\n", imageName)
	if buildReport != nil {
		fmt.Printf("buildReport: %+v\n", buildReport)
	}
}
