#!/bin/bash

BASE_URL="http://localhost:8080"

echo "Getting all e-shops:"
curl -s http://localhost:8080/api/eshops/ | jq
echo ""

echo "🔍 GET /api/eshops"
curl -s -X GET "$BASE_URL/api/eshops/" | jq
echo ""

echo "🔍 GET /api/merchants"
curl -s -X GET "$BASE_URL/api/merchants/" | jq
echo ""

