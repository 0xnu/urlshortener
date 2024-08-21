#!/bin/bash
# chmod +x start.sh && ./start.sh

# # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                       #
#     created by: Finbarrs Oketunji <f@finbarrs.eu>     #
#     created on: 25/08/2023                            #
#                                                       #        
# # # # # # # # # # # # # # # # # # # # # # # # # # # # #

export HTTP_PORT=8090
export MYSQL_USER=root
export MYSQL_PASSWORD=rootroot
export MYSQL_HOST=localhost
export MYSQL_PORT=3306
export MYSQL_DATABASE=url_shortener

go mod init urlshortener
go mod tidy
go run cmd/main.go
