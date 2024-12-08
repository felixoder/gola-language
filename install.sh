#!/bin/bash

# Check if the user is root (for system-wide installation)
if [ "$(id -u)" -ne 0 ]; then
    echo "You must run this script as root (use sudo)."
    exit 1
fi

# Determine OS and architecture
OS=$(uname -s)
ARCH=$(uname -m)

# Download the appropriate binary based on OS and architecture
if [ "$OS" == "Linux" ] && [ "$ARCH" == "x86_64" ]; then
    URL="https://github.com/felixoder/gola-language/releases/download/v1.0.01/gola-linux-x86_64.tar.gz"
elif [ "$OS" == "Darwin" ] && [ "$ARCH" == "x86_64" ]; then
    URL="https://github.com/felixoder/gola-language/releases/download/v1.0.01/gola-darwin-x86_64.tar.gz"
elif [ "$OS" == "Darwin" ] && [ "$ARCH" == "arm64" ]; then
    URL="https://github.com/felixoder/gola-language/releases/download/v1.0.1/gola-darwin-arm64.tar.gz"

else
    echo "Unsupported OS or architecture"
    exit 1
fi

# Download the appropriate file
echo "Downloading Gola for $OS $ARCH..."
curl -L -o gola.tar.gz $URL

# Extract the downloaded file
echo "Extracting Gola..."
tar -xvzf gola.tar.gz

# Move the binary to a directory in the PATH
echo "Installing Gola..."
chmod +x gola
mv gola /usr/local/bin/

# Cleanup
rm -f gola.tar.gz

# Confirm installation
echo "Gola installed successfully! You can run 'gola' from anywhere."
