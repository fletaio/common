package validator

import (
	"runtime"
	"sync"
	"sync/atomic"

	"git.fleta.io/fleta/common"
	"git.fleta.io/fleta/common/hash"
)

// BatchValidator TODO
type BatchValidator struct {
	batchRequestChan  chan *batchRequest
	batchResponseChan chan *batchResponse
	batchRequestCount int
	aliveWorker       int64
	errorCount        int64
	batchRequestEnd   int64
}

// NewBatchValidator TODO
func NewBatchValidator() *BatchValidator {
	vr := &BatchValidator{
		batchRequestChan:  make(chan *batchRequest, runtime.NumCPU()*5000),
		batchResponseChan: make(chan *batchResponse, runtime.NumCPU()*5000),
		aliveWorker:       int64(runtime.NumCPU()),
	}
	for i := int64(0); i < vr.aliveWorker; i++ {
		go vr.worker()
	}
	return vr
}

// DoRequest TODO
func (vr *BatchValidator) DoRequest(fn func()) {
	go func() {
		defer func() {
			recover()
		}()
		fn()
		atomic.AddInt64(&vr.batchRequestEnd, 1)
	}()
}

// Request TODO
func (vr *BatchValidator) Request(ID int, h hash.Hash256, sig common.Signature, addr common.Address) {
	req := batchRequestPool.Get().(*batchRequest)
	req.ID = ID
	req.Hash = h
	req.Signature = sig
	req.Address = addr
	vr.batchRequestChan <- req
	vr.batchRequestCount++
}

// WaitResponse TODO
func (vr *BatchValidator) WaitResponse() (ID int, err error) {
	Count := 0
	for res := range vr.batchResponseChan {
		ResErr := res.Error
		batchResponsePool.Put(res)
		if ResErr != nil {
			atomic.AddInt64(&vr.errorCount, 1)
			close(vr.batchRequestChan)
			ID = res.ID
			err = ErrInvalidPublicKey
		}
		Count++
		if Count >= vr.batchRequestCount && vr.batchRequestEnd > 0 {
			close(vr.batchRequestChan)
			break
		}
	}
	return
}

func (vr *BatchValidator) worker() {
	for req := range vr.batchRequestChan {
		if vr.errorCount == 0 {
			pubkey, _ := common.RecoverPubkey(req.Hash, req.Signature)
			addr := common.AddressFromPubkey(pubkey)
			res := batchResponsePool.Get().(*batchResponse)
			res.ID = req.ID
			if !addr.Equal(req.Address) {
				res.Error = ErrInvalidPublicKey
			} else {
				res.Error = nil
			}
			vr.batchResponseChan <- res
		}
		batchRequestPool.Put(req)
	}
	if atomic.AddInt64(&vr.aliveWorker, -1) == 0 {
		close(vr.batchResponseChan)
	}
}

type batchRequest struct {
	ID        int
	Hash      hash.Hash256
	Signature common.Signature
	Address   common.Address
}

type batchResponse struct {
	ID    int
	Error error
}

var batchRequestPool = sync.Pool{
	New: func() interface{} {
		return &batchRequest{}
	},
}

var batchResponsePool = sync.Pool{
	New: func() interface{} {
		return &batchResponse{}
	},
}
