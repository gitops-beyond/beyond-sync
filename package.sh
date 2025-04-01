#/bin/bash

# Compile go code into binary for Linux ARM
GOOS=linux GOARCH=arm64 go build -C cmd/api/ -o ../../bin/beyond-sync-api
GOOS=linux GOARCH=arm64 go build -C cmd/worker/ -o ../../bin/beyond-sync-worker

# Package binaries into tar.gz
tar --no-xattrs -czvf bin/beyond-sync.tar.gz -C bin beyond-sync-api beyond-sync-worker