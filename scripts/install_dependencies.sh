#!/bin/bash
# Update and install Node.js, npm, and Go
#sudo apt-get update -y
mkdir -p /home/ec2-user/frontend
mkdir -p /home/ec2-user/backend
sudo yum install -y nodejs npm golang

# Install frontend dependencies
cd /home/ec2-user/frontend
npm install
npm run build


# Ensure backend dependencies are ready (Go modules)
cd /home/ec2-user/backend
go mod tidy
