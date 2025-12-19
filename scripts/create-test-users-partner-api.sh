#!/usr/bin/env bash

# Script to create test users via Partner API
# This creates users directly and returns their API keys
# Then adds them to the specified account as Super Administrators
# User naming format: tf-acct-{product-group}-{test-job}@cfapi.net

set -e

# Configuration
USER_EMAIL_DOMAIN="@cfapi.net"
USER_BASE_NAME="tf-acct"
API_BASE_URL="https://api.cloudflare.com/client/v4"
ACCOUNT_ID="f037e56e89293a057740de681ac9abbe"
SUPER_ADMIN_ROLE_ID="05784afa30c1afe1440e79d9351c7430"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check required environment variables
check_requirements() {
    if [ -z "$CLOUDFLARE_API_KEY" ]; then
        echo -e "${RED}Error: CLOUDFLARE_API_KEY must be set${NC}"
        echo -e "${YELLOW}This is your partner/tenant admin API key${NC}"
        exit 1
    fi

    if [ -z "$CLOUDFLARE_EMAIL" ]; then
        echo -e "${RED}Error: CLOUDFLARE_EMAIL must be set${NC}"
        echo -e "${YELLOW}This is your partner/tenant admin email${NC}"
        exit 1
    fi

    if [ -z "$CLOUDFLARE_TENANT_ID" ]; then
        echo -e "${RED}Error: CLOUDFLARE_TENANT_ID must be set${NC}"
        echo -e "${YELLOW}This is your tenant ID (required if you manage multiple tenants)${NC}"
        exit 1
    fi

    echo -e "${GREEN}✓ Environment variables configured${NC}"
}


# Add user to account as Super Administrator
add_user_to_account() {
    local email="$1"

    echo -e "${BLUE}  Adding user to account ${ACCOUNT_ID} as Super Administrator...${NC}"

    local response=$(curl -s -w "\n%{http_code}" "${API_BASE_URL}/accounts/${ACCOUNT_ID}/members" \
        -X POST \
        -H "X-Auth-Email: ${CLOUDFLARE_EMAIL}" \
        -H "X-Auth-Key: ${CLOUDFLARE_API_KEY}" \
        -H "Content-Type: application/json" \
        --data-raw "{\"email\":\"${email}\",\"roles\":[\"${SUPER_ADMIN_ROLE_ID}\"],\"status\":\"accepted\"}")

    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')

    if [ "$http_code" == "200" ]; then
        echo -e "${GREEN}  ✓ User added to account successfully${NC}"
        return 0
    else
        echo -e "${RED}  ✗ Failed to add user to account${NC}"
        echo "  HTTP Code: $http_code"
        echo "  Response: $body"
        return 1
    fi
}

# Create a user via Partner API
create_user() {
    local email="$1"

    echo -e "${BLUE}Creating user: ${email}${NC}"

    local response=$(curl -s -w "\n%{http_code}" "${API_BASE_URL}/users" \
        -X POST \
        -H "X-Auth-Email: ${CLOUDFLARE_EMAIL}" \
        -H "X-Auth-Key: ${CLOUDFLARE_API_KEY}" \
        -H "Content-Type: application/json" \
        --data-raw "{\"email\":\"${email}\",\"tenant\":{\"id\":\"${CLOUDFLARE_TENANT_ID}\"}}")

    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')

    if [ "$http_code" == "200" ]; then
        local user_id=$(echo "$body" | jq -r '.result.id')
        local user_email=$(echo "$body" | jq -r '.result.email')
        local user_apikey=$(echo "$body" | jq -r '.result.api_key')

        echo -e "${GREEN}  ✓ User created successfully${NC}"
        echo "  User ID: ${user_id}"
        echo "  Email: ${user_email}"
        echo "  API Key: ${user_apikey}"

        # Add user to account
        if ! add_user_to_account "${user_email}"; then
            echo -e "${YELLOW}  ⚠ User created but failed to add to account${NC}"
        fi

        # Store to file
        echo "${user_email},${user_apikey}" >> test-users-credentials.csv

        return 0
    else
        echo -e "${RED}  ✗ Failed to create user${NC}"
        echo "  HTTP Code: $http_code"
        echo "  Response: $body"
        return 1
    fi
}

# Main execution
main() {
    local dry_run=false
    local product_group=""
    local test_job=""

    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            --dry-run)
                dry_run=true
                shift
                ;;
            --product-group)
                product_group="$2"
                shift 2
                ;;
            --test-job)
                test_job="$2"
                shift 2
                ;;
            --help|-h)
                cat <<EOF
Usage: $0 --product-group GROUP --test-job JOB [OPTIONS]

Create a test user via Cloudflare Partner API.
Users are created immediately with API keys - no password reset needed!
Each user is automatically added to account f037e56e89293a057740de681ac9abbe as Super Administrator.

Options:
  --product-group GROUP  Product group for the test user (e.g., workers, zone, security) [Required]
  --test-job JOB         Test job name (e.g., workers-core, zone-settings) [Required]
  --dry-run              Show what would be created without actually creating
  --help, -h             Show this help message

Environment Variables:
  CLOUDFLARE_API_KEY    Your partner/tenant admin API key [Required]
  CLOUDFLARE_EMAIL      Your partner/tenant admin email [Required]
  CLOUDFLARE_TENANT_ID  Your tenant ID [Required]

Output:
  Appends to test-users-credentials.csv with format: email,api_key

Examples:
  # Create a specific user
  $0 --product-group workers --test-job workers-core

  # Dry run to see what would be created
  $0 --product-group workers --test-job workers-core --dry-run

EOF
                exit 0
                ;;
            *)
                echo -e "${RED}Unknown option: $1${NC}"
                echo "Use --help for usage information"
                exit 1
                ;;
        esac
    done

    # Validate required arguments
    if [ -z "$product_group" ] || [ -z "$test_job" ]; then
        echo -e "${RED}Error: Both --product-group and --test-job are required${NC}"
        echo "Use --help for usage information"
        exit 1
    fi

    echo -e "${GREEN}=== Cloudflare Test User Creation (Partner API) ===${NC}"
    echo ""

    check_requirements
    echo ""

    local email="${USER_BASE_NAME}-${product_group}-${test_job}${USER_EMAIL_DOMAIN}"

    echo -e "${BLUE}Creating user: ${email}${NC}"
    echo ""

    if [ "$dry_run" = true ]; then
        echo -e "${YELLOW}DRY RUN MODE - No users will be created${NC}"
        echo ""
        echo "Would create user: ${email}"
        exit 0
    fi

    # Initialize or append to credentials file
    if [ ! -f test-users-credentials.csv ]; then
        echo "email,api_key" > test-users-credentials.csv
    fi
    echo -e "${BLUE}Credentials will be saved to: test-users-credentials.csv${NC}"
    echo ""

    if create_user "$email"; then
        echo ""
        echo -e "${GREEN}✓ User created successfully${NC}"
    else
        echo ""
        echo -e "${RED}✗ User creation failed${NC}"
        exit 1
    fi
}

main "$@"
