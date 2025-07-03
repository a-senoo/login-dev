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

# Copy the Go module files to the working directory
COPY go.mod go.sum ./
# Set the environment variable to disable CGO
RUN go mod download

# Set the working directory inside the container
ENV APP_HOME /work/backend
WORKDIR $APP_HOME

# Download the Go module dependencies
COPY . .

# Set the environment variable to disable CGO
RUN go build -v -o /usr/local/bin/app .

VOLUME $APP_HOME

CMD ["app"]