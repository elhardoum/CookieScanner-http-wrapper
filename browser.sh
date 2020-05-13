#!/usr/bin/env sh

# $1 = running port
chromium-browser \
    --headless \
    --disable-gpu \
    --disable-software-rasterizer \
    --disable-dev-shm-usage \
    --no-sandbox \
    --remote-debugging-address=0.0.0.0 \
    --remote-debugging-port=$1
