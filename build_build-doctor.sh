#!/usr/bin/env bash
set -ue

podman build -t build-doctor --build-arg ARTIFACT_REPOSITORY_DOMAIN=docker:/ .

