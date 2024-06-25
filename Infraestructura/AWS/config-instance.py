import subprocess
import json
import boto3

def run_terraform():
    # Inicializa Terraform
    subprocess.run(["terraform", "init"], check=True)

    # Aplica la configuración de Terraform y captura la salida
    result = subprocess.run(["terraform", "apply", "-auto-approve", "-json"], capture_output=True, text=True, check=True)

    # Analiza la salida JSON para obtener el instance_id
    apply_output = json.loads(result.stdout)
    instance_id = None

    for resource in apply_output['resource_changes']:
        if resource['type'] == 'aws_instance':
            instance_id = resource['change']['after']['id']
            break

    if instance_id is None:
        raise Exception("No se pudo encontrar el instance_id en la salida de Terraform.")

    return instance_id

def execute_bash_script(instance_id):
    ssm_client = boto3.client('ssm')

    bash_script = """#!/bin/bash
export DEBIAN_FRONTEND=noninteractive
apt-get update -y
apt-get install -y docker.io
apt-get install -y docker-compose
apt-get install -y openssl
apt-get install -y git
mkdir -p /hyperledger && cd /hyperledger
wget https://github.com/hyperledger/firefly-cli/releases/download/v1.3.0/firefly-cli_1.3.0_Linux_x86_64.tar.gz
tar -zxf firefly-cli_1.3.0_Linux_x86_64.tar.gz -C /usr/local/bin ff && rm firefly-cli_1.3.0_Linux_x86_64.tar.gz
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
./install-fabric.sh
wget https://github.com/hyperledger/fabric/releases/download/v2.5.8/hyperledger-fabric-linux-amd64-2.5.8.tar.gz
tar -zxf hyperledger-fabric-linux-amd64-2.5.8.tar.gz && cd hyperledger-fabric-linux-amd64-2.5.8
cd bin && cp * /usr/bin && cd ..
cd ..
git clone https://github.com/hyperledger/fabric-samples.git
mkdir smart-contracts
cd
ff init test-prod1 2 -b "fabric" -p 8000
ff start test-prod1 -v
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
    instance_id = run_terraform()
    print(f"Created instance with ID: {instance_id}")
    execute_bash_script(instance_id)
