package validator

import (
	"runtime"
	"sync"

	"git.fleta.io/fleta/common"
	"git.fleta.io/fleta/common/hash"
)

// Validator TODO
type Validator struct {
	requestChan chan *request
	aliveWorker int64
	errorCount  int64
	requestEnd  int64
}

// NewValidator TODO
func NewValidator() *Validator {
	vr := &Validator{
		requestChan: make(chan *request, runtime.NumCPU()*5000),
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		go vr.worker()
	}
	return vr
}

// Request TODO
func (vr *Validator) Request(ID int, h hash.Hash256, sig common.Signature, addr common.Address) *Response {
	req := requestPool.Get().(*request)
	req.ID = ID
	req.Hash = h
	req.Signature = sig
	req.Address = addr
	vr.requestChan <- req
	return req.Response
}

func (vr *Validator) worker() {
	for req := range vr.requestChan {
		pubkey, _ := common.RecoverPubkey(req.Hash, req.Signature)
		addr := common.AddressFromPubkey(pubkey)
		if !addr.Equal(req.Address) {
			req.Response.C <- ErrInvalidPublicKey
		} else {
			req.Response.C <- nil
		}
		requestPool.Put(req)
	}
}

type request struct {
	ID        int
	Hash      hash.Hash256
	Signature common.Signature
	Address   common.Address
	Response  *Response
}

// Response TODO
type Response struct {
	C chan error
}

var requestPool = sync.Pool{
	New: func() interface{} {
		return &request{
			Response: &Response{
				C: make(chan error),
			},
		}
	},
}
