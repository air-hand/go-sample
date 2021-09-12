module main

go 1.17

require local.packages/web v0.0.0

require (
	github.com/go-chi/chi/v5 v5.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace local.packages/web => ./web

replace local.packages/fundamentals => ./fundamentals
