module main

go 1.17

require local.packages/web v0.0.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/chi/v5 v5.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/gorm v1.21.15 // indirect
	local.packages/types v0.0.0 // indirect
)

replace local.packages/web => ./web

replace local.packages/types => ./types

replace local.packages/fundamentals => ./fundamentals
