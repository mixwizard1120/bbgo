package v3

//go:generate -command GetRequest requestgen -method GET
//go:generate -command PostRequest requestgen -method POST
//go:generate -command DeleteRequest requestgen -method DELETE

import (
	maxapi "github.com/c9s/bbgo/pkg/exchange/max/maxapi"
)

// create type alias
type WalletType = maxapi.WalletType
type Order = maxapi.Order
type Trade = maxapi.Trade
type Account = maxapi.Account

// OrderService manages the Order endpoint.
type OrderService struct {
	Client *maxapi.RestClient
}
