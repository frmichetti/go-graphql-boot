#!/bin/sh
mkdir gin-graphql-postgres
cd gin-graphql-postgres
go mod init github.com/anjaneyulubatta505/gin-graphql-postgres
go get github.com/99designs/gqlgen
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go run github.com/99designs/gqlgen init
