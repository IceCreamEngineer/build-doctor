#!/usr/bin/env bash
set -ue

housecall() {
  go run -tags "containers_image_openpgp exclude_graphdriver_btrfs" .
}

housecall

