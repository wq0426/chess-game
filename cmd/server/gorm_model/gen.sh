#!/bin/bash

go run main.go -db=$1 -tb=$2 -dsn='root:123123@tcp(127.0.0.1:3306)/'$1'?charset=utf8mb4&parseTime=True&loc=Local'
