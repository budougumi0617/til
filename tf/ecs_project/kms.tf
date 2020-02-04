# 11.1.1 カスタマーマスターキーの定義
resource "aws_kms_key" "example" {
  description             = "Example Customer Master Key"
  # 年に1回のローテーション
  enable_key_rotation     = true
  is_enabled              = true
  # 本当に削除してしまうまでの期間。
  # 削除はせずに、is_enmabledをfalseにして無効化するくらいのほうがよい。
  deletion_window_in_days = 30
}