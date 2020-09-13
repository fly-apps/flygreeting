#!/bin/sh -l

# Install curl
apk --no-cache add curl

# Use curl to install flyctl
curl -L https://fly.io/install.sh | sh

# Set flyctl install path
export FLYCTL_INSTALL="/root/.fly"
export PATH="$FLYCTL_INSTALL/bin:$PATH"
echo "Successfully Installed Flyctl"

# deploy app
sh -c "flyctl deploy -t $FLYCTL_AUTH_TOKEN"

# get app Information
sh -c "flyctl info -t $FLYCTL_AUTH_TOKEN"

exit 0
