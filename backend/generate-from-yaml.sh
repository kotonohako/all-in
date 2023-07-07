#!/bin/bash

oapi-codegen -generate types -package generated ../api.yaml > presentation/generated/types.gen.go
oapi-codegen -generate server -package generated ../api.yaml > presentation/generated/api.gen.go

