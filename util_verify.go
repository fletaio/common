package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	"git.fleta.io/fleta/common/hash"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
)

// RecoverPubkey TODO
func RecoverPubkey(h hash.Hash256, sig Signature) (PublicKey, error) {
	bs, err := ecrypto.Ecrecover(h[:], sig[:])
	if err != nil {
		return PublicKey{}, err
	}
	X, Y := elliptic.Unmarshal(ecrypto.S256(), bs)
	key := ecrypto.CompressPubkey(&ecdsa.PublicKey{
		Curve: ecrypto.S256(),
		X:     X,
		Y:     Y,
	})
	var pubkey PublicKey
	copy(pubkey[:], key)
	return pubkey, nil
}

// VerifySignature TODO
func VerifySignature(pubkey PublicKey, h []byte, sig Signature) error {
	if !ecrypto.VerifySignature(pubkey[:], h, sig[:64]) {
		return ErrInvalidSignature
	}
	return nil
}
