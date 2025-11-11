#!/bin/bash

author="$1"
message="$2"

if [ -z "$author" ] || [ -z "$message" ]; then
  echo "Authpr and message are required, please see README.md for usage"
  exit 1
fi

for i in {1..10}; do
  sleep 1
  curl -s -X POST http://localhost:8080/messages \
    -H "Content-Type: application/json" \
    -d "{\"author\":\"${author}\",\"body\":\"${message} #${i}\"}"
done
