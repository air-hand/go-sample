module main

go 1.17

require local.packages/web v0.0.0

require (
	github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace local.packages/web => ./web

replace local.packages/fundamentals => ./fundamentals
