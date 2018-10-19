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
func (vr *Validator) Request(h hash.Hash256, sigs []common.Signature) []common.PublicHash {
	req := requestPool.Get().(*request)
	req.Hash = h
	req.Signatures = sigs
	vr.requestChan <- req
	return <-req.Response.C
}

func (vr *Validator) worker() {
	for req := range vr.requestChan {
		PublicHashes := make([]common.PublicHash, 0, len(req.Signatures))
		for _, sig := range req.Signatures {
			pubkey, _ := common.RecoverPubkey(req.Hash, sig)
			pubhash := common.NewPublicHash(pubkey)
			PublicHashes = append(PublicHashes, pubhash)
		}
		req.Response.C <- PublicHashes
		requestPool.Put(req)
	}
}

type request struct {
	Hash       hash.Hash256
	Signatures []common.Signature
	Response   *Response
}

// Response TODO
type Response struct {
	C chan []common.PublicHash
}

var requestPool = sync.Pool{
	New: func() interface{} {
		return &request{
			Response: &Response{
				C: make(chan []common.PublicHash),
			},
		}
	},
}
