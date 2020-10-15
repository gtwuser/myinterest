variable "ingress_rule" {
  type = list(number)
}

variable "egress_rule" {
  type = list(number)
}

resource "aws_security_group" "sg" {
  name = "first sg"
  dynamic "ingress" {
    iterator = port
    for_each = var.ingress_rule
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = [
        "0.0.0.0/0"]
    }
  }

  dynamic "egress" {
    iterator = port
    for_each = var.egress_rule
    content {
      from_port = port.value
      to_port = port.value
      cidr_blocks = [
        "0.0.0.0/0"]
      protocol = "TCP"
    }
  }
}

output "sg-name" {
  value = aws_security_group.sg.name
}
