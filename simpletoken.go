package simpletoken

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"reflect"

	"github.com/vmihailenco/msgpack"
)

var (
	algorithms = map[string]func() hash.Hash{
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}
)

type SimpleToken struct {
	hashFunc  func() hash.Hash
	algorithm string
	secret    []byte
}

type Token struct {
	bytes []byte
}

func (t *Token) String() string {
	return hex.EncodeToString(t.bytes)
}

func (t *Token) Bytes() []byte {
	return t.bytes
}

func New(algorithm string, secret []byte) (*SimpleToken, error) {
	hashFunc, ok := algorithms[algorithm]
	if !ok {
		return nil, fmt.Errorf("invalid hash algorithm %s. selected algorithm does not supported by simpletoken", algorithm)
	}
	return &SimpleToken{
		hashFunc:  hashFunc,
		algorithm: algorithm,
		secret:    secret,
	}, nil
}

func (s *SimpleToken) getHMAC(bytes []byte) []byte {
	h := hmac.New(s.hashFunc, s.secret)
	h.Write(bytes)
	return h.Sum(nil)
}

func (s *SimpleToken) Generate(payload interface{}) (*Token, error) {
	bytes, err := msgpack.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return &Token{
		bytes: append(bytes, s.getHMAC(bytes)...),
	}, nil
}

func (s *SimpleToken) ParseString(token string, output interface{}) error {
	bytes, err := hex.DecodeString(token)
	if err != nil {
		return err
	}
	return s.Parse(bytes, output)
}

func (s *SimpleToken) Parse(token []byte, output interface{}) error {
	hmacSize := s.hashFunc().Size()
	length := len(token)
	if length < hmacSize {
		return fmt.Errorf("invalid token size. length must be greater than %d bytes", hmacSize)
	}
	index := length - hmacSize
	payloadBytes := token[:index]
	hmacBytes := token[index:]
	if !reflect.DeepEqual(hmacBytes, s.getHMAC(payloadBytes)) {
		return fmt.Errorf("token is not validated by hmac")
	}
	return msgpack.Unmarshal(payloadBytes, output)
}
