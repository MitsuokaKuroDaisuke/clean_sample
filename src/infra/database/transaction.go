package database

import "context"

type ctxKey string

var TxCtxKey ctxKey = "transaction_context_key"

type Transaction struct {
	SQLHandler
}

func (tr *Transaction) DoInTx(f func(ctx context.Context) error) error {
	tx := tr.db.Begin()

	ctx := context.WithValue(context.Background(), TxCtxKey, tx)

	if err := f(ctx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
