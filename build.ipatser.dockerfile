# Start from golang base image
FROM golang:alpine as build-stage

# Install git.
# alpine image does not have git in it
# - issue: stdlib.h: No such file or directory
#   soln : musl-dev
RUN apk update && apk add --no-cache git wget curl make gcc musl-dev
#jq

# Setup current working directory
# Copy the source code
WORKDIR /app
RUN mkdir -p /app/src/github.com/ipatser
WORKDIR /app/src/github.com/ipatser
COPY . .

# Enable go modules
# Set go env variables
ENV GO111MODULE=auto
ENV GOPATH=/app
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV GOMOD=$GOPATH/src/github.com/ipatser/go.mod

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files if not done
# Download all dependencies.
RUN go mod download

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Setup current working directory as vcs_ipatser
WORKDIR /app/src/github.com/ipatser/vcs_ipatser

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install
