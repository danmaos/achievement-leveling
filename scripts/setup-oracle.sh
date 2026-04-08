#!/bin/bash
# Run this script on a fresh Oracle Cloud Ubuntu VM to set up the deployment environment.
# Usage: ssh ubuntu@<your-oracle-ip> 'bash -s' < scripts/setup-oracle.sh

set -euo pipefail

echo "=== Updating system ==="
sudo apt update && sudo apt upgrade -y

echo "=== Installing Docker ==="
sudo apt install -y docker.io docker-compose-v2
sudo systemctl enable docker
sudo systemctl start docker
sudo usermod -aG docker $USER

echo "=== Installing Git ==="
sudo apt install -y git

echo "=== Opening firewall ports ==="
sudo iptables -I INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 443 -j ACCEPT
sudo iptables -I INPUT -p tcp --dport 8080 -j ACCEPT
# Persist iptables rules
sudo apt install -y iptables-persistent
sudo netfilter-persistent save

echo "=== Cloning repository ==="
git clone https://github.com/danmaos/achievement-leveling.git ~/achievement-leveling
cd ~/achievement-leveling

echo "=== Setup complete! ==="
echo ""
echo "Next steps:"
echo "1. Create .env file:  cp .env.example .env && nano .env"
echo "2. Fill in GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, JWT_SECRET"
echo "3. Update GOOGLE_REDIRECT_URL to your server's public URL"
echo "4. Update FRONTEND_URL to your server's public URL"
echo "5. Start services:  docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build"
echo ""
echo "NOTE: Log out and back in for docker group membership to take effect."
