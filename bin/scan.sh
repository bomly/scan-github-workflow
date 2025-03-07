#!/bin/sh

echo "Running scan..."
result=$(/scan | base64 -w 0)
echo "RESULT=${result}" >> $GITHUB_OUTPUT
