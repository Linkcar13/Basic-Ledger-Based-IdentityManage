import subprocess
import json
import re
import time
import boto3
import os

def run_terraform(terraform_dir):
    os.chdir(terraform_dir)
    
    # Initialize and apply Terraform configuration
    init_result = subprocess.run(["terraform", "init"], capture_output=True, text=True)
    print(init_result.stdout)
    
    apply_result = subprocess.run(["terraform", "apply", "-auto-approve"], capture_output=True, text=True)
    output = apply_result.stdout
    print(output)
    
    # Extract instance_id from Terraform apply output using regular expressions
    match = re.search(r'aws_instance\.server-blockchain: Creation complete after \d+s \[id=(i-[a-z0-9]+)\]', output)
    if match:
        instance_id = match.group(1)
        return instance_id
    else:
        raise Exception("No se pudo encontrar el instance_id en la salida de Terraform.")

def wait_for_instance_running(instance_id):
    ec2_client = boto3.client('ec2',region_name="us-east-1" )
    waiter = ec2_client.get_waiter('instance_running')
    waiter.wait(InstanceIds=[instance_id])
    print(f"Instance {instance_id} is now running.")

def restart_instance(instance_id):
    ec2_client = boto3.client('ec2', region_name="us-east-1")
    print(f"Restarting instance {instance_id}...")
    ec2_client.reboot_instances(InstanceIds=[instance_id])
    print(f"Instance {instance_id} rebooted successfully.")

def execute_bash_script(instance_id):
    ssm_client = boto3.client('ssm', )
    bash_script = '''#!/bin/bash
export DEBIAN_FRONTEND=noninteractive
apt-get update -y
#uninstall versions of docker
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do apt-get remove -y $pkg; done
#Install docker and dependencies
apt-get install -y ca-certificates curl
sudo install -y -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update -y
apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo groupadd docker
sudo usermod -aG docker ubuntu
newgrp docker
apt-get install -y docker-compose
systemctl restart docker.socket
apt-get install -y openssl
apt-get install -y git
mkdir /home/ubuntu/hyperledger && cd /home/ubuntu/hyperledger
wget https://github.com/hyperledger/firefly-cli/releases/download/v1.3.0/firefly-cli_1.3.0_Linux_x86_64.tar.gz
sudo tar -zxf /home/ubuntu/hyperledger/firefly-cli_*.tar.gz -C /usr/local/bin ff && rm /home/ubuntu/hyperledger/firefly-cli_*.tar.gz
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
./install-fabric.sh
wget https://github.com/hyperledger/fabric/releases/download/v2.5.8/hyperledger-fabric-linux-amd64-2.5.8.tar.gz
mkdir /home/ubuntu/hyperledger/hyperledger-binaries
tar -zxf hyperledger-fabric-linux*.tar.gz -C /home/ubuntu/hyperledger/hyperledger-binaries && cd /home/ubuntu/hyperledger/hyperledger-binaries
cd bin && sudo cp * /usr/bin && cd ..
cd ..
#git clone https://github.com/hyperledger/fabric-samples.git
mkdir /home/ubuntu/hyperledger/smart-contracts
cd /home/ubuntu/hyperledger/smart-contracts
#clonar repo de smartcontracts
cd /home/ubuntu/
sudo systemctl enable docker.socket
sudo systemctl restart docker.socket
sudo systemctl daemon-reload
sleep 180
cd /home/ubuntu/
source .bashrc
source .profile
sleep 180
ff --help
ff init test 2 -b "fabric" -p 8000
ff start test -v -b'''
    response = ssm_client.send_command(
        InstanceIds=[instance_id],
        DocumentName="AWS-RunShellScript",
        Parameters={'commands': [bash_script]},
        TimeoutSeconds=600,
    )
    #command_id = response['Command']['CommandId']
    print("Command sent successfully.")
    return response

if __name__ == "__main__":
    terraform_directory = "../Terraform/"
    instance_id = run_terraform(terraform_directory)
    print(f"Created instance with ID: {instance_id}")

    time.sleep(300)

    wait_for_instance_running(instance_id)
    
    # Restart the instance before executing bash script
    restart_instance(instance_id)
    print("Waiting for instance to restart...")
    time.sleep(300)  # Adjust time as necessary
    print("Executing bash script...")

    try:
        execute_bash_script(instance_id)
    except Exception as e:
        print(f"Error executing bash script: {e}")
