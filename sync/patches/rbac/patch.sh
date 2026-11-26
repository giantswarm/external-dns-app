#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

repo_dir=$(git rev-parse --show-toplevel) ; readonly repo_dir
script_dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd ) ; readonly script_dir

cd "${repo_dir}"

readonly script_dir_rel=".${script_dir#"${repo_dir}"}"

set -x
git apply "${script_dir_rel}/000-rbac.patch"
mkdir -p ./helm/external-dns-app/tests
cp "${script_dir_rel}/tests/"*_test.yaml ./helm/external-dns-app/tests/

{ set +x; } 2>/dev/null
