module runtime-go-ping

go 1.25

require github.com/f4ah6o/direct-go-sdk/daab-go v0.0.0

require (
	github.com/f4ah6o/direct-go-sdk/direct-go v0.0.0-20251215153455-bd7b7d48fdd6
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
)

replace github.com/f4ah6o/direct-go-sdk/daab-go => ../../../daab-go

replace github.com/f4ah6o/direct-go-sdk/direct-go => ../../../direct-go
