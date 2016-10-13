# geojson-go

## Overview

A simple GeoJSON serializer / deserializer, designed to work with all GeoJSON objects. Based on [draft-ietf-geojson](https://datatracker.ietf.org/doc/draft-ietf-geojson/).

Use it as a library to get back an object that can be cast to a GeoJSON object (FeatureCollection, Feature, or any of the 7 geometry types):

```
func Parse(bytes []byte) (interface{}, error)
```

When used from the command line, it attempts to read the input file, write out Go's internal representation (`%#v`), and round-trip it back to GeoJSON.

### Why Yet Another?
There are other GeoJSON Go libraries out there. This library handles any GeoJSON object as input and returns an `interface{}`. You don't need to know what it is beforehand. This is useful when your input can come from a variety of places. (If you know enough about the input to call a function called, let's say, "ParseFeatureCollection", you know enough about the output to call `output.(geojson.FeatureCollection)`.) 

### What does it not do well?
1. It is not designed to validate input. The library assumes that the input is valid GeoJSON. Behavior for invalid GeoJSON is undefined and could panic.
2. Foreign members are not supported. If your GeoJSON input has predictable foreign members, the best thing to do is to unmarshal them into a separate struct from the GeoJSON object.
