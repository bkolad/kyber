package edwards25519

import (
	"crypto/sha256"
	"hash"
	"io"
	"reflect"

	"github.com/dedis/fixbuf"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/cipher/sha3"
	"github.com/dedis/kyber/group/internal/marshalling"
)

// SuiteEd25519 implements some basic functionalities such as Group, HashFactory
// and CipherFactory.
type SuiteEd25519 struct {
	Curve
}

// Hash return a newly instanciated sha256 hash function
func (s *SuiteEd25519) Hash() hash.Hash {
	return sha256.New()
}

// Cipher returns the SHA3/SHAKE128 Sponge Cipher
func (s *SuiteEd25519) Cipher(key []byte, options ...interface{}) kyber.Cipher {
	return sha3.NewShakeCipher128(key, options...)
}

func (s *SuiteEd25519) Read(r io.Reader, objs ...interface{}) error {
	return fixbuf.Read(r, s, objs...)
}

func (s *SuiteEd25519) Write(w io.Writer, objs ...interface{}) error {
	return fixbuf.Write(w, objs)
}

// New implements the kyber.Encoding interface
func (s *SuiteEd25519) New(t reflect.Type) interface{} {
	return marshalling.GroupNew(s, t)
}

// NewAES128SHA256Ed25519 returns a cipher suite based on AES-128, SHA-256, and
// the Ed25519 curve.
func NewAES128SHA256Ed25519() *SuiteEd25519 {
	suite := new(SuiteEd25519)
	return suite
}
