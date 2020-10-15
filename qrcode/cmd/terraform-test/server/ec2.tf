variable "instance_id" {
  type = string
}

variable "instance_name" {
  type = string
}

variable "inst_type" {
  type = string
}
resource "aws_instance" "ec2" {
  ami = var.instance_id
  instance_type = var.inst_type
  security_groups = [
    module.sg.sg-name]
  tags = {
    Name = "New Server"
  }
}


module "eip" {
  source = "../eip"
  ec2-id = aws_instance.ec2.id
}

module "sg" {
  source = "../sg"
  ingress_rule = [8080, 8081, 8002]
  egress_rule = [8083, 8084, 8003]
}

output "ec2-output" {
  value = [
    var.instance_id,
    var.instance_name,
    aws_instance.ec2.public_ip,
    //    data.aws_instance.new_server.availability_zone
  ]
}