package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	ecrypto "git.fleta.io/fleta/common/crypto"
	"git.fleta.io/fleta/common/hash"
)

// RecoverPubkey recover the public key using the hash value and the signature
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

// VerifySignature checks the signature with the public key and the hash value
func VerifySignature(pubkey PublicKey, h hash.Hash256, sig Signature) error {
	if !ecrypto.VerifySignature(pubkey[:], h[:], sig[:64]) {
		return ErrInvalidSignature
	}
	return nil
}
