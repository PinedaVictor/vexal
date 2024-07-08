#!/bin/sh

cd ..

# Source environment variables from .env.development
source .env.production

# Build application with injected environment variables using ldflags
echo "Building Production"
go build -ldflags "\
  -X 'vx/config.ENVIRONMENT=$ENVIRONMENT' \
  -X 'vx/config.APP_URL=$APP_URL' \
  -X 'vx/config.SERVER_REDIRECT_ADDR=$SERVER_REDIRECT_ADDR' \
  -X 'vx/config.GCP_PROJECT_ID=$GCP_PROJECT_ID' \
  -X 'vx/config.GCP_ACCOUNT_TYPE=$GCP_ACCOUNT_TYPE' \
  -X 'vx/config.FB_ADMIN_PRIVATE_KEY=$FB_ADMIN_PRIVATE_KEY' \
  -X 'vx/config.FB_ADMIN_CLIENT_EMAIL=$FB_ADMIN_CLIENT_EMAIL' \
  -X 'vx/config.SECRETS_MAN_ACCOUNT_TYPE=$SECRETS_MAN_ACCOUNT_TYPE' \
  -X 'vx/config.SECRETS_MAN_PRIVATE_KEY=$SECRETS_MAN_PRIVATE_KEY' \
  -X 'vx/config.SECRETS_MAN_EMAIL=$SECRETS_MAN_EMAIL'" -o vx
