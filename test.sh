#!/bin/bash

echo "🧪 Complete Trading Platform Test..."

# Step 1: Signup (Create User)
echo "1. Creating user with signup..."
SIGNUP_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}')

echo "Signup Response:"
echo $SIGNUP_RESPONSE | jq .

# Check if signup was successful
ACCESS_TOKEN=$(echo $SIGNUP_RESPONSE | jq -r '.access_token')
if [ "$ACCESS_TOKEN" = "null" ] || [ "$ACCESS_TOKEN" = "" ]; then
  echo "❌ Signup failed, trying login instead..."
  
  # Step 2: Try Login (User might already exist)
  echo "2. Trying login..."
  LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
    -H "Content-Type: application/json" \
    -d '{"email":"test@example.com","password":"password123"}')
  
  echo "Login Response:"
  echo $LOGIN_RESPONSE | jq .
  
  ACCESS_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.access_token')
  REFRESH_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.refresh_token')
else
  echo "✅ Signup successful!"
  REFRESH_TOKEN=$(echo $SIGNUP_RESPONSE | jq -r '.refresh_token')
fi

if [ "$ACCESS_TOKEN" = "null" ] || [ "$ACCESS_TOKEN" = "" ]; then
  echo "❌ Failed to get access token"
  exit 1
fi

echo "3. Got tokens successfully!"
echo "Access Token: ${ACCESS_TOKEN:0:50}..."
echo "Refresh Token: ${REFRESH_TOKEN:0:50}..."

# Step 4: Test Protected Endpoints
echo "4. Testing protected endpoints..."

echo "Testing Holdings:"
curl -s -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:8080/api/v1/holdings | jq .

echo "Testing Orderbook:"
curl -s -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:8080/api/v1/orderbook | jq .

echo "Testing Positions:"
curl -s -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:8080/api/v1/positions | jq .

# Step 5: Test Refresh Token
echo "5. Testing refresh token..."
REFRESH_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d "{\"refresh_token\":\"$REFRESH_TOKEN\"}")

echo "Refresh Response:"
echo $REFRESH_RESPONSE | jq .

NEW_ACCESS_TOKEN=$(echo $REFRESH_RESPONSE | jq -r '.access_token')
if [ "$NEW_ACCESS_TOKEN" != "null" ] && [ "$NEW_ACCESS_TOKEN" != "" ]; then
  echo "✅ Refresh token flow successful!"
  echo "New Access Token: ${NEW_ACCESS_TOKEN:0:50}..."
else
  echo "❌ Refresh token flow failed"
fi

echo "🎉 All tests completed!"