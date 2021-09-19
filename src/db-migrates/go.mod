module local.packages/db-migrates

go 1.17

replace local.packages/web => ../web

replace local.packages/types => ../types

require (
	github.com/golang-migrate/migrate/v4 v4.14.1
	local.packages/web v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-chi/chi/v5 v5.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	gorm.io/gorm v1.21.15 // indirect
	local.packages/types v0.0.0 // indirect
)
