// 変数宣言。実行時に-var 'example_instance_type=t3.nano'というように変更できる。
variable "example_instance_type" {
  default = "t3.micro"
}

// コマンド実行時に上書きできない。
locals {
  example_instance_type = "t3.micro"
}

resource "aws_instance" "example" {
  ami           = "ami-0f9ae750e8274075b"
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
