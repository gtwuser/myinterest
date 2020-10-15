provider "aws" {
  region = "us-east-1"
}

module "ec2" {
  source = "./server"
  instance_id = data.aws_ami.ami_instance.id
  inst_type = "t2.micro"
  instance_name = data.aws_ami.ami_instance.name
}

output "ami_details" {
  value = module.ec2.ec2-output
}

data "aws_ami" "ami_instance" {
  most_recent = true
  owners = [
    "amazon"]

  filter {
    name = "name"
    values = [
      "amzn-ami-hvm*"]
  }

  filter {
    name = "virtualization-type"
    values = ["hvm"]
  }

}
