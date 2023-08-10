// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
	"time"
)

// databaseSQLComposeOptions is a holder of options
type databaseSQLComposeOptions struct {
	panicCallback func(e interface{})
}

// DatabaseSQLOption specified DatabaseSQL compose option
type DatabaseSQLComposeOption func(o *databaseSQLComposeOptions)

// WithDatabaseSQLPanicCallback specified behavior on panic
func WithDatabaseSQLPanicCallback(cb func(e interface{})) DatabaseSQLComposeOption {
	return func(o *databaseSQLComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new DatabaseSQL which has functional fields composed both from t and x.
func (t *DatabaseSQL) Compose(x *DatabaseSQL, opts ...DatabaseSQLComposeOption) *DatabaseSQL {
	var ret DatabaseSQL
	options := databaseSQLComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnConnectorConnect
		h2 := x.OnConnectorConnect
		ret.OnConnectorConnect = func(d DatabaseSQLConnectorConnectStartInfo) func(DatabaseSQLConnectorConnectDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnectorConnectDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnectorConnectDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnPing
		h2 := x.OnConnPing
		ret.OnConnPing = func(d DatabaseSQLConnPingStartInfo) func(DatabaseSQLConnPingDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnPingDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnPingDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnPrepare
		h2 := x.OnConnPrepare
		ret.OnConnPrepare = func(d DatabaseSQLConnPrepareStartInfo) func(DatabaseSQLConnPrepareDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnPrepareDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnPrepareDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnClose
		h2 := x.OnConnClose
		ret.OnConnClose = func(d DatabaseSQLConnCloseStartInfo) func(DatabaseSQLConnCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnBegin
		h2 := x.OnConnBegin
		ret.OnConnBegin = func(d DatabaseSQLConnBeginStartInfo) func(DatabaseSQLConnBeginDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnBeginDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnBeginDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnQuery
		h2 := x.OnConnQuery
		ret.OnConnQuery = func(d DatabaseSQLConnQueryStartInfo) func(DatabaseSQLConnQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnQueryDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnExec
		h2 := x.OnConnExec
		ret.OnConnExec = func(d DatabaseSQLConnExecStartInfo) func(DatabaseSQLConnExecDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLConnExecDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLConnExecDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnTxQuery
		h2 := x.OnTxQuery
		ret.OnTxQuery = func(d DatabaseSQLTxQueryStartInfo) func(DatabaseSQLTxQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLTxQueryDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLTxQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnTxExec
		h2 := x.OnTxExec
		ret.OnTxExec = func(d DatabaseSQLTxExecStartInfo) func(DatabaseSQLTxExecDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLTxExecDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLTxExecDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnTxCommit
		h2 := x.OnTxCommit
		ret.OnTxCommit = func(d DatabaseSQLTxCommitStartInfo) func(DatabaseSQLTxCommitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLTxCommitDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLTxCommitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnTxRollback
		h2 := x.OnTxRollback
		ret.OnTxRollback = func(d DatabaseSQLTxRollbackStartInfo) func(DatabaseSQLTxRollbackDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLTxRollbackDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLTxRollbackDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnStmtQuery
		h2 := x.OnStmtQuery
		ret.OnStmtQuery = func(d DatabaseSQLStmtQueryStartInfo) func(DatabaseSQLStmtQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLStmtQueryDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLStmtQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnStmtExec
		h2 := x.OnStmtExec
		ret.OnStmtExec = func(d DatabaseSQLStmtExecStartInfo) func(DatabaseSQLStmtExecDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLStmtExecDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLStmtExecDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnStmtClose
		h2 := x.OnStmtClose
		ret.OnStmtClose = func(d DatabaseSQLStmtCloseStartInfo) func(DatabaseSQLStmtCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLStmtCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLStmtCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnDoTx
		h2 := x.OnDoTx
		ret.OnDoTx = func(d DatabaseSQLDoTxStartInfo) func(DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(DatabaseSQLDoTxDoneInfo)
				if r != nil {
					r2 = r(d)
				}
				if r1 != nil {
					r3 = r1(d)
				}
				return func(d DatabaseSQLDoTxDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(d)
					}
					if r3 != nil {
						r3(d)
					}
				}
			}
		}
	}
	return &ret
}

func (t *DatabaseSQL) onConnectorConnect(d DatabaseSQLConnectorConnectStartInfo) func(DatabaseSQLConnectorConnectDoneInfo) {
	fn := t.OnConnectorConnect
	if fn == nil {
		return func(DatabaseSQLConnectorConnectDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnectorConnectDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnPing(d DatabaseSQLConnPingStartInfo) func(DatabaseSQLConnPingDoneInfo) {
	fn := t.OnConnPing
	if fn == nil {
		return func(DatabaseSQLConnPingDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnPingDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnPrepare(d DatabaseSQLConnPrepareStartInfo) func(DatabaseSQLConnPrepareDoneInfo) {
	fn := t.OnConnPrepare
	if fn == nil {
		return func(DatabaseSQLConnPrepareDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnPrepareDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnClose(d DatabaseSQLConnCloseStartInfo) func(DatabaseSQLConnCloseDoneInfo) {
	fn := t.OnConnClose
	if fn == nil {
		return func(DatabaseSQLConnCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnCloseDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnBegin(d DatabaseSQLConnBeginStartInfo) func(DatabaseSQLConnBeginDoneInfo) {
	fn := t.OnConnBegin
	if fn == nil {
		return func(DatabaseSQLConnBeginDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnBeginDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnQuery(d DatabaseSQLConnQueryStartInfo) func(DatabaseSQLConnQueryDoneInfo) {
	fn := t.OnConnQuery
	if fn == nil {
		return func(DatabaseSQLConnQueryDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnQueryDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onConnExec(d DatabaseSQLConnExecStartInfo) func(DatabaseSQLConnExecDoneInfo) {
	fn := t.OnConnExec
	if fn == nil {
		return func(DatabaseSQLConnExecDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLConnExecDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onTxQuery(d DatabaseSQLTxQueryStartInfo) func(DatabaseSQLTxQueryDoneInfo) {
	fn := t.OnTxQuery
	if fn == nil {
		return func(DatabaseSQLTxQueryDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLTxQueryDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onTxExec(d DatabaseSQLTxExecStartInfo) func(DatabaseSQLTxExecDoneInfo) {
	fn := t.OnTxExec
	if fn == nil {
		return func(DatabaseSQLTxExecDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLTxExecDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onTxCommit(d DatabaseSQLTxCommitStartInfo) func(DatabaseSQLTxCommitDoneInfo) {
	fn := t.OnTxCommit
	if fn == nil {
		return func(DatabaseSQLTxCommitDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLTxCommitDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onTxRollback(d DatabaseSQLTxRollbackStartInfo) func(DatabaseSQLTxRollbackDoneInfo) {
	fn := t.OnTxRollback
	if fn == nil {
		return func(DatabaseSQLTxRollbackDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLTxRollbackDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onStmtQuery(d DatabaseSQLStmtQueryStartInfo) func(DatabaseSQLStmtQueryDoneInfo) {
	fn := t.OnStmtQuery
	if fn == nil {
		return func(DatabaseSQLStmtQueryDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLStmtQueryDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onStmtExec(d DatabaseSQLStmtExecStartInfo) func(DatabaseSQLStmtExecDoneInfo) {
	fn := t.OnStmtExec
	if fn == nil {
		return func(DatabaseSQLStmtExecDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLStmtExecDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onStmtClose(d DatabaseSQLStmtCloseStartInfo) func(DatabaseSQLStmtCloseDoneInfo) {
	fn := t.OnStmtClose
	if fn == nil {
		return func(DatabaseSQLStmtCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLStmtCloseDoneInfo) {
			return
		}
	}
	return res
}

func (t *DatabaseSQL) onDoTx(d DatabaseSQLDoTxStartInfo) func(DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
	fn := t.OnDoTx
	if fn == nil {
		return func(DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
			return func(DatabaseSQLDoTxDoneInfo) {
				return
			}
		}
	}
	res := fn(d)
	if res == nil {
		return func(DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
			return func(DatabaseSQLDoTxDoneInfo) {
				return
			}
		}
	}
	return func(d DatabaseSQLDoTxIntermediateInfo) func(DatabaseSQLDoTxDoneInfo) {
		res := res(d)
		if res == nil {
			return func(DatabaseSQLDoTxDoneInfo) {
				return
			}
		}
		return res
	}
}

func DatabaseSQLOnConnectorConnect(t *DatabaseSQL, c *context.Context) func(_ error, session tableSessionInfo) {
	var p DatabaseSQLConnectorConnectStartInfo
	p.Context = c
	res := t.onConnectorConnect(p)
	return func(e error, session tableSessionInfo) {
		var p DatabaseSQLConnectorConnectDoneInfo
		p.Error = e
		p.Session = session
		res(p)
	}
}

func DatabaseSQLOnConnPing(t *DatabaseSQL, c *context.Context) func(error) {
	var p DatabaseSQLConnPingStartInfo
	p.Context = c
	res := t.onConnPing(p)
	return func(e error) {
		var p DatabaseSQLConnPingDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnConnPrepare(t *DatabaseSQL, c *context.Context, query string) func(error) {
	var p DatabaseSQLConnPrepareStartInfo
	p.Context = c
	p.Query = query
	res := t.onConnPrepare(p)
	return func(e error) {
		var p DatabaseSQLConnPrepareDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnConnClose(t *DatabaseSQL) func(error) {
	var p DatabaseSQLConnCloseStartInfo
	res := t.onConnClose(p)
	return func(e error) {
		var p DatabaseSQLConnCloseDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnConnBegin(t *DatabaseSQL, c *context.Context) func(tx tableTransactionInfo, _ error) {
	var p DatabaseSQLConnBeginStartInfo
	p.Context = c
	res := t.onConnBegin(p)
	return func(tx tableTransactionInfo, e error) {
		var p DatabaseSQLConnBeginDoneInfo
		p.Tx = tx
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnConnQuery(t *DatabaseSQL, c *context.Context, query string, mode string, idempotent bool, idleTime time.Duration) func(error) {
	var p DatabaseSQLConnQueryStartInfo
	p.Context = c
	p.Query = query
	p.Mode = mode
	p.Idempotent = idempotent
	p.IdleTime = idleTime
	res := t.onConnQuery(p)
	return func(e error) {
		var p DatabaseSQLConnQueryDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnConnExec(t *DatabaseSQL, c *context.Context, query string, mode string, idempotent bool, idleTime time.Duration) func(error) {
	var p DatabaseSQLConnExecStartInfo
	p.Context = c
	p.Query = query
	p.Mode = mode
	p.Idempotent = idempotent
	p.IdleTime = idleTime
	res := t.onConnExec(p)
	return func(e error) {
		var p DatabaseSQLConnExecDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnTxQuery(t *DatabaseSQL, c *context.Context, txContext context.Context, tx tableTransactionInfo, query string, idempotent bool) func(error) {
	var p DatabaseSQLTxQueryStartInfo
	p.Context = c
	p.TxContext = txContext
	p.Tx = tx
	p.Query = query
	p.Idempotent = idempotent
	res := t.onTxQuery(p)
	return func(e error) {
		var p DatabaseSQLTxQueryDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnTxExec(t *DatabaseSQL, c *context.Context, txContext context.Context, tx tableTransactionInfo, query string, idempotent bool) func(error) {
	var p DatabaseSQLTxExecStartInfo
	p.Context = c
	p.TxContext = txContext
	p.Tx = tx
	p.Query = query
	p.Idempotent = idempotent
	res := t.onTxExec(p)
	return func(e error) {
		var p DatabaseSQLTxExecDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnTxCommit(t *DatabaseSQL, c *context.Context, tx tableTransactionInfo) func(error) {
	var p DatabaseSQLTxCommitStartInfo
	p.Context = c
	p.Tx = tx
	res := t.onTxCommit(p)
	return func(e error) {
		var p DatabaseSQLTxCommitDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnTxRollback(t *DatabaseSQL, c *context.Context, tx tableTransactionInfo) func(error) {
	var p DatabaseSQLTxRollbackStartInfo
	p.Context = c
	p.Tx = tx
	res := t.onTxRollback(p)
	return func(e error) {
		var p DatabaseSQLTxRollbackDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnStmtQuery(t *DatabaseSQL, c *context.Context, query string) func(error) {
	var p DatabaseSQLStmtQueryStartInfo
	p.Context = c
	p.Query = query
	res := t.onStmtQuery(p)
	return func(e error) {
		var p DatabaseSQLStmtQueryDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnStmtExec(t *DatabaseSQL, c *context.Context, query string) func(error) {
	var p DatabaseSQLStmtExecStartInfo
	p.Context = c
	p.Query = query
	res := t.onStmtExec(p)
	return func(e error) {
		var p DatabaseSQLStmtExecDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnStmtClose(t *DatabaseSQL) func(error) {
	var p DatabaseSQLStmtCloseStartInfo
	res := t.onStmtClose(p)
	return func(e error) {
		var p DatabaseSQLStmtCloseDoneInfo
		p.Error = e
		res(p)
	}
}

func DatabaseSQLOnDoTx(t *DatabaseSQL, c *context.Context, iD string, idempotent bool) func(error) func(attempts int, _ error) {
	var p DatabaseSQLDoTxStartInfo
	p.Context = c
	p.ID = iD
	p.Idempotent = idempotent
	res := t.onDoTx(p)
	return func(e error) func(int, error) {
		var p DatabaseSQLDoTxIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p DatabaseSQLDoTxDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
