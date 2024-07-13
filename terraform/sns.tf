resource "aws_sns_topic" "count_topic" {
  name = "counterapi-count-topic"
  tags = {
    Service = "CounterAPI"
  }
}