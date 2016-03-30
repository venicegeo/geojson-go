# geojson-go

## Overview

A simple GeoJSON serializer / deserializer, designed to work with all GeoJSON objects.

Use it as a library to get back an object that can be cast to a GeoJSON object:

```
func Parse(bytes []byte) (interface{}, error)
```

Use it from the command line:
```
go build
./geojson-go [file1] [file2]...
```

When used from the command line, it attempts to read the input file, write out Go's internal representation (`%#v`), and round-trip it back to GeoJSON.



