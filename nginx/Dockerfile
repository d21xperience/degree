# Stage 1: Build Nginx with gRPC support
FROM ubuntu:22.04 AS builder

# Install dependencies
RUN apt update && \
    apt install -y \
    build-essential \
    git \
    curl \
    wget \
    libpcre3-dev \
    zlib1g-dev \
    libssl-dev \
    libxml2-dev \
    libxslt1-dev

# Set working directory
WORKDIR /tmp

# Define Nginx version
ARG NGINX_VERSION=1.26.0

# Download and extract Nginx
RUN wget http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz && \
    tar zxf nginx-${NGINX_VERSION}.tar.gz

# Clone the official NJS module for gRPC
RUN git clone https://github.com/nginx/njs.git 

# Go into Nginx source dir
WORKDIR /tmp/nginx-${NGINX_VERSION}

# Configure Nginx with gRPC support
RUN ./configure \
    --add-module=/tmp/njs/nginx \
    --prefix=/etc/nginx \
    --sbin-path=/usr/sbin/nginx \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --with-http_ssl_module \
    --with-http_v2_module \
    --with-http_stub_status_module \
    --with-threads \
    --with-file-aio

# Build Nginx
RUN make -j$(nproc)
RUN make install

# Stage 2: Minimal runtime image
FROM ubuntu:22.04

# Install minimal runtime dependencies
RUN apt update && \
    apt install -y \
        libssl-dev \
        libxml2 \
        zlib1g \
        libpcre3 \
        && apt clean && rm -rf /var/lib/apt/lists/*

# Copy binaries and config from builder
COPY --from=builder /usr/sbin/nginx /usr/sbin/nginx
COPY --from=builder /etc/nginx /etc/nginx
COPY --from=builder /usr/lib/x86_64-linux-gnu/libnjs.* /usr/lib/

# Create log directory
RUN mkdir -p /var/log/nginx

# Copy our custom config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy certs if needed
COPY certs/ /etc/nginx/certs/

# Expose ports
EXPOSE 443
EXPOSE 50001

CMD ["nginx", "-g", "daemon off;"]