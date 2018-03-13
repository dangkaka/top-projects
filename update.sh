#!/bin/sh
git pull
go run main.go
git add .
git commit -m "Update automatically by script"
git push
