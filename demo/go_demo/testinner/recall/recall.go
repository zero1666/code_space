package recall

import "learngo/testinner/rctx"

type Recall struct {
	*rctx.RecContext
}

func (r *Recall) Init(rc *rctx.RecContext) {
	r.RecContext = rc
}
