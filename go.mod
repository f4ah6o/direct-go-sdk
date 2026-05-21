module github.com/f4ah6o/direct-go-sdk/direct-teams-bridge

go 1.25

require (
	github.com/f4ah6o/direct-go-sdk/direct-go v0.0.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	golang.org/x/term v0.13.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

replace github.com/f4ah6o/direct-go-sdk/direct-go => ./direct-go
