package domain

import "context"

type ctxKey string

var TxCtxKey ctxKey = "transaction_context_key"

type Transaction interface {
	DoInTx(f func(ctx context.Context) error) error
}
