#!/usr/bin/env bash
set -ue

housecall() {
  if out_of_office; then
    page_the_doctor
  fi
  administer_the_patient || administer_the_doctor
}

out_of_office() {
  ! podman image exists "build-doctor"
}

page_the_doctor() {
  podman build -t build-doctor --build-arg ARTIFACT_REPOSITORY_DOMAIN=docker:/ .
}

administer_the_patient() {
  podman run -it --rm --runtime=crun --cgroups=disabled build-doctor:latest
}

administer_the_doctor() {
  podman run -it --rm --runtime=crun --cgroups=disabled --entrypoint /bin/bash build-doctor:latest
}

housecall

