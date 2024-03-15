# Pack Service

This allows to set up pack sizes and calculate packs for items.

## Pre requisites

- Docker
- Golang v1.21+

## Running App

1. Build and run the app container.

`make app`

2. Inspect logs using docker

`docker logs pack-svc-go -f`

3. Run tests in local docker using

`make test`

4. Run functional tests
    - Start the app using step 1
    - Trigger the dockerized functional tests
      `make functional-test`

## Building the Docker Image

1. Navigate to the root directory of the project where the Dockerfile is located.

2. Build the Docker image using the following command:
   ```docker build -t pack-svc .```

3. This command performs the following steps:
    - Uses the golang:alpine image as a base to create a lightweight container for building the Go application.
    - Sets the working directory to /pack-svc.
    - Copies go.mod and go.sum to the working directory and downloads the Go modules.
    - Copies the source code into the container.
    - Compiles the Go source code into a static binary named pack-svc.

   The build process creates a final image based on scratch, which is a minimal image with just the compiled
   application, the static frontend assets (if any), and the environment configuration.

## Running the Container

1. To run the pack-svc service, use the following command:
   ```
   docker run -p 8080:8080 \
    -e SERVICE_NAME="pack service" \
    -e LOG_LEVEL=debug \
    -e HTTP_SERVER_HOST=localhost \
    -e PORT=8080 \
    -e HTTP_SERVER_READ_TIMEOUT_IN_SEC=5 \
    -e HTTP_SERVER_WRITE_TIMEOUT_IN_SEC=5 \
    pack-svc
   ```
2. This command runs the pack-svc container and maps the container’s port 8080 to the host's port 8080, allowing the
   service to be accessed via localhost:8080.

## Environment Configuration

The service configuration is managed through a .env file, which is copied into the container. Here are the configurable
environment variables:

- SERVICE_NAME: The name of the service (e.g., "pack service").
- LOG_LEVEL: The logging level of the service (e.g., debug, info, error).
- HTTP_SERVER_HOST: The hostname for the HTTP server (e.g., localhost).
- PORT: The port on which the HTTP server listens (e.g., 8080).
- HTTP_SERVER_READ_TIMEOUT_IN_SEC: The read timeout for the HTTP server in seconds.
- HTTP_SERVER_WRITE_TIMEOUT_IN_SEC: The write timeout for the HTTP server in seconds.

## Live Website

This is the hosted link. However since it's free, it might take some to come up if it's unused.

https://pack-service.onrender.com/

Free instances spin down with inactivity

