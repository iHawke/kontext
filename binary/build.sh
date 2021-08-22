#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

platforms=("windows/amd64" "windows/386"
           "darwin/amd64" "darwin/arm64"
           "linux/386" "linux/amd64" "linux/arm64"
           "openbsd/386" "openbsd/amd64")

cd ../cmd

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    OUT=$package'-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        OUT+='.exe'
    fi
    echo "Building [$OUT]"
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o ../binary/$OUT "$package".go
    if [ $? -ne 0 ]; then
        echo "An error has occurred! skipping $OUT"
    fi
    sleep 1
done