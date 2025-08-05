#!/bin/bash

BASE_URL="http://localhost:8080"

echo "🔍 GET /"
curl -s -X GET "$BASE_URL" # | jq
echo ""

echo "🔍 GET /api/eshops"
curl -s -X GET "$BASE_URL/api/eshops" # | jq
echo ""

echo "🔍 GET /api/merchants"
curl -s -X GET "$BASE_URL/api/merchants" # | jq
echo ""

