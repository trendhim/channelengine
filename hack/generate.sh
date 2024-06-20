#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

cd "$(git rev-parse --show-toplevel)"

rm -rf ./{channel,merchant}/

openapi-generator generate \
  --input-spec https://demo.channelengine.net/api/swagger/channel/swagger.json \
  --skip-validate-spec \
  --generator-name go \
  --output ./channel \
  --git-user-id trendhim \
  --git-repo-id channelengine/channel \
  --additional-properties packageName=channel,packageVersion=1.0.0,enumClassPrefix=true,structPrefix=true

openapi-generator generate \
  --input-spec https://demo.channelengine.net/api/swagger/merchant/swagger.json \
  --skip-validate-spec \
  --generator-name go \
  --output ./merchant \
  --git-user-id trendhim \
  --git-repo-id channelengine/merchant \
  --additional-properties packageName=merchant,packageVersion=1.0.0,enumClassPrefix=true,structPrefix=true

rm -rf ./{channel,merchant}/{.travis.yml,git_push.sh}

cd "$(git rev-parse --show-toplevel)/channel"
go mod tidy

cd "$(git rev-parse --show-toplevel)/merchant"
go mod tidy
