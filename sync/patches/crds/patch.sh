#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

repo_dir=$(git rev-parse --show-toplevel) ; readonly repo_dir
script_dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd ) ; readonly script_dir

cd "${repo_dir}"

readonly script_dir_rel=".${script_dir#"${repo_dir}"}"

set -x
# rm -rf ./helm/external-dns-app/crds

{ set +x; } 2>/dev/null

cp -R "${script_dir_rel}/files" ./helm/external-dns-app/
cp -R "${script_dir_rel}/crds.yaml" ./helm/external-dns-app/templates

set -x

{ set +x; } 2>/dev/null
