#!/usr/bin/env bash
# https://sharats.me/posts/shell-script-best-practices/

# https://cdkworkshop.com/15-prerequisites.html
# https://dashboard.eventengine.run/dashboard
# https://cdkworkshop.com/60-go.html

set -o errexit
set -o nounset
set -o pipefail
if [[ "${TRACE-0}" == "1" ]]; then
    set -o xtrace
fi

#brew install go
brew install awscli aws-cdk
#export AWS_DEFAULT_REGION=eu-west-1
#export AWS_ACCESS_KEY_ID=ASIA
#export AWS_SECRET_ACCESS_KEY=Vu
#export AWS_SESSION_TOKEN=IQ
# git config user.email david.aparicio@free.fr

#mkdir cdk-workshop && cd cdk-workshop
#cdk init sample-app --language go
#go mod tidy

cdk synth
cdk bootstrap
cdk deploy
cdk diff
cdk deploy