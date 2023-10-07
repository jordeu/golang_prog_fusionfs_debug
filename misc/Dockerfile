FROM debian:12.1-slim
RUN apt-get update && apt-get install -y procps && rm -rf /var/lib/apt/lists/*

COPY fusionfs_debug /usr/local/bin/
