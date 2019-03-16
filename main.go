package randomhash

// #include "RandomHash.h"

import "C"
import (
	"math/big"
	"unsafe"
)

type RandomHash struct {
	hashrate int64
}

func (rh *RandomHash) ScanHash(block []byte, stop <-chan struct{}, index int) uint64 {

}

func (rh *RandomHash) ScanHash() {

}

func (rh *RandomHash) GetHashrate() int64 {
	return rh.hashrate
}

func (rh *RandomHash) Verify(block []byte) bool {
	// C.RandomHash_Simple()
}

// Search , aka ScanHash
func (rh *RandomHash) Search(block []byte, index int) uint64 {

}

func (rh *RandomHash) CalcHash(block []byte, nonce uint64) *big.Int {
	var in unsafe.Pointer = C.CBytes(block)
	var out unsafe.Pointer = C.malloc(HashLength)
	C.RandomHash_Simple((*C.uint8_t)(in), (*C.uint8_t)(out))
	C.free(in)
	C.free(out)

	return hash
}
