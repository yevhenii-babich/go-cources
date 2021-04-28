package main

//go:generate go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
//go:generate oapi-codegen -generate types -o ./service/models/models.gen.go -package models sample.yaml
//go:generate oapi-codegen -generate server -o server.gen.go -package main -import-mapping=sample.yaml:github.com/yevhenii-babich/go-cources/oapigen/service/models sample.yaml
//go:generate
