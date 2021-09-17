package cache_replacement_golang

import "container/list"

func FILO() func(int) Cache {
	return Build(Spec{
		Wrap: func(key string, value interface{}) Node {
			return &NodeImpl{key: key, value: value}
		},
		Evict: func(ctx *Ctx) {
			ctx.List.Remove(ctx.List.Front())
		},
		Found: func(ctx *Ctx, item *list.Element) {

		},
	})
}
