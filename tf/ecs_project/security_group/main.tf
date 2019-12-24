# セキュリティグループの名前
variable "name" {}
# VPCのID
variable "vpc_id" {}
# 通信を許可するポート番号
variable "port" {}

# 通信を許可するCIDRブロック
variable "cidr_blocks" {
  # 文字列配列以外を不許可。
  type = list(string)
}

resource "aws_security_group" "default" {
  name   = var.name
  vpc_id = var.vpc_id
}

resource "aws_security_group_rule" "ingress" {
  type              = "ingress"
  from_port         = var.port
  to_port           = var.port
  protocol          = "tcp"
  cidr_blocks       = var.cidr_blocks
  security_group_id = aws_security_group.default.id
}

resource "aws_security_group_rule" "egress" {
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.default.id
}

output "security_group_id" {
  value = aws_security_group.default.id
}