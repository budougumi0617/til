# 10.5 CloudWatchイベントルールの定義
resource "aws_cloudwatch_event_rule" "example_batch" {
  name                = "example-batch"
  description         = "とても重要なバッチ処理です"
  # 2分おきに実行するさ（"?"は任意の曜日）
  schedule_expression = "cron(*/2 * * * ? *)"
}

# 10.6 CloudWatchイベントターゲットの定義
resource "aws_cloudwatch_event_target" "example_batch" {
  target_id = "example-batch"
  # スケジューリングルール
  rule      = aws_cloudwatch_event_rule.example_batch.name
  # 10.4で作ったやつ
  role_arn  = module.ecs_events_role.iam_role_arn
  arn       = aws_ecs_cluster.example.arn

  # ロードバランサやヘルスチェックの設定はない
  ecs_target {
    launch_type         = "FARGATE"
    task_count          = 1
    platform_version    = "1.3.0"
    # バッチ
    task_definition_arn = aws_ecs_task_definition.example_batch.arn

    network_configuration {
      assign_public_ip = "false"
      subnets          = [aws_subnet.private_0.id]
    }
  }
}