package singleflight

import "sync"

// 管理不同 key 的请求(call)。
type Group struct {
	lock sync.Mutex
	m map[string]*Call
}

// 代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入。
type Call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

func (g *Group) Do(key string, fn func()(interface{}, error)) (interface{}, error)  {
	g.lock.Lock()
	if g.m == nil {
		g.m = make(map[string]*Call)
	}

	if call, ok := g.m[key]; ok {
		g.lock.Unlock()
		call.wg.Wait()
		return call.val, call.err
	}

	call := new(Call)
	call.wg.Add(1)
	g.m[key] = call
	g.lock.Unlock()
	call.val, call.err = fn()
	call.wg.Done()

	g.lock.Lock()
	delete(g.m, key)
	g.lock.Unlock()
	return call.val, call.err

}