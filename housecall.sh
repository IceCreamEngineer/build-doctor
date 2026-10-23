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
  echo "Paging the doctor... (building build-doctor)"
  podman build -t build-doctor --build-arg ARTIFACT_REPOSITORY_DOMAIN=docker:/ .
}

administer_the_patient() {
  echo "Administering the patient..."
  administer
}

administer() {
  podman run -it --rm --runtime=crun --cgroups=disabled \
    --device /dev/fuse \
    --cap-add=sys_admin,mknod \
    --security-opt label=disable \
    --security-opt unmask=/proc/all \
    ${TROUBLESHOOT:+--entrypoint /bin/bash} \
    build-doctor:latest
}

administer_the_doctor() {
  echo "Looks like the doctor is the sick one - take a look under the hood..."
  TROUBLESHOOT=1 administer
}

housecall

