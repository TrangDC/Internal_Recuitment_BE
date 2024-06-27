#!/bin/sh
echo "Running start.sh script..."
go run import_master_data.go
exec ./server api
