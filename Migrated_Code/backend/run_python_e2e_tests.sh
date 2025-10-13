#!/bin/bash

echo "============================================"
echo "OBP API - Python E2E Test Runner"
echo "============================================"
echo ""

if [ ! -d "../../python-e2e-tests" ]; then
    echo "❌ Error: Python E2E tests not found at ../../python-e2e-tests"
    echo "Please extract the python-e2e-tests.tar.gz file first"
    exit 1
fi

if ! curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "❌ Error: Backend server is not running at localhost:8080"
    echo ""
    echo "Please start the server first:"
    echo "  cd Migrated_Code/backend"
    echo "  go run cmd/main.go"
    echo ""
    exit 1
fi

echo "✅ Backend server is running"
echo ""

cd ../../python-e2e-tests || exit 1

if [ ! -d "venv" ]; then
    echo "Creating Python virtual environment..."
    python3 -m venv venv
fi

source venv/bin/activate

echo "Installing Python dependencies..."
pip install -q -r requirements.txt

echo ""
echo "Running Python E2E tests..."
echo "============================================"
echo ""

echo "Test 1: DirectLogin Authentication"
python3 -m pytest tests/test_authentication.py::TestDirectLoginAuthentication::test_valid_authentication -v -s

echo ""
echo "Test 2: Bank Operations with Authorization"
python3 -m pytest tests/test_bank_operations.py::TestBankEndpoints::test_get_banks_authorized -v -s

echo ""
echo "============================================"
echo "To run all tests: python3 -m pytest tests/ -v"
echo "============================================"
