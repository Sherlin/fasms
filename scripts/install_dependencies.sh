#!/bin/bash
# Update and install Node.js, npm, and Go
#sudo apt-get update -y
mkdir -p /home/ec2-user/frontend
mkdir -p /home/ec2-user/backend
sudo yum install -y nodejs npm golang nginx
mkdir -p /var/www/html/
cp -r /home/ec2-user/frontend/build /var/www/html/

