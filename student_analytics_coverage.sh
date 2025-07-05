#!/bin/bash

# 📦 Test Coverage Script for Student Analytics & Placement Suite
# Usage: ./student_analytics_coverage.sh

# Set of specific test targets
FILES=(
    "analytics1_test.go"
    "analytics2_test.go"
    "analytics3_test.go"
    "analytics1.go"
    "analytics3.go"
    "analytics2.go"
    "main.go"
)

echo "🧪 Running tests and generating full coverage profile..."
go test -coverprofile=coverage_all.out ./internal

if [ $? -ne 0 ]; then
    echo "❌ Tests failed. Aborting."
    exit 1
fi

echo "📊 Filtering coverage for student analytics & placement suite only..."

# Construct pattern to match only relevant Go files
PATTERN=""
for file in "${FILES[@]}"; do
    if [ -z "$PATTERN" ]; then
        PATTERN="$file"
    else
        PATTERN="$PATTERN|$file"
    fi
done

# Extract relevant lines from full coverage report
awk -v pat="$PATTERN" 'NR==1 || $0 ~ pat' coverage_all.out > filtered_coverage.out

# Display summary
echo "🎯 Student Analytics & Placement Module Coverage:"
go tool cover -func=filtered_coverage.out

TOTAL_COVERAGE=$(go tool cover -func=filtered_coverage.out | grep total | awk '{print $3}')

echo ""
echo "🧾 TOTAL COVERAGE: $TOTAL_COVERAGE"

# Optional HTML output
read -p "🌐 Generate HTML coverage report? (y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    go tool cover -html=filtered_coverage.out -o coverage_report.html
    echo "📄 HTML report saved as: coverage_report.html"
fi

# Clean up (optional)
# rm -f filtered_coverage.out

echo "✅ Done!"
