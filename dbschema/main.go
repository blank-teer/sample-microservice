package dbschema

// go get -u github.com/jteeuwen/go-bindata/...

//go:generate go-bindata -ignore .+\.go$ -pkg dbschema -o bindata.go ./...
//go:generate gofmt -w bindata.go
