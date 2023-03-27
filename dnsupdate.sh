#!/bin/bash

# Replace with your own values
DOMAIN="$DOMAIN"
USERNAME="$USERNAME"
PASSWORD="$PASSWORD"

# Get the current public IP address
IP=$(dig +short myip.opendns.com @resolver1.opendns.com)

# Update the Google domain using dynamic DNS
curl -s "https://domains.google.com/nic/update?hostname=$DOMAIN&myip=$IP" \
  --user "$USERNAME:$PASSWORD" \
  --header "User-Agent: ddclient/3.8.3" \
  --header "Content-Type: application/x-www-form-urlencoded"
