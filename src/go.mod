module main

go 1.17

require local.packages/web v0.0.0

require (
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace local.packages/web => ./web

replace local.packages/fundamentals => ./fundamentals
