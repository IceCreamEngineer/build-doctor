#!/usr/bin/env bash
set -u

cleanup() {
  podman rmi build-doctor-demo:latest
  podman rmi build-doctor:latest
}

cleanup

