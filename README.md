# withkoa
Money converter NGN, GHS and KSH.

## How to use
`go run app/main.go`

Open another terminal to convert with curl
`curl -d '{"currency":"ghs", "amount":100.0}' -X POST localhost:3000/shillings`
