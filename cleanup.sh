#!/usr/bin/env bash
set -ue

cleanup() {
  podman rmi build-doctor-demo:latest
}

cleanup

