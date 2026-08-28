// Separate module so the zip-packing helper's dependency on golang.org/x/mod
// never enters golang.org/x/net's own go.mod.
module sealsecurity.local/modzip

go 1.24.0
