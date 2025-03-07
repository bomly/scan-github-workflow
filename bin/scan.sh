#!/bin/sh

echo "Running scan..."
result=$(/scan | base64 -w 0)
echo "RESULT=${result}" >> $GITHUB_OUTPUT
echo "---------------------------------"
echo "Found dependencies:"
echo $result | base64 -d
echo "---------------------------------"
echo "Displaying file system:"
find . -type f
