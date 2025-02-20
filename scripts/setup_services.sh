#!/bin/bash

# Install frontend dependencies
sudo chown -R ec2-user: /home/ec2-user/frontend
cp -r /home/ec2-user/frontend/build /var/www/html/

sudo chown -R ec2-user: /home/ec2-user/backend
sudo chmod 0755 /home/ec2-user/backend/fasms_exe
# Ensure backend dependencies are ready (Go modules)
cd /home/ec2-user/backend
#go mod tidy

# Create backend systemd service
cat <<EOF | sudo tee /etc/systemd/system/backend.service
[Unit]
Description=Golang HTTP Server
After=network.target

[Service]
ExecStart=/home/ec2-user/backend/fasms_exe
Restart=always
User=ec2-user
WorkingDirectory=/home/ec2-user/backend
RestartSec=90
EnvironmentFile=/home/ec2-user/.bash_profile

[Install]
WantedBy=multi-user.target
EOF



# Reload systemd to recognize new services
sudo systemctl daemon-reload
sudo systemctl enable backend.service
sudo systemctl restart nginx
#sudo systemctl enable frontend.service