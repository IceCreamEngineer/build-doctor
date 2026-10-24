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
		fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Use an isolated local socket inside the container filesystem
	socketPath := "/tmp/podman.sock"
	socket := "unix://" + socketPath

	// 1. Start the embedded local Podman API Service
	cmd := exec.CommandContext(ctx, "podman", "system", "service", "--time=0", socket)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start embedded podman service: %w", err)
	}
	defer func() {
		if cmd.Process != nil {
			cmd.Process.Kill() // Ensure the service is cleaned up on exit
		}
	}()

	// Wait briefly for the local socket file to be generated
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(socketPath); err == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 2. Connect to the isolated local service
	conn, err := bindings.NewConnection(ctx, socket)
	if err != nil {
		return fmt.Errorf("connect to local podman service at %q: %w", socket, err)
	}

	containerfile := filepath.Join("/workspace", "Containerfile")
	imageName := "localhost/build-doctor-demo:latest"

	response, err := images.Build(conn, []string{containerfile}, types.BuildOptions{
		BuildOptions: define.BuildOptions{
			AdditionalTags:   []string{imageName},
			ContextDirectory: "/workspace", // Explicitly bind context to avoid working directory mismatches
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
