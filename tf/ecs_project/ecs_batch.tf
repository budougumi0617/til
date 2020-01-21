# 第10章 バッチ
# バッチで大事なこと
# エラーハンドリング。失敗したときのエラー通知。ロギングもないと調査できない
# リトライ。バッチが失敗したときにリトライデキること。自動でリトライできれば理想。最低限手動で動くように。再実行可能な設計になっていること。
# 依存関係制御。バッチごとに順序があるならば、制御する必要がある。時間をずらすだけでは失敗する可能性があるのでアンチパターン

# 10.1 バッチ用CloundWatch Logs
# バッチごとに作ったほうが運用はラク
resource "aws_cloudwatch_log_group" "for_ecs_scheduled_tasks" {
  name              = "/ecs-scheduled-tasks/example"
  retention_in_days = 180
}

# 10.2 バッチ用タスク定義
resource "aws_ecs_task_definition" "example_batch" {
  family                   = "example-batch"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  # JSON以外は "aws_ecs_task_definition" "example" と同じ
  # 時刻を刻むだけのバッチ
  container_definitions = file("./batch_container_definitions.json")
  execution_role_arn    = module.ecs_task_execution_role.iam_role_arn
}


