#!/usr/bin/env bash
set -ue

cleanup() {
  podman rmi build-doctor-demo:latest || podman rmi build-doctor:latest
}

cleanup

