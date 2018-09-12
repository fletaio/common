package common

import (
	"crypto/ecdsa"
	"io"
	"math/big"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// PrivateKey TODO
type PrivateKey struct {
	privkey *ecdsa.PrivateKey
	pubkey  PublicKey
}

// NewPrivateKey TODO
func NewPrivateKey() (*PrivateKey, error) {
	privKey, err := ecrypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	ac := &PrivateKey{
		privkey: privKey,
	}
	if err := ac.calcPubkey(); err != nil {
		return nil, err
	}
	return ac, nil
}

// NewPrivateKeyFromBytes TODO
func NewPrivateKeyFromBytes(pk []byte) (*PrivateKey, error) {
	ac := &PrivateKey{
		privkey: &ecdsa.PrivateKey{
			PublicKey: ecdsa.PublicKey{
				Curve: ecrypto.S256(),
			},
			D: new(big.Int),
		},
	}
	ac.privkey.D.SetBytes(pk)
	ac.privkey.PublicKey.X, ac.privkey.PublicKey.Y = ac.privkey.Curve.ScalarBaseMult(ac.privkey.D.Bytes())
	if err := ac.calcPubkey(); err != nil {
		return nil, err
	}
	return ac, nil
}

// PublicKey TODO
func (ac *PrivateKey) calcPubkey() error {
	pk := ecrypto.CompressPubkey(&ac.privkey.PublicKey)
	copy(ac.pubkey[:], pk[:])
	return nil
}

// PublicKey TODO
func (ac *PrivateKey) PublicKey() PublicKey {
	return ac.pubkey
}

// Sign TODO
func (ac *PrivateKey) Sign(h hash.Hash256) (Signature, error) {
	bs, err := ecrypto.Sign(h[:], ac.privkey)
	if err != nil {
		return Signature{}, err
	}
	var sig Signature
	copy(sig[:], bs)
	return sig, nil
}

// Verify TODO
func (ac *PrivateKey) Verify(h hash.Hash256, sig Signature) bool {
	return ecrypto.VerifySignature(ac.pubkey[:], h[:], sig[:])
}

// WriteTo TODO
func (ac *PrivateKey) WriteTo(w io.Writer) (int64, error) {
	var wrote int64
	bs := ac.privkey.D.Bytes()
	if n, err := util.WriteUint16(w, uint16(len(bs))); err != nil {
		return wrote, err
	} else {
		wrote += n
	}
	if n, err := w.Write(bs); err != nil {
		return wrote, err
	} else {
		wrote += int64(n)
	}
	return wrote, nil
}

// ReadFrom TODO
func (ac *PrivateKey) ReadFrom(r io.Reader) (int64, error) {
	var read int64
	if Len, n, err := util.ReadUint16(r); err != nil {
		return read, err
	} else {
		read += n
		bs := make([]byte, Len)
		if n, err := r.Read(bs); err != nil {
			return read, err
		} else {
			read += int64(n)
			ac.privkey = &ecdsa.PrivateKey{
				PublicKey: ecdsa.PublicKey{
					Curve: ecrypto.S256(),
				},
				D: new(big.Int),
			}
			ac.privkey.D.SetBytes(bs)
			ac.privkey.PublicKey.X, ac.privkey.PublicKey.Y = ac.privkey.Curve.ScalarBaseMult(ac.privkey.D.Bytes())
			if err := ac.calcPubkey(); err != nil {
				return read, err
			}
		}
	}
	return read, nil
}
