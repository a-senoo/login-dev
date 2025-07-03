FROM golang:1.24-bookworm

# Use the official Golang image as the base image
SHELL ["/bin/bash", "-euxo", "pipefail", "-c"]

# Set the timezone and locale environment variables
ENV TZ=UTC \
    LANG=en_US.UTF-8 \
    LANGUAGE=en_US:en \
    LC_ALL=en_US.UTF-8

RUN apt-get update && apt-get upgrade -y && apt-get clean


# タイムゾーン & ツール類
RUN apt-get update && apt-get install -y --no-install-recommends \
    tzdata \
    bash \
    wget \
    curl \
    gnupg \
    procps \
    inotify-tools \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# ランタイム系のライブラリ
RUN apt-get update && apt-get install -y --no-install-recommends \
    libuv1 \
    libicu-dev \
    libzip-dev \
    libonig-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Node.js (version 18.x)
RUN apt-get update && apt-get install -y curl gnupg && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# Set environment variables
ENV APP_HOME=/work/backend \
    CGO_ENABLED=0

# Set working directory before any file copying
WORKDIR $APP_HOME

# Copy Go module files into the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Download the Go module dependencies
COPY . .

# Set the environment variable to disable CGO
RUN go build -v -o /usr/local/bin/app .

VOLUME $APP_HOME

CMD ["app"]