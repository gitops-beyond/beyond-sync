#!/bin/bash

# Exit on any error
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

# Configuration variables
INSTALL_DIR="/opt/beyond-sync"
SERVICE_USER="beyond-sync"
GITHUB_REPO="gitops-beyond/beyond-sync"

# Log helper function
log() {
    echo -e "${GREEN}[BEYOND-SYNC]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    error "Please run as root"
fi

# Create service user and group
create_user() {
    log "Creating service user and group..."
    if ! getent group $SERVICE_USER >/dev/null; then
        groupadd $SERVICE_USER
    fi
    if ! getent passwd $SERVICE_USER >/dev/null; then
        useradd -r -g $SERVICE_USER -d $INSTALL_DIR -s /sbin/nologin $SERVICE_USER
    fi
}

# Create installation directory
setup_directories() {
    log "Creating installation directory..."
    mkdir -p $INSTALL_DIR
    chown -R $SERVICE_USER:$SERVICE_USER $INSTALL_DIR
}

# Install system dependencies
install_dependencies() {
    log "Installing system dependencies..."
    
    # Check distribution
    if [ -f /etc/debian_version ]; then
        # Debian/Ubuntu
        apt-get update
        apt-get install -y redis ansible git curl
    elif [ -f /etc/redhat-release ]; then
        # RHEL/CentOS
        yum install -y epel-release
        yum install -y redis ansible git curl
    else
        error "Unsupported distribution"
    fi
    systemctl start redis
}

# Download and install binaries
install_binaries() {
    log "Installing binaries..."
    
    # Download the tar file
    log "Downloading latest release tar file..."
    curl -L "https://github.com/gitops-beyond/beyond-sync/releases/download/test/beyond-sync.tar.gz" -o "/tmp/beyond-sync.tar.gz"
    
    # Extract the tar file to installation directory
    log "Extracting binaries..."
    tar -xzvf /tmp/beyond-sync.tar.gz -C "$INSTALL_DIR"
    
    # Cleanup temporary tar
    rm -f "/tmp/beyond-sync.tar.gz"
    
    # Set permissions
    chmod +x $INSTALL_DIR/beyond-sync-api
    chmod +x $INSTALL_DIR/beyond-sync-worker
    chown $SERVICE_USER:$SERVICE_USER $INSTALL_DIR/beyond-sync-api
    chown $SERVICE_USER:$SERVICE_USER $INSTALL_DIR/beyond-sync-worker
    
    log "Binaries installed successfully"
}

# Create environment file
#setup_env() {
#    log "Setting up environment file..."
#    
#    if [ ! -f "$INSTALL_DIR/.env" ]; then
#        cat > "$INSTALL_DIR/.env" << EOF
#REDIS_HOST=localhost
#USERNAME=
#TOKEN=
#REPONAME=
## Add other required environment variables
#EOF
#        chown $SERVICE_USER:$SERVICE_USER "$INSTALL_DIR/.env"
#        chmod 600 "$INSTALL_DIR/.env"
#        log "Please configure the environment variables in $INSTALL_DIR/.env"
#    fi
#}

# Install systemd services
install_services() {
    log "Installing systemd services..."
    
    # Copy service files
    cp init/beyond-sync-api.service /etc/systemd/system/
    cp init/beyond-sync-worker.service /etc/systemd/system/
    
    # Reload systemd
    systemctl daemon-reload
    
    # Enable services
    systemctl enable beyond-sync-api
    systemctl enable beyond-sync-worker
}

# Main installation process
main() {
  log "Starting Beyond Sync installation..."
  create_user
  setup_directories
  install_dependencies
  install_binaries
  #setup_env
  install_services
  log "Installation completed successfully!"
}

main