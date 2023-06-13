# Tiktok Assignment

This is a instant messaging service implementation with Redis as required for Tiktok Immersion 2023.

## Installation

Requirement:

- golang 1.18+
- docker

To install dependency tools:

```bash
make pre
```

## Run

```bash
docker-compose up -d
```

Check if it's running:

```bash
curl localhost:8080/ping
```
