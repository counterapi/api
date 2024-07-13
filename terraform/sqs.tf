resource "aws_sqs_queue" "count_queue" {
  name                      = "counterapi-count-queue"
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

  tags = {
    Service = "CounterAPI"
  }
}

resource "aws_sqs_queue_policy" "example_queue_policy" {
  queue_url = aws_sqs_queue.count_queue.id

  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Id" : "sqspolicy",
      "Statement" : [
        {
          "Sid" : "001",
          "Effect" : "Allow",
          "Principal" : "*",
          "Action" : "sqs:SendMessage",
          "Resource" : aws_sqs_queue.count_queue.arn,
          "Condition" : {
            "ArnEquals" : {
              "aws:SourceArn" : aws_sqs_queue.count_queue.arn
            }
          }
        }
      ]
  })
}

resource "aws_sns_topic_subscription" "results_updates_sqs_target" {
  topic_arn = aws_sns_topic.count_topic.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.count_queue.arn
}