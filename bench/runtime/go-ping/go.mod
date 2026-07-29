module runtime-go-ping

go 1.25

require (
	github.com/f4ah6o/direct-go-sdk/daab-go v0.0.0
	github.com/f4ah6o/direct-go-sdk/direct-go v0.0.0
)

require gopkg.in/yaml.v3 v3.0.1

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	go.opentelemetry.io/otel/metric v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.38.0 // indirect
)

replace github.com/f4ah6o/direct-go-sdk/daab-go => ../../../daab-go

replace github.com/f4ah6o/direct-go-sdk/direct-go => ../../../direct-go
