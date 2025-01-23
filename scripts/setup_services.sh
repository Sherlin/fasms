#!/bin/bash

# Install frontend dependencies
cd /home/ec2-user/frontend
npm install --legacy-peer-deps
npm run build


# Ensure backend dependencies are ready (Go modules)
cd /home/ec2-user/backend
go mod tidy

# Create backend systemd service
cat <<EOF | sudo tee /etc/systemd/system/backend.service
[Unit]
Description=Golang HTTP Server
After=network.target

[Service]
ExecStart=/usr/bin/go run /home/ec2-user/backend/cmd/main.go
Restart=always
User=ec2-user
WorkingDirectory=/home/ec2-user/backend

[Install]
WantedBy=multi-user.target
EOF

# Create frontend systemd service
cat <<EOF | sudo tee /etc/systemd/system/frontend.service
[Unit]
Description=React Frontend Application
After=network.target

[Service]
ExecStart=/usr/bin/npm start
Restart=always
User=ec2-user
WorkingDirectory=/home/ec2-user/frontend
Environment=PORT=3000

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd to recognize new services
sudo systemctl daemon-reload
sudo systemctl enable backend.service
sudo systemctl enable frontend.service