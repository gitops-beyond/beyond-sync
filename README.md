# Beyond Sync - GitOps for Ansible Configuration Management

Beyond Sync is a groundbreaking GitOps tool designed to bring GitOps methodology beyond Kubernetes and into the world of Ansible configuration management. It's possibly the first GitOps tool specifically designed for configuration management outside of Kubernetes environments.

## What is Beyond Sync?

Beyond Sync enables continuous deployment and automation of Ansible configurations using GitOps principles. It monitors your Git repository and automatically applies changes to your infrastructure when modifications are detected in your Ansible playbooks and configurations.

Key Features:
- Automated sync of Ansible configurations
- GitOps workflow for infrastructure management
- Real-time status monitoring
- REST API for integration
- Redis-backed state management

## Installation
### Prerequisites
Installation server:
- Linux **ARM** server (Debian or CentOS/RHEL)
- Ports 6379, 8080 opened
- Can access Ansible target servers from inventory
- Contains SSH keys required for Ansible project

Ansible tracked project:
- Hosted in Github
- Github PAT
- Ansible project in `ansible/` dir of the repo
- Inventory in `ansible/inventory`
- Playbook in `ansible/playbook.yml`

### Installation Steps

1. Clone the repo:
```
https://github.com/gitops-beyond/beyond-sync.git
```
2. Run the installation script as root:
```
sudo ./install.sh
```
Input the information about the GitHub repository with Ansible project you want to track\
3. Check that installed services are running:
```
sudo systemctl status beyond-sync-api
bsudo systemctl status beyond-sync-worker
```

# Release
### Release Steps
1. Run the packaging script:
```
./package.sh
```
2. In `beyond-sync` repo create release and archive generated tar.gz archive; in `install.sh` change the release download path

# Swagger
To check API documentation one can access swagger with following link:
```
<beyond-sync host>:8080/swagger/index.html#/
```
To generate/update swagger docs one can run:
```
go get -u github.com/swaggo/gin-swagger
swag init -g cmd/api/main.go -o docs
```