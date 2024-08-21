#!/bin/bash
# chmod +x create.sh && ./create.sh

# # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                       #
#     created by: Finbarrs Oketunji <f@finbarrs.eu>     #
#     created on: 25/08/2023                            #
#                                                       #        
# # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# Create directories
mkdir -p cmd
mkdir -p internal/app/urlshortener
mkdir -p internal/domain/url
mkdir -p internal/infrastructure/persistence/memory
mkdir -p internal/infrastructure/persistence/sql
mkdir -p internal/infrastructure/api/http
mkdir -p internal/config
mkdir -p test/unit

# Create files
touch cmd/main.go
touch internal/app/urlshortener/application.go
touch internal/app/urlshortener/handlers.go
touch internal/app/urlshortener/errors.go
touch internal/domain/url/model.go
touch internal/domain/url/repository.go
touch internal/infrastructure/persistence/memory/url_repository.go
touch internal/infrastructure/persistence/sql/url_repository.go
touch internal/infrastructure/api/http/router.go
touch internal/infrastructure/api/http/handlers.go
touch internal/config/config.go
touch test/unit/urlshortener_test.go
touch go.mod
touch go.sum

echo "Directory and file structure has been created."
