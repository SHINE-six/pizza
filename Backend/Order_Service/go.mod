module Order_Service

go 1.22.4

require (
	github.com/99designs/gqlgen v0.17.49
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/spf13/viper v1.19.0
	github.com/vektah/gqlparser/v2 v2.5.16
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.11
	pizza/db v0.0.0-00010101000000-000000000000
)

replace pizza/db => git-codecommit.ap-southeast-1.amazonaws.com/v1/repos/Database_Struct.git v0.0.5

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)