# Port Scanner in Go

### Usage:

go run main.go --h hostname

#### Other flags:

--p <porst to scan>
- e.g -> go run main.go --h hostname --p 80,120,200

--pc <protocol> [tcp or udp]
- e.g -> go run main.go --h hostname --p ports --pc udp