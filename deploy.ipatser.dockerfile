# Start from golang base image
FROM ipatser/build as build-stage

# Finally our multi-stage to build a small image
# Start a new stage from scratch
#FROM scratch
FROM ipatser/build as deployment-stage

# Setup current working directory
WORKDIR /

# Copy the Pre-built binary file
COPY --from=build-stage /app/bin/vcs_ipatser .
