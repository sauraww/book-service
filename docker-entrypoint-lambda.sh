#!/bin/bash

export GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn

target="build"

if [[ $VERSION == "" ]]
then
  echo "No version supplied!"
  exit 1
fi

echo "Building ${SERVICE_NAME} for track '${TRACK}'. Version ${VERSION}"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR
echo $DIR

set -ex

rm -rf *.zip
zip -r -j --exclude=*.git* ${SERVICE_NAME}-${VERSION}.zip /app/main /app/logger_config.yaml /app/config.yaml
#
# upload zip to destination if the build has been triggered for lambdas
echo "destinations -> $LAMBDA_DESTINATIONS"
if [[ $LAMBDA_DESTINATIONS == "" ]]
then
  echo "Skipping push. No Lambda destinations provided."
else
  # upload to s3 destinations
  filename="${SERVICE_NAME}-${VERSION}.zip"
  destinations=$((echo $LAMBDA_DESTINATIONS | tr -d '[]') | tr -s ',' '\n')
  while read line; do
    destination=`sed -e 's/^"//' -e 's/"$//' <<<"$line"`
    echo "LAMBDA DESTINATION - ${destination}"
    aws s3 cp $filename  "s3://${destination}"
  done <<< "$destinations"
fi

set +ex