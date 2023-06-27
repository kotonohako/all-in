#!/bin/bash

oapi-codegen -generate types -package generated ../api.yaml > presentation/generated/types/types.gen.go
oapi-codegen -generate server -package generated ../api.yaml > presentation/generated/api/api.gen.go

