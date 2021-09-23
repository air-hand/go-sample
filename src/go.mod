module main

go 1.17

require local.packages/web v0.0.0

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-chi/chi/v5 v5.0.4 // indirect
	github.com/go-redis/redis/v8 v8.11.3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null/v8 v8.1.2 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/volatiletech/sqlboiler/v4 v4.6.0 // indirect
	github.com/volatiletech/strmangle v0.0.1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	local.packages/types v0.0.0 // indirect
)

replace local.packages/web => ./web

replace local.packages/types => ./types

replace local.packages/fundamentals => ./fundamentals
