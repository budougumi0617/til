provider "aws" {
  region = "ap-northeast-1"
}

// 変数宣言。実行時に-var 'example_instance_type=t3.nano'というように変更できる。
variable "example_instance_type" {
  default = "t3.micro"
}

// コマンド実行時に上書きできない。
locals {
  example_instance_type = "t3.micro"
}

// 外部データを参照する。filterで検索条件を指定して最新のAMIを取得しているだけ。
data "aws_ami" "recent_amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-2.0.????????-x86_64-gp2"]
  }

  filter {
    name   = "state"
    values = ["available"]
  }
}

resource "aws_instance" "example" {
  ami           = data.aws_ami.recent_amazon_linux_2.image_id
  instance_type = var.example_instance_type
  // instance_type = locals.example_instance_type


  tags = {
    Name = "example"
  }
  user_data = <<EOF
    #!/bin/bash
    yum install -y httpd
    systemctl start httpd.service
EOF
}

/**
 * apply時に実行結果の最後に結果が表示される。
 *
 * $terraform apply
 * aws_instance.example: Refreshing state... [id=i-057617a019734946a]

 * Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

 * Outputs:

 * example_instance_id = i-057617a019734946a
*/
output "example_instance_id" {
  value = aws_instance.example.id
}
