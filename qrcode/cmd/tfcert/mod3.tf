provider "aws" {
  region = "us-east-1"
}

variable "private_key_path" {
  type = string
}
variable "key_name" {
  type = string
}

data "aws_ami" "ami-linux" {
  most_recent = true
  owners = ["amazon"]
  filter {
    name = "name"
    values = ["amzn-ami-hvm*"]
  }
  filter {
    name = "virtualization-type"
    values = ["hvm"]
  }
  filter {
    name = "root-device-type"
    values = ["ebs"]
  }
}

resource "aws_default_vpc" "aws-vpc" {}

variable "igr" {
  type = list(number)
  default = [22,80]
}

variable "egr" {
  type = list(number)
  default = [0]
}

resource "aws_security_group" "sg" {
  name = "first-sg"
  vpc_id = aws_default_vpc.aws-vpc.id
  dynamic "ingress" {
    iterator = port
    for_each = var.igr
    content {
      to_port = port.value
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
      from_port = port.value
    }
  }
  dynamic "egress" {
    iterator = port
    for_each = var.egr
    content {
      to_port = port.value
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
      from_port = port.value
    }
  }
}


resource "aws_instance" "aws-linux" {
  ami = data.aws_ami.ami-linux.id
  instance_type = "t2.micro"
  vpc_security_group_ids = [aws_security_group.sg.id]
  key_name = var.key_name
  connection {
    host = self.public_ip
    user = "ec2-user"
    type = "ssh"
    private_key = file(var.private_key_path)
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum install nginx -y",
      "sudo service nginx start"
    ]
  }
}

output "aws_linux_details" {
  value = aws_instance.aws-linux.public_dns
}