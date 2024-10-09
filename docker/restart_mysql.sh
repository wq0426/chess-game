#!/bin/bash

docker run --name mysql-server -e MYSQL_ROOT_PASSWORD=123123 -d mysql:5.6
