#!/bin/bash

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 默认值
HOST=${1:-"localhost:8080"}
GEMINI_KEY=${GEMINI_API_KEY:-""}

# 检查 jq 是否安装
if ! command -v jq &> /dev/null; then
    echo -e "${YELLOW}Warning: jq is not installed. JSON responses will not be formatted.${NC}"
    JQ_INSTALLED=false
else
    JQ_INSTALLED=true
fi

# 辅助函数
format_json() {
    if [ "$JQ_INSTALLED" = true ]; then
        jq '.'
    else
        cat
    fi
}

print_header() {
    echo -e "\n${YELLOW}=== $1 ===${NC}"
}

check_response() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✓ Test passed${NC}"
    else
        echo -e "${RED}✗ Test failed${NC}"
        exit 1
    fi
}

# 测试 CORS 预检请求
print_header "Testing CORS preflight request"
curl -s -I -X OPTIONS "http://$HOST/api/generate" \
    -H "Origin: http://localhost:3000" \
    -H "Access-Control-Request-Method: POST" \
    -H "Access-Control-Request-Headers: Content-Type"
check_response $?

# 测试基本名字生成
print_header "Testing basic name generation"
curl -s -X POST "http://$HOST/api/generate" \
    -H "Content-Type: application/json" \
    -d '{"english_name":"Michael"}' | format_json
check_response $?

# 测试带空格的名字
print_header "Testing name with spaces"
curl -s -X POST "http://$HOST/api/generate" \
    -H "Content-Type: application/json" \
    -d '{"english_name":"John Smith"}' | format_json
check_response $?

# 测试特殊字符
print_header "Testing name with special characters"
curl -s -X POST "http://$HOST/api/generate" \
    -H "Content-Type: application/json" \
    -d '{"english_name":"Mary-Jane"}' | format_json
check_response $?

# 测试错误情况
print_header "Testing empty name (should fail)"
curl -s -X POST "http://$HOST/api/generate" \
    -H "Content-Type: application/json" \
    -d '{"english_name":""}' | format_json
check_response $?

# 测试无效的 JSON
print_header "Testing invalid JSON (should fail)"
curl -s -X POST "http://$HOST/api/generate" \
    -H "Content-Type: application/json" \
    -d '{invalid json}' | format_json
check_response $?

echo -e "\n${GREEN}All tests completed!${NC}" 