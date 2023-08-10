// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// discoveryComposeOptions is a holder of options
type discoveryComposeOptions struct {
	panicCallback func(e interface{})
}

// DiscoveryOption specified Discovery compose option
type DiscoveryComposeOption func(o *discoveryComposeOptions)

// WithDiscoveryPanicCallback specified behavior on panic
func WithDiscoveryPanicCallback(cb func(e interface{})) DiscoveryComposeOption {
	return func(o *discoveryComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Discovery which has functional fields composed both from t and x.
func (t *Discovery) Compose(x *Discovery, opts ...DiscoveryComposeOption) *Discovery {
	var ret Discovery
	options := discoveryComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnDiscover
		h2 := x.OnDiscover
		ret.OnDiscover = func(d DiscoveryDiscoverStartInfo) func(DiscoveryDiscoverDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DiscoveryDiscoverDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DiscoveryDiscoverDoneInfo) {
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
		h1 := t.OnWhoAmI
		h2 := x.OnWhoAmI
		ret.OnWhoAmI = func(d DiscoveryWhoAmIStartInfo) func(DiscoveryWhoAmIDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DiscoveryWhoAmIDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DiscoveryWhoAmIDoneInfo) {
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
	return &ret
}

func (t *Discovery) onDiscover(d DiscoveryDiscoverStartInfo) func(DiscoveryDiscoverDoneInfo) {
	fn := t.OnDiscover
	if fn == nil {
		return func(DiscoveryDiscoverDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DiscoveryDiscoverDoneInfo) {
			return
		}
	}
	return res
}

func (t *Discovery) onWhoAmI(d DiscoveryWhoAmIStartInfo) func(DiscoveryWhoAmIDoneInfo) {
	fn := t.OnWhoAmI
	if fn == nil {
		return func(DiscoveryWhoAmIDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DiscoveryWhoAmIDoneInfo) {
			return
		}
	}
	return res
}

func DiscoveryOnDiscover(t *Discovery, c *context.Context, address string, database string) func(location string, endpoints []EndpointInfo, _ error) {
	var p DiscoveryDiscoverStartInfo
	p.Context = c
	p.Address = address
	p.Database = database
	res := t.onDiscover(p)
	return func(location string, endpoints []EndpointInfo, e error) {
		var p DiscoveryDiscoverDoneInfo
		p.Location = location
		p.Endpoints = endpoints
		p.Error = e
		res(p)
	}
}

func DiscoveryOnWhoAmI(t *Discovery, c *context.Context) func(user string, groups []string, _ error) {
	var p DiscoveryWhoAmIStartInfo
	p.Context = c
	res := t.onWhoAmI(p)
	return func(user string, groups []string, e error) {
		var p DiscoveryWhoAmIDoneInfo
		p.User = user
		p.Groups = groups
		p.Error = e
		res(p)
	}
}
