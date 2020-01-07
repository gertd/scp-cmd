#!/usr/bin/env bash

set -e

export SCP_CLIENT_ID=$(cat ./appreg.json | jq -r .clientId)
export SCP_CLIENT_SECRET=$(cat ./appreg.json | jq -r .clientSecret)
