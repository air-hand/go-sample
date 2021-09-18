module local.packages/web

go 1.17

require local.packages/types v0.0.0

require (
	github.com/go-chi/chi/v5 v5.0.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lestrrat-go/strftime v1.0.5
	gorm.io/gorm v1.21.15
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace local.packages/types => ../types
