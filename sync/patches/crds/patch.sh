#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

repo_dir=$(git rev-parse --show-toplevel) ; readonly repo_dir
script_dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd ) ; readonly script_dir

cd "${repo_dir}"

readonly script_dir_rel=".${script_dir#"${repo_dir}"}"

set -x
rm -rf ./crds
mkdir -p ./crds
cp ./helm/external-dns-app/crds/*.yaml ./crds/
rm -rf ./helm/external-dns-app/crds

mkdir -p ./helm/external-dns-app/templates/crds
cp "${script_dir_rel}/crds-helpers.tpl" ./helm/external-dns-app/templates/crds/_helpers.tpl
cp "${script_dir_rel}/crds-serviceaccount.yaml" ./helm/external-dns-app/templates/crds/serviceaccount.yaml
cp "${script_dir_rel}/crds-rbac.yaml" ./helm/external-dns-app/templates/crds/rbac.yaml
cp "${script_dir_rel}/crds-networkpolicy.yaml" ./helm/external-dns-app/templates/crds/networkpolicy.yaml
cp "${script_dir_rel}/crds-job.yaml" ./helm/external-dns-app/templates/crds/job.yaml

{ set +x; } 2>/dev/null
