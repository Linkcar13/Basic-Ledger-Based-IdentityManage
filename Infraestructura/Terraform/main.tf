provider "aws" {
  shared_config_files = ["~/.aws/config"]
  shared_credentials_files = ["~/.aws/credentials"]
}

resource "aws_vpc" "tesis_vpc" {
  #CIDR
  cidr_block = "10.0.0.0/16"

  tags = {
    Name: "VPC-Tesis"
  }

}

resource "aws_subnet" "tesis_public_subnet" {
  cidr_block = "10.0.1.0/24"
  vpc_id = aws_vpc.tesis_vpc.id
}

resource "aws_subnet" "tesis_private_subnet" {
  cidr_block = "10.0.2.0/24"
  vpc_id = aws_vpc.tesis_vpc.id
}

resource "aws_internet_gateway" "tesis_public_internet_gateway" {
  vpc_id = aws_vpc.tesis_vpc.id
}

resource "aws_route_table" "tesis_public_subnet_route_table" {
  vpc_id = aws_vpc.tesis_vpc.id

  route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.tesis_public_internet_gateway.id
  }
  route {
        ipv6_cidr_block = "::/0"
        gateway_id = aws_internet_gateway.tesis_public_internet_gateway.id
  }
}

resource "aws_route_table_association" "tesis_public_route_table_association" {
    
    route_table_id = aws_route_table.tesis_public_subnet_route_table.id
    subnet_id = aws_subnet.tesis_public_subnet.id

}

resource "aws_security_group" "tesis_blokchain_server_security_group" {
  vpc_id = aws_vpc.tesis_vpc.id
  
  ingress {
        description = "inbound rule for HTTP protocol"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
        description = "inbound rule for HTTPS  protocol"
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
        description = "inbound rule for HTTP alt protocol"
        from_port = 8000
        to_port = 8000
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
        description = "inbound rule for HTTP alt protocol"
        from_port = 8001
        to_port = 8001
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
        description = "inbound rule for HTTP alt protocol"
        from_port = 5108
        to_port = 5108
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
        description = "inbound rule for HTTP alt protocol"
        from_port = 5208
        to_port = 5208
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
  }
  
    ingress {
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
        description = "Allow all trafic"
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "blockchain security group"
    }

}

//Bloque de datos se usa para definir informacion especifica de un recurso
data "aws_ami" "ubuntu" {
  most_recent = "true"
  filter {
    name = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }
  filter {
    name = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] #canonical
}

resource "aws_instance" "server-blockchain" {

    ami = data.aws_ami.ubuntu.id
    instance_type = "t2.medium"

    root_block_device {
    volume_size = 20  # Tamaño del volumen en GB
    volume_type = "gp2"  # Tipo de volumen, puede ser gp2, io1, st1, sc1, etc.
    delete_on_termination = true  # Si el volumen debe ser eliminado cuando se termine la instancia
    }

    network_interface {
      network_interface_id = aws_network_interface.tesis_network_interface.id
      device_index = 0
    }

    tags = {
      Name = "Server-Prod-Blockchain"
    }
  
}

resource "aws_network_interface" "tesis_network_interface" {
  
  subnet_id = aws_subnet.tesis_public_subnet.id
  private_ips = ["10.0.1.100"]
  security_groups = [aws_security_group.tesis_blokchain_server_security_group.id]

    tags = {
    Name = "Network-Interface-Prod-Server"
  }

}

resource "aws_eip" "server_blockchain_elastic_ip" {

  associate_with_private_ip = tolist(aws_network_interface.tesis_network_interface.private_ips)[0]
  network_interface = aws_network_interface.tesis_network_interface.id
  instance = aws_instance.server-blockchain.id
  tags = {
    Name = "Elasctic-IP-For-Prod-Server"
  }
}

