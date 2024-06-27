import paramiko
import time

def execute_commands_via_ssh(host, username, key_file, commands):
    # Crear un cliente SSH
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    # Conectarse al servidor
    try:
        ssh.connect(hostname=host, username=username, key_filename=key_file)
        print(f"Conectado a {host}")

        for command in commands:
            stdin, stdout, stderr = ssh.exec_command(command)
            print(f"Ejecutando comando: {command}")
            
            # Esperar a que el comando termine y obtener la salida
            stdout.channel.recv_exit_status()
            output = stdout.read().decode()
            errors = stderr.read().decode()
            
            print("Salida:", output)
            if errors:
                print("Errores:", errors)

        # Cerrar la conexión SSH
        ssh.close()
        print("Conexión SSH cerrada")
    except Exception as e:
        print(f"Error al conectar o ejecutar comandos: {e}")

if __name__ == "__main__":
    host = "34.202.27.130"
    username = "ubuntu"  # O el usuario adecuado para tu instancia EC2
    key_file = "../Terraform/keys/id_rsa"  # Ruta a tu clave privada
    
    commands = [
        "export DEBIAN_FRONTEND=noninteractive",
        "sudo apt-get update -y",
        "for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove -y $pkg; done",
        "sudo apt-get install -y ca-certificates curl",
        "sudo install -y -m 0755 -d /etc/apt/keyrings",
        "sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc",
        "sudo chmod a+r /etc/apt/keyrings/docker.asc",
        "echo 'deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable' | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null",
        "sudo apt-get update -y",
        "sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
        "sudo groupadd docker",
        "sudo usermod -aG docker ubuntu",
        "newgrp docker",
        "sudo apt-get install -y docker-compose",
        "sudo systemctl restart docker.socket",
        "sudo apt-get install -y openssl",
        "sudo apt-get install -y git",
        "mkdir /home/ubuntu/hyperledger && cd /home/ubuntu/hyperledger",
        "wget https://github.com/hyperledger/firefly-cli/releases/download/v1.3.0/firefly-cli_1.3.0_Linux_x86_64.tar.gz",
        "sudo tar -zxf /home/ubuntu/hyperledger/firefly-cli_*.tar.gz -C /usr/local/bin ff && rm /home/ubuntu/hyperledger/firefly-cli_*.tar.gz",
        "curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh",
        "./install-fabric.sh",
        "wget https://github.com/hyperledger/fabric/releases/download/v2.5.8/hyperledger-fabric-linux-amd64-2.5.8.tar.gz",
        "mkdir /home/ubuntu/hyperledger/hyperledger-binaries",
        "tar -zxf hyperledger-fabric-linux*.tar.gz -C /home/ubuntu/hyperledger/hyperledger-binaries && cd /home/ubuntu/hyperledger/hyperledger-binaries",
        "cd bin && sudo cp * /usr/bin && cd ..",
        "cd ..",
        "mkdir /home/ubuntu/hyperledger/smart-contracts",
        "cd /home/ubuntu/hyperledger/smart-contracts",
        "cd /home/ubuntu/",
        "sudo systemctl enable docker.socket",
        "sudo systemctl restart docker.socket",
        "sudo systemctl daemon-reload",
        "sleep 180",
        "cd /home/ubuntu/",
        "source .bashrc",
        "source .profile",
        "sleep 180",
        "ff --help",
        "ff init test 2 -b 'fabric' -p 8000",
        "ff start test -v -b"
    ]

    execute_commands_via_ssh(host, username, key_file, commands)
