module local.packages/web

go 1.17

require local.packages/types v0.0.0

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi/v5 v5.0.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lestrrat-go/strftime v1.0.5
	github.com/spf13/viper v1.9.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/sqlboiler/v4 v4.6.0
	github.com/volatiletech/strmangle v0.0.1
)

require (
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace local.packages/types => ../types
