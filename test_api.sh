#!/bin/bash

API_URL="http://localhost:8080"
POST_ENDPOINT="/post"
GET_ENDPOINT="/get"

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "❌ jq is required but not installed. Install it using: sudo apt install jq"
    exit 1
fi

# Test POST request
echo "Testing POST request..."
RESPONSE=$(curl -s -X POST "$API_URL$POST_ENDPOINT" \
     -H "Content-Type: application/json" \
     -d '{"name": "Sheila", "age": 25}')

# Extract the user ID
USER_ID=$(echo "$RESPONSE" | jq -r '.id')

if [ "$USER_ID" == "null" ]; then
    echo "❌ POST request failed! Response:"
    echo "$RESPONSE"
    exit 1
fi

echo "✅ POST request successful! User ID: $USER_ID"

# Test GET request
echo "Testing GET request..."
RESPONSE=$(curl -s -X GET "$API_URL$GET_ENDPOINT/$USER_ID")

if echo "$RESPONSE" | jq . > /dev/null 2>&1; then
    echo "✅ GET request successful! Response:"
    echo "$RESPONSE"
else
    echo "❌ GET request failed! Response:"
    echo "$RESPONSE"
    exit 1
fi
