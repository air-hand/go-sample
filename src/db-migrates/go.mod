module local.packages/db-migrates

go 1.17

replace local.packages/web => ../web

replace local.packages/types => ../types

require (
	github.com/golang-migrate/migrate/v4 v4.14.1
	local.packages/types v0.0.0
	local.packages/web v0.0.0-00010101000000-000000000000
)

require (
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-chi/chi/v5 v5.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null/v8 v8.1.2 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/volatiletech/sqlboiler/v4 v4.6.0 // indirect
	github.com/volatiletech/strmangle v0.0.1 // indirect
	golang.org/x/net v0.0.0-20210503060351-7fd8e65b6420 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
