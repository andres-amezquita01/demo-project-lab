#!/bin/bash
cluster_name=$1
task_arn=$2
task_description=$(aws ecs describe-tasks --cluster $cluster_name --tasks $task_arn)
private_ip=$(echo $task_description | jq -r '.tasks[0].attachments[0].details[] | select(.name=="privateIPv4Address") | .value')
echo "private staging ip: $private_ip"