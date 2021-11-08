module github.com/ikuokuo/start-golang/_/start-ent

go 1.17

require (
	entgo.io/ent v0.9.1
	github.com/mattn/go-sqlite3 v1.14.9
)

require (
	github.com/google/uuid v1.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ikuokuo/start-golang/_/start-ent/ent => ./ent
