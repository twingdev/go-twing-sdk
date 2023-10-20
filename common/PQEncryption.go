package common

import (
	"context"
	"crypto/rand"
	"github.com/cloudflare/circl/hpke"
	"github.com/cloudflare/circl/kem"
	"github.com/libp2p/go-libp2p/core/peer"
)

// PQEncryption is the official selected post quantum encryption scheme (public key) by the NIST

type PQEncryption struct {
	suite    hpke.Suite
	pubKey   kem.PublicKey
	privKey  kem.PrivateKey
	receiver *hpke.Receiver
	sealer   hpke.Sealer
	ct       []byte
	enc      []byte
	aad      []byte
}

func NewPQEncryption(ctx context.Context) *PQEncryption {
	kemID := hpke.KEM_P384_HKDF_SHA384
	kdfID := hpke.KDF_HKDF_SHA384
	aeadID := hpke.AEAD_AES256GCM
	suite := hpke.NewSuite(kemID, kdfID, aeadID)
	pqe := &PQEncryption{suite: suite}
	pub, priv, err := kemID.Scheme().GenerateKeyPair()
	pqe.pubKey = pub
	pqe.privKey = priv
	privKeyBytes, _ := priv.MarshalBinary()
	LocalCache.Set(ctx, "pqe_priv_key", privKeyBytes)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	return pqe

}

func (pqe *PQEncryption) PrepareReceiver(id peer.ID) {
	receiver, _ := pqe.suite.NewReceiver(pqe.privKey, []byte(peer.ID.String(id)))
	pqe.receiver = receiver
}

func (pqe *PQEncryption) PrepareSealer(id peer.ID) []byte {
	sender, _ := pqe.suite.NewSender(pqe.pubKey, []byte(peer.ID.String(id)))
	enc, sealer, err := sender.Setup(rand.Reader)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	pqe.sealer = sealer
	pqe.enc = enc
	return enc
}

func (pqe *PQEncryption) Encrypt(msg, aad []byte) []byte {
	pqe.aad = aad
	ct, err := pqe.sealer.Seal(msg, aad)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	pqe.ct = ct
	return ct

}

func (pqe *PQEncryption) Decrypt() []byte {
	opener, err := pqe.receiver.Setup(pqe.enc)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	dec, err := opener.Open(pqe.ct, pqe.aad)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	return dec

}
