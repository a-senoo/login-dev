# Use the official Golang image as the base image for building the application
FROM golang:1.24-bookworm AS builder

# Set environment variables
ENV APP_HOME=/work/backend \
    CGO_ENABLED=0

# Set working directory before any file copying
WORKDIR $APP_HOME

# Copy Go module files into the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Build the Go application
RUN go build -trimpath -ldflags="-s -w" -o /usr/local/bin/app .

# Use a minimal base image for the final stage
FROM debian:bookworm-slim

# Set timezone environment variable
ENV DEBIAN_FRONTEND=noninteractive TZ=UTC

# まず sh で bash を入れる
RUN apt-get update && apt-get install -y --no-install-recommends bash ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/*

# ここから RUN は bash で
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# 残りのツールを入れるなら
RUN apt-get update && apt-get install -y --no-install-recommends wget curl gnupg procps unzip \
    && rm -rf /var/lib/apt/lists/*


# ロケール環境変数を設定
ENV TZ=UTC \
    LANG=en_US.UTF-8 \
    LANGUAGE=en_US:en \
    LC_ALL=en_US.UTF-8

# Set environment variables
ENV APP_HOME=/work/backend

# Set working directory before any file copying
WORKDIR $APP_HOME

# Copy the built Go application from the builder stage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# 非rootで実行
RUN useradd -r -u 10001 -m appuser
USER appuser

# # Install Node.js (version 18.x)
# RUN apt-get update && apt-get install -y curl gnupg && \
#     curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
#     apt-get install -y nodejs

VOLUME $APP_HOME

# 起動
CMD ["/usr/local/bin/app"]