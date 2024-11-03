# syntax=docker/dockerfile:1

# Comments are provided throughout this file to help you get started.
# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/go/dockerfile-reference/

# Want to help us make this template better? Share your feedback here: https://forms.gle/ybq9Krt8jtBL3iCk7

################################################################################
# Create a stage for building the application.
ARG GO_VERSION=1.22.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Install build dependencies
RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    gcc \
    g++ \
    libc-dev \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*


RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    GOARCH=$TARGETARCH go build -o /bin/server ./cmd/server

################################################################################
# By specifying the "latest" tag, it will also use whatever happens to be the
# most recent version of that image when you build your Dockerfile. If
# reproducability is important, consider using a versioned tag
# Using ubuntu:latest as the base image
FROM ubuntu:latest AS final

# Install dependencies required for the Confluent Kafka Go client.
# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    tzdata \
    adduser \
    && rm -rf /var/lib/apt/lists/*

# Create a non-privileged user that the app will run under.
# See https://docs.docker.com/go/dockerfile-user-best-practices/
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy the executable from the "build" stage.
COPY --from=build /bin/server /bin/

# Set environment variables.
ARG POSTGRES_DATABASE_URL
ENV POSTGRES_DATABASE_URL=${POSTGRES_DATABASE_URL}
ENV SECRETKEY=OKb_JT2nE-FExnL8tw4jdNJHLFYIOdhQMJav4H_kjrY=
ENV BOOTSTRAP_SERVERS=pkc-312o0.ap-southeast-1.aws.confluent.cloud:9092
ENV SASL_USERNAME=MALHSLV25KMK2JXU
ENV SASL_PASSWORD=9DTawTMGr2L7kvdXbM+jStqcf/q4P9M4fzBoEx1Oak5rZmfz6KAMbBP/+VC13jP1
ENV SECURITY_PROTOCOL=SASL_SSL
ENV SASL_MECHANISMS=PLAIN
ENV ACKS=all

# Expose the port that the application listens on.
EXPOSE 50051

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
