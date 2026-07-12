#!/bin/bash
BASE_URL="http://localhost:8080"

# ============================================================
# 1. SUCCESSFUL CASES
# ============================================================

# Create a pending bill
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-001",
  "Title": "Electricity Bill",
  "Description": "Monthly electricity payment",
  "Amount": 150.00,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-08-01T00:00:00Z"
}' | python3 -m json.tool

# Create a paid bill
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-002",
  "Title": "Internet Bill",
  "Description": "Monthly internet subscription",
  "Amount": 79.99,
  "Currency": "USD",
  "Status": 1,
  "DueDate": "2026-07-15T00:00:00Z"
}' | python3 -m json.tool

# Create a cancelled bill
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-003",
  "Title": "Gym Membership",
  "Description": "Cancelled membership",
  "Amount": 49.99,
  "Currency": "USD",
  "Status": 2,
  "DueDate": "2026-07-20T00:00:00Z"
}' | python3 -m json.tool

# ============================================================
# 2. FAILED CASES (400 Bad Request)
# ============================================================

# Invalid JSON (missing closing brace)
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-004",
  "Title": "Broken Bill"
  "Amount": 100.00
}' | python3 -m json.tool

# Empty body
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '' | python3 -m json.tool

# Wrong content type
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: text/plain" -d 'not json' | python3 -m json.tool

# Null body
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d 'null' | python3 -m json.tool

# ============================================================
# 3. EDGE CASES
# ============================================================

# Zero amount
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-001",
  "Title": "Zero Amount Bill",
  "Description": "Bill with zero amount",
  "Amount": 0,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool

# Negative amount
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-002",
  "Title": "Negative Amount Bill",
  "Description": "Refund bill",
  "Amount": -50.00,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool

# Very large amount
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-003",
  "Title": "Large Amount Bill",
  "Description": "Bill with very large amount",
  "Amount": 999999999.99,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool

# Empty string fields
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "",
  "Title": "",
  "Description": "",
  "Amount": 10.00,
  "Currency": "",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool

# Missing all optional fields (only BillNo provided)
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-004"
}' | python3 -m json.tool

# Past due date
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-005",
  "Title": "Past Due Bill",
  "Description": "Bill with past due date",
  "Amount": 25.00,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2020-01-01T00:00:00Z"
}' | python3 -m json.tool

# Invalid status value (99)
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-006",
  "Title": "Invalid Status Bill",
  "Description": "Bill with invalid status",
  "Amount": 30.00,
  "Currency": "USD",
  "Status": 99,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool

# Unknown/extra fields (should be ignored by JSON binding)
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-007",
  "Title": "Extra Fields Bill",
  "Description": "Has unknown fields",
  "Amount": 40.00,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z",
  "HackerField": "drop table bill;"
}' | python3 -m json.tool

# Extremely long strings
curl -s -X POST "$BASE_URL/bills" -H "Content-Type: application/json" -d '{
  "BillNo": "INV-EDGE-008",
  "Title": "'$(python3 -c "print('A' * 1000)")'",
  "Description": "Normal description",
  "Amount": 50.00,
  "Currency": "USD",
  "Status": 0,
  "DueDate": "2026-07-25T00:00:00Z"
}' | python3 -m json.tool
