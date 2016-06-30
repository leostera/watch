#!/bin/sh

mkdir -p ./dist

for OS in "freebsd" "linux" "darwin" "windows"; do
  for ARCH in "386" "amd64"; do
    VERSION="$(git describe --tags $1)"
    echo "Building v${VERSION} for ${OS}.${ARCH}..."
    GOOS=$OS CGO_ENABLED=0 GOARCH=$ARCH go build -ldflags "-X main.Version=$VERSION" -o watch
    ARCHIVE="watch-$VERSION-$OS-$ARCH.tar.gz"
    tar -czf ./dist/$ARCHIVE watch
    echo $ARCHIVE
  done
done
