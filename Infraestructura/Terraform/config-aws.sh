#!/bin/bash
sudo apt-get update -y
#uninstall versions of docker
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove -y $pkg; done
#Install docker and dependencies
sudo apt-get install -y  ca-certificates curl
sudo install -y  -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update -y
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
sudo apt-get install -y docker-compose
sudo systemctl restart docker.socket
sudo apt-get install -y openssl
sudo apt-get install -y git
mkdir hyperledger && cd hyperledger
wget https://github.com/hyperledger/firefly-cli/releases/download/v1.3.0/firefly-cli_1.3.0_Linux_x86_64.tar.gz
sudo tar -zxf $PWD/firefly-cli_*.tar.gz -C /usr/local/bin ff && rm $PWD/firefly-cli_*.tar.gz
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
./install-fabric.sh
wget https://github.com/hyperledger/fabric/releases/download/v2.5.8/hyperledger-fabric-linux-amd64-2.5.8.tar.gz
mkdir hyperledger-binaries
tar -zxf hyperledger-fabric-linux*.tar.gz -C hyperledger-binaries  && cd hyperledger-binaries
cd bin && sudo cp * /usr/bin && cd ..
cd ..
#git clone https://github.com/hyperledger/fabric-samples.git
mkdir smart-contracts
cd smart-contracts
cd
sudo systemctl restart docker.socket
ff init test-prod1 2 -b fabric -p 8000
sudo systemctl daemon-reload
source ~/.profile
source ~/.bashrc
sleep(600)
ff start test-prod1 -v
