import subprocess
import json
import os
import boto3

def run_terraform(terraform_dir):
    # Cambia el directorio de trabajo a la carpeta de Terraform
    os.chdir(terraform_dir)

    # Inicializa Terraform
    subprocess.run(["terraform", "init"], check=True)

    # Aplica la configuración de Terraform y captura la salida
    result = subprocess.run(["terraform", "apply", "-auto-approve", "-json"], capture_output=True, text=True, check=True)
    
    # Vuelve al directorio original
    os.chdir("..")

    # Guardar la salida en una variable para su análisis
    output = result.stdout

    # Extraer solo la parte JSON de la salida
    json_output = extract_json_from_output(output)

    # Cargar el JSON extraído
    apply_output = json.loads(json_output)

    # Buscar el instance_id en el JSON
    instance_id = None
    for resource in apply_output['resource_changes']:
        if resource['type'] == 'aws_instance':
            instance_id = resource['change']['after']['id']
            break

    if instance_id is None:
        raise Exception("No se pudo encontrar el instance_id en la salida de Terraform.")

    return instance_id

def extract_json_from_output(output):
    """
    Extraer solo la parte JSON de la salida de Terraform
    """
    json_start = output.find('{')
    json_end = output.rfind('}') + 1
    json_output = output[json_start:json_end]
    return json_output

def execute_bash_script(instance_id):
    ssm_client = boto3.client('ssm')

    bash_script = """##!/bin/bash
export DEBIAN_FRONTEND=noninteractive
apt-get update -y
#uninstall versions of docker
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do apt-get remove -y $pkg; done
#Install docker and dependencies
apt-get install -y  ca-certificates curl
sudo install -y  -m 0755 -d /etc/apt/keyrings
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
tar -zxf hyperledger-fabric-linux*.tar.gz -C /home/ubuntu/hyperledger/hyperledger-binaries  && cd /home/ubuntu/hyperledger/hyperledger-binaries
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
ff start test -v -b
"""

    response = ssm_client.send_command(
        InstanceIds=[instance_id],
        DocumentName="AWS-RunShellScript",
        Parameters={'commands': [bash_script]}
    )

    command_id = response['Command']['CommandId']

    # Esperar a que el comando se complete
    ssm_client.get_waiter('command_executed').wait(
        CommandId=command_id,
        InstanceId=instance_id
    )

    # Obtener la salida del comando
    output = ssm_client.get_command_invocation(
        CommandId=command_id,
        InstanceId=instance_id
    )

    print("Command Output:")
    print(output['StandardOutputContent'])
    print("Command Errors:")
    print(output['StandardErrorContent'])

if __name__ == "__main__":
    terraform_directory = "../Terraform/"  # Directorio que contiene el script de Terraform
    instance_id = run_terraform(terraform_directory)
    print(f"Created instance with ID: {instance_id}")
    execute_bash_script(instance_id)
