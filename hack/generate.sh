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

removeDecoderDisallowUnknownFields() {
  local folder="$1"

  find "$folder" -type f -name "*.go" | while read -r file; do
    sed -i '' 's/\(decoder\.DisallowUnknownFields()\)/\/\/ \1/' "$file"
  done
}

removeDecoderDisallowUnknownFields "channel"
removeDecoderDisallowUnknownFields "merchant"

cd "$(git rev-parse --show-toplevel)/channel"
go mod tidy

cd "$(git rev-parse --show-toplevel)/merchant"
go mod tidy
