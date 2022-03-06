// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// Compose returns a new Driver which has functional fields composed
// both from t and x.
func (t Driver) Compose(x Driver) (ret Driver) {
	switch {
	case t.OnInit == nil:
		ret.OnInit = x.OnInit
	case x.OnInit == nil:
		ret.OnInit = t.OnInit
	default:
		h1 := t.OnInit
		h2 := x.OnInit
		ret.OnInit = func(i InitStartInfo) func(InitDoneInfo) {
			r1 := h1(i)
			r2 := h2(i)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(i InitDoneInfo) {
					r1(i)
					r2(i)
				}
			}
		}
	}
	switch {
	case t.OnClose == nil:
		ret.OnClose = x.OnClose
	case x.OnClose == nil:
		ret.OnClose = t.OnClose
	default:
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(c CloseStartInfo) func(CloseDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CloseDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnNetRead == nil:
		ret.OnNetRead = x.OnNetRead
	case x.OnNetRead == nil:
		ret.OnNetRead = t.OnNetRead
	default:
		h1 := t.OnNetRead
		h2 := x.OnNetRead
		ret.OnNetRead = func(n NetReadStartInfo) func(NetReadDoneInfo) {
			r1 := h1(n)
			r2 := h2(n)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(n NetReadDoneInfo) {
					r1(n)
					r2(n)
				}
			}
		}
	}
	switch {
	case t.OnNetWrite == nil:
		ret.OnNetWrite = x.OnNetWrite
	case x.OnNetWrite == nil:
		ret.OnNetWrite = t.OnNetWrite
	default:
		h1 := t.OnNetWrite
		h2 := x.OnNetWrite
		ret.OnNetWrite = func(n NetWriteStartInfo) func(NetWriteDoneInfo) {
			r1 := h1(n)
			r2 := h2(n)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(n NetWriteDoneInfo) {
					r1(n)
					r2(n)
				}
			}
		}
	}
	switch {
	case t.OnNetDial == nil:
		ret.OnNetDial = x.OnNetDial
	case x.OnNetDial == nil:
		ret.OnNetDial = t.OnNetDial
	default:
		h1 := t.OnNetDial
		h2 := x.OnNetDial
		ret.OnNetDial = func(n NetDialStartInfo) func(NetDialDoneInfo) {
			r1 := h1(n)
			r2 := h2(n)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(n NetDialDoneInfo) {
					r1(n)
					r2(n)
				}
			}
		}
	}
	switch {
	case t.OnNetClose == nil:
		ret.OnNetClose = x.OnNetClose
	case x.OnNetClose == nil:
		ret.OnNetClose = t.OnNetClose
	default:
		h1 := t.OnNetClose
		h2 := x.OnNetClose
		ret.OnNetClose = func(n NetCloseStartInfo) func(NetCloseDoneInfo) {
			r1 := h1(n)
			r2 := h2(n)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(n NetCloseDoneInfo) {
					r1(n)
					r2(n)
				}
			}
		}
	}
	switch {
	case t.OnResolve == nil:
		ret.OnResolve = x.OnResolve
	case x.OnResolve == nil:
		ret.OnResolve = t.OnResolve
	default:
		h1 := t.OnResolve
		h2 := x.OnResolve
		ret.OnResolve = func(r ResolveStartInfo) func(ResolveDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r ResolveDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnConnStateChange == nil:
		ret.OnConnStateChange = x.OnConnStateChange
	case x.OnConnStateChange == nil:
		ret.OnConnStateChange = t.OnConnStateChange
	default:
		h1 := t.OnConnStateChange
		h2 := x.OnConnStateChange
		ret.OnConnStateChange = func(c ConnStateChangeStartInfo) func(ConnStateChangeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnStateChangeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnInvoke == nil:
		ret.OnConnInvoke = x.OnConnInvoke
	case x.OnConnInvoke == nil:
		ret.OnConnInvoke = t.OnConnInvoke
	default:
		h1 := t.OnConnInvoke
		h2 := x.OnConnInvoke
		ret.OnConnInvoke = func(c ConnInvokeStartInfo) func(ConnInvokeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnInvokeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnNewStream == nil:
		ret.OnConnNewStream = x.OnConnNewStream
	case x.OnConnNewStream == nil:
		ret.OnConnNewStream = t.OnConnNewStream
	default:
		h1 := t.OnConnNewStream
		h2 := x.OnConnNewStream
		ret.OnConnNewStream = func(c ConnNewStreamStartInfo) func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
					r11 := r1(c)
					r21 := r2(c)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(c ConnNewStreamDoneInfo) {
							r11(c)
							r21(c)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnConnTake == nil:
		ret.OnConnTake = x.OnConnTake
	case x.OnConnTake == nil:
		ret.OnConnTake = t.OnConnTake
	default:
		h1 := t.OnConnTake
		h2 := x.OnConnTake
		ret.OnConnTake = func(c ConnTakeStartInfo) func(ConnTakeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnTakeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnUsagesChange == nil:
		ret.OnConnUsagesChange = x.OnConnUsagesChange
	case x.OnConnUsagesChange == nil:
		ret.OnConnUsagesChange = t.OnConnUsagesChange
	default:
		h1 := t.OnConnUsagesChange
		h2 := x.OnConnUsagesChange
		ret.OnConnUsagesChange = func(c ConnUsagesChangeInfo) {
			h1(c)
			h2(c)
		}
	}
	switch {
	case t.OnConnPark == nil:
		ret.OnConnPark = x.OnConnPark
	case x.OnConnPark == nil:
		ret.OnConnPark = t.OnConnPark
	default:
		h1 := t.OnConnPark
		h2 := x.OnConnPark
		ret.OnConnPark = func(c ConnParkStartInfo) func(ConnParkDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnParkDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnClose == nil:
		ret.OnConnClose = x.OnConnClose
	case x.OnConnClose == nil:
		ret.OnConnClose = t.OnConnClose
	default:
		h1 := t.OnConnClose
		h2 := x.OnConnClose
		ret.OnConnClose = func(c ConnCloseStartInfo) func(ConnCloseDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnCloseDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterInit == nil:
		ret.OnClusterInit = x.OnClusterInit
	case x.OnClusterInit == nil:
		ret.OnClusterInit = t.OnClusterInit
	default:
		h1 := t.OnClusterInit
		h2 := x.OnClusterInit
		ret.OnClusterInit = func(c ClusterInitStartInfo) func(ClusterInitDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterInitDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterClose == nil:
		ret.OnClusterClose = x.OnClusterClose
	case x.OnClusterClose == nil:
		ret.OnClusterClose = t.OnClusterClose
	default:
		h1 := t.OnClusterClose
		h2 := x.OnClusterClose
		ret.OnClusterClose = func(c ClusterCloseStartInfo) func(ClusterCloseDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterCloseDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterGet == nil:
		ret.OnClusterGet = x.OnClusterGet
	case x.OnClusterGet == nil:
		ret.OnClusterGet = t.OnClusterGet
	default:
		h1 := t.OnClusterGet
		h2 := x.OnClusterGet
		ret.OnClusterGet = func(c ClusterGetStartInfo) func(ClusterGetDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterGetDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterInsert == nil:
		ret.OnClusterInsert = x.OnClusterInsert
	case x.OnClusterInsert == nil:
		ret.OnClusterInsert = t.OnClusterInsert
	default:
		h1 := t.OnClusterInsert
		h2 := x.OnClusterInsert
		ret.OnClusterInsert = func(c ClusterInsertStartInfo) func(ClusterInsertDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterInsertDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterUpdate == nil:
		ret.OnClusterUpdate = x.OnClusterUpdate
	case x.OnClusterUpdate == nil:
		ret.OnClusterUpdate = t.OnClusterUpdate
	default:
		h1 := t.OnClusterUpdate
		h2 := x.OnClusterUpdate
		ret.OnClusterUpdate = func(c ClusterUpdateStartInfo) func(ClusterUpdateDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterUpdateDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterRemove == nil:
		ret.OnClusterRemove = x.OnClusterRemove
	case x.OnClusterRemove == nil:
		ret.OnClusterRemove = t.OnClusterRemove
	default:
		h1 := t.OnClusterRemove
		h2 := x.OnClusterRemove
		ret.OnClusterRemove = func(c ClusterRemoveStartInfo) func(ClusterRemoveDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterRemoveDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnPessimizeNode == nil:
		ret.OnPessimizeNode = x.OnPessimizeNode
	case x.OnPessimizeNode == nil:
		ret.OnPessimizeNode = t.OnPessimizeNode
	default:
		h1 := t.OnPessimizeNode
		h2 := x.OnPessimizeNode
		ret.OnPessimizeNode = func(p PessimizeNodeStartInfo) func(PessimizeNodeDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PessimizeNodeDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnRepeaterWakeUp == nil:
		ret.OnRepeaterWakeUp = x.OnRepeaterWakeUp
	case x.OnRepeaterWakeUp == nil:
		ret.OnRepeaterWakeUp = t.OnRepeaterWakeUp
	default:
		h1 := t.OnRepeaterWakeUp
		h2 := x.OnRepeaterWakeUp
		ret.OnRepeaterWakeUp = func(r RepeaterTickStartInfo) func(RepeaterTickDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RepeaterTickDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnGetCredentials == nil:
		ret.OnGetCredentials = x.OnGetCredentials
	case x.OnGetCredentials == nil:
		ret.OnGetCredentials = t.OnGetCredentials
	default:
		h1 := t.OnGetCredentials
		h2 := x.OnGetCredentials
		ret.OnGetCredentials = func(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
			r1 := h1(g)
			r2 := h2(g)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(g GetCredentialsDoneInfo) {
					r1(g)
					r2(g)
				}
			}
		}
	}
	return ret
}
func (t Driver) onInit(i InitStartInfo) func(InitDoneInfo) {
	fn := t.OnInit
	if fn == nil {
		return func(InitDoneInfo) {
			return
		}
	}
	res := fn(i)
	if res == nil {
		return func(InitDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClose(c1 CloseStartInfo) func(CloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(CloseDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetRead(n NetReadStartInfo) func(NetReadDoneInfo) {
	fn := t.OnNetRead
	if fn == nil {
		return func(NetReadDoneInfo) {
			return
		}
	}
	res := fn(n)
	if res == nil {
		return func(NetReadDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetWrite(n NetWriteStartInfo) func(NetWriteDoneInfo) {
	fn := t.OnNetWrite
	if fn == nil {
		return func(NetWriteDoneInfo) {
			return
		}
	}
	res := fn(n)
	if res == nil {
		return func(NetWriteDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetDial(n NetDialStartInfo) func(NetDialDoneInfo) {
	fn := t.OnNetDial
	if fn == nil {
		return func(NetDialDoneInfo) {
			return
		}
	}
	res := fn(n)
	if res == nil {
		return func(NetDialDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetClose(n NetCloseStartInfo) func(NetCloseDoneInfo) {
	fn := t.OnNetClose
	if fn == nil {
		return func(NetCloseDoneInfo) {
			return
		}
	}
	res := fn(n)
	if res == nil {
		return func(NetCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onResolve(r ResolveStartInfo) func(ResolveDoneInfo) {
	fn := t.OnResolve
	if fn == nil {
		return func(ResolveDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(ResolveDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnStateChange(c1 ConnStateChangeStartInfo) func(ConnStateChangeDoneInfo) {
	fn := t.OnConnStateChange
	if fn == nil {
		return func(ConnStateChangeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnStateChangeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnInvoke(c1 ConnInvokeStartInfo) func(ConnInvokeDoneInfo) {
	fn := t.OnConnInvoke
	if fn == nil {
		return func(ConnInvokeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnInvokeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnNewStream(c1 ConnNewStreamStartInfo) func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
	fn := t.OnConnNewStream
	if fn == nil {
		return func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
	}
	return func(c ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
		res := res(c)
		if res == nil {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Driver) onConnTake(c1 ConnTakeStartInfo) func(ConnTakeDoneInfo) {
	fn := t.OnConnTake
	if fn == nil {
		return func(ConnTakeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnTakeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnUsagesChange(c1 ConnUsagesChangeInfo) {
	fn := t.OnConnUsagesChange
	if fn == nil {
		return
	}
	fn(c1)
}
func (t Driver) onConnPark(c1 ConnParkStartInfo) func(ConnParkDoneInfo) {
	fn := t.OnConnPark
	if fn == nil {
		return func(ConnParkDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnParkDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnClose(c1 ConnCloseStartInfo) func(ConnCloseDoneInfo) {
	fn := t.OnConnClose
	if fn == nil {
		return func(ConnCloseDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterInit(c1 ClusterInitStartInfo) func(ClusterInitDoneInfo) {
	fn := t.OnClusterInit
	if fn == nil {
		return func(ClusterInitDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterInitDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterClose(c1 ClusterCloseStartInfo) func(ClusterCloseDoneInfo) {
	fn := t.OnClusterClose
	if fn == nil {
		return func(ClusterCloseDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterGet(c1 ClusterGetStartInfo) func(ClusterGetDoneInfo) {
	fn := t.OnClusterGet
	if fn == nil {
		return func(ClusterGetDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterGetDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterInsert(c1 ClusterInsertStartInfo) func(ClusterInsertDoneInfo) {
	fn := t.OnClusterInsert
	if fn == nil {
		return func(ClusterInsertDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterInsertDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterUpdate(c1 ClusterUpdateStartInfo) func(ClusterUpdateDoneInfo) {
	fn := t.OnClusterUpdate
	if fn == nil {
		return func(ClusterUpdateDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterUpdateDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterRemove(c1 ClusterRemoveStartInfo) func(ClusterRemoveDoneInfo) {
	fn := t.OnClusterRemove
	if fn == nil {
		return func(ClusterRemoveDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterRemoveDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onPessimizeNode(p PessimizeNodeStartInfo) func(PessimizeNodeDoneInfo) {
	fn := t.OnPessimizeNode
	if fn == nil {
		return func(PessimizeNodeDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PessimizeNodeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onRepeaterWakeUp(r RepeaterTickStartInfo) func(RepeaterTickDoneInfo) {
	fn := t.OnRepeaterWakeUp
	if fn == nil {
		return func(RepeaterTickDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RepeaterTickDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onGetCredentials(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
	fn := t.OnGetCredentials
	if fn == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	res := fn(g)
	if res == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	return res
}
func DriverOnInit(t Driver, c *context.Context, endpoint string, database string, secure bool) func(error) {
	var p InitStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Database = database
	p.Secure = secure
	res := t.onInit(p)
	return func(e error) {
		var p InitDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnClose(t Driver, c *context.Context) func(error) {
	var p CloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(e error) {
		var p CloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnNetRead(t Driver, address string, buffer int) func(received int, _ error) {
	var p NetReadStartInfo
	p.Address = address
	p.Buffer = buffer
	res := t.onNetRead(p)
	return func(received int, e error) {
		var p NetReadDoneInfo
		p.Received = received
		p.Error = e
		res(p)
	}
}
func DriverOnNetWrite(t Driver, address string, bytes int) func(sent int, _ error) {
	var p NetWriteStartInfo
	p.Address = address
	p.Bytes = bytes
	res := t.onNetWrite(p)
	return func(sent int, e error) {
		var p NetWriteDoneInfo
		p.Sent = sent
		p.Error = e
		res(p)
	}
}
func DriverOnNetDial(t Driver, c *context.Context, address string) func(error) {
	var p NetDialStartInfo
	p.Context = c
	p.Address = address
	res := t.onNetDial(p)
	return func(e error) {
		var p NetDialDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnNetClose(t Driver, address string) func(error) {
	var p NetCloseStartInfo
	p.Address = address
	res := t.onNetClose(p)
	return func(e error) {
		var p NetCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnResolve(t Driver, target string, resolved []string) func(error) {
	var p ResolveStartInfo
	p.Target = target
	p.Resolved = resolved
	res := t.onResolve(p)
	return func(e error) {
		var p ResolveDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnStateChange(t Driver, endpoint EndpointInfo, state ConnState) func(state ConnState) {
	var p ConnStateChangeStartInfo
	p.Endpoint = endpoint
	p.State = state
	res := t.onConnStateChange(p)
	return func(state ConnState) {
		var p ConnStateChangeDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnInvoke(t Driver, c *context.Context, endpoint EndpointInfo, m Method) func(_ error, issues []Issue, opID string, state ConnState) {
	var p ConnInvokeStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Method = m
	res := t.onConnInvoke(p)
	return func(e error, issues []Issue, opID string, state ConnState) {
		var p ConnInvokeDoneInfo
		p.Error = e
		p.Issues = issues
		p.OpID = opID
		p.State = state
		res(p)
	}
}
func DriverOnConnNewStream(t Driver, c *context.Context, endpoint EndpointInfo, m Method) func(error) func(state ConnState, _ error) {
	var p ConnNewStreamStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Method = m
	res := t.onConnNewStream(p)
	return func(e error) func(ConnState, error) {
		var p ConnNewStreamRecvInfo
		p.Error = e
		res := res(p)
		return func(state ConnState, e error) {
			var p ConnNewStreamDoneInfo
			p.State = state
			p.Error = e
			res(p)
		}
	}
}
func DriverOnConnTake(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p ConnTakeStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnTake(p)
	return func(e error) {
		var p ConnTakeDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnUsagesChange(t Driver, endpoint EndpointInfo, usages int) {
	var p ConnUsagesChangeInfo
	p.Endpoint = endpoint
	p.Usages = usages
	t.onConnUsagesChange(p)
}
func DriverOnConnPark(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p ConnParkStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnPark(p)
	return func(e error) {
		var p ConnParkDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnClose(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p ConnCloseStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnClose(p)
	return func(e error) {
		var p ConnCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnClusterInit(t Driver, c *context.Context) func() {
	var p ClusterInitStartInfo
	p.Context = c
	res := t.onClusterInit(p)
	return func() {
		var p ClusterInitDoneInfo
		res(p)
	}
}
func DriverOnClusterClose(t Driver, c *context.Context) func(error) {
	var p ClusterCloseStartInfo
	p.Context = c
	res := t.onClusterClose(p)
	return func(e error) {
		var p ClusterCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnClusterGet(t Driver, c *context.Context) func(endpoint EndpointInfo, _ error) {
	var p ClusterGetStartInfo
	p.Context = c
	res := t.onClusterGet(p)
	return func(endpoint EndpointInfo, e error) {
		var p ClusterGetDoneInfo
		p.Endpoint = endpoint
		p.Error = e
		res(p)
	}
}
func DriverOnClusterInsert(t Driver, c *context.Context, endpoint EndpointInfo) func(inserted bool, state ConnState) {
	var p ClusterInsertStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onClusterInsert(p)
	return func(inserted bool, state ConnState) {
		var p ClusterInsertDoneInfo
		p.Inserted = inserted
		p.State = state
		res(p)
	}
}
func DriverOnClusterUpdate(t Driver, c *context.Context, endpoint EndpointInfo) func(state ConnState) {
	var p ClusterUpdateStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onClusterUpdate(p)
	return func(state ConnState) {
		var p ClusterUpdateDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnClusterRemove(t Driver, c *context.Context, endpoint EndpointInfo) func(state ConnState, removed bool) {
	var p ClusterRemoveStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onClusterRemove(p)
	return func(state ConnState, removed bool) {
		var p ClusterRemoveDoneInfo
		p.State = state
		p.Removed = removed
		res(p)
	}
}
func DriverOnPessimizeNode(t Driver, c *context.Context, endpoint EndpointInfo, state ConnState, cause error) func(state ConnState) {
	var p PessimizeNodeStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.State = state
	p.Cause = cause
	res := t.onPessimizeNode(p)
	return func(state ConnState) {
		var p PessimizeNodeDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnRepeaterWakeUp(t Driver, c *context.Context, name string, event string) func(error) {
	var p RepeaterTickStartInfo
	p.Context = c
	p.Name = name
	p.Event = event
	res := t.onRepeaterWakeUp(p)
	return func(e error) {
		var p RepeaterTickDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnGetCredentials(t Driver, c *context.Context) func(token string, _ error) {
	var p GetCredentialsStartInfo
	p.Context = c
	res := t.onGetCredentials(p)
	return func(token string, e error) {
		var p GetCredentialsDoneInfo
		p.Token = token
		p.Error = e
		res(p)
	}
}
