#!/bin/sh
PATH=/usr/local/bin:/usr/local/sbin:~/bin:/usr/bin:/bin:/usr/sbin:/sbin
git pull
go run main.go
git add .
git commit -m "Update automatically by script"
git push
