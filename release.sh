#!/bin/bash -e

readonly DIST=./dist
readonly NAME="$(basename $(pwd))"
readonly VERSION="$(git describe --tags $1)"

cleanup () {
  rm -rf ${DIST}
}
trap cleanup ERR

mkdir -p ${DIST}

echo "Building v${VERSION}:"
for OS in "freebsd" "linux" "darwin" "windows"; do
  for ARCH in "386" "amd64"; do
    echo -n "-> ${OS}.${ARCH}..."
    CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags "-X main.Version=${VERSION}" -o ${NAME}
    ARCHIVE="${NAME}-${VERSION}-${OS}-${ARCH}.tar.gz"
    tar -czf ${DIST}/${ARCHIVE} ${NAME}
    echo ✔︎
  done
done
