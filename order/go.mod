module order

go 1.13

require (
	github.com/asim/go-micro/plugins/config/source/consul/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.7.0 // indirect
	github.com/asim/go-micro/v3 v3.7.1
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	google.golang.org/protobuf v1.26.0
)
require common v0.0.0
replace common => ../common