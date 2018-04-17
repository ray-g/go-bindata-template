#!/usr/bin/env bash

echo "Start go-bindata..."
go-bindata -o bindata/bindata.go -pkg bin templates/...
go-bindata -nocompress -o nocomp/bindata.go -pkg nocomp templates/...
go-bindata -nometadata -o nometa/bindata.go -pkg nometa templates/...
go-bindata -nocompress -nometadata -o nocompmeta/bindata.go -pkg nocompmeta templates/...
echo "Done go-bindata..."

echo "Start build..."
go build main.go
echo "Done build..."

./main
