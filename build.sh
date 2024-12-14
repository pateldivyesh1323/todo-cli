#!/bin/bash

# Application name
APP_NAME="todo"

# Create output directory
OUTPUT_DIR="bin"
mkdir -p $OUTPUT_DIR

# Array of target platforms
platforms=("windows/amd64" "linux/amd64" "darwin/amd64" "darwin/arm64")

# Build for each platform
for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$OUTPUT_DIR/$APP_NAME-$GOOS-$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building $APP_NAME for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
    if [ $? -ne 0 ]; then
        echo "Failed to build for $GOOS/$GOARCH"
        exit 1
    fi
done

echo "Build complete! Executables are in the $OUTPUT_DIR directory."
