variable "ec2-id" {
  type = string
}

resource "aws_eip" "eip" {
  instance = var.ec2-id
}