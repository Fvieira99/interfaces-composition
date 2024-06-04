package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	// "Inherited" the properties and methods from bytes.Reader Struct
	// Including the function Read() that is needed to fulfill the HashReader interface
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	// Transform bytes into hashed string
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println("Hashed String from the bytes of the payload: ", hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	// 1 Char = 1 Byte
	// Once the reader is read from someone it cannot be read again
	// Thats why hashReader Struct now has a buf property, so it can be hashed.
	// Also Readers are read-only and Buffers can be read and writen
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("String from the bytes of the payload: ", string(b))
	return nil
}

func main() {
	payload := []byte("Hello World")
	fmt.Println("Array of Bytes from payload: ", payload)
	hashAndBroadcast(NewHashReader(payload))
}
