package randomhash

/* 
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -lstdc++ -lrh 
#include "bridge.h"
#include <stdlib.h>
*/
import "C"

import (
	"math/big"
	"unsafe"
)

type RandomHash struct {
	hashrate int64
}

func (rh *RandomHash) ScanHash(block []byte, stop <-chan struct{}, index int) uint64 {
	return 0
}

func (rh *RandomHash) GetHashrate() int64 {
	return rh.hashrate
}

func (rh *RandomHash) Verify(block []byte) bool {
	// C.RandomHash_Simple()
	return false
}

// Search , aka ScanHash
func (rh *RandomHash) Search(block []byte, index int) uint64 {
	return 0
}

func (rh *RandomHash) CalcHash(block []byte, nonce uint64) *big.Int {
	var in unsafe.Pointer = C.CBytes(block)
	var out unsafe.Pointer = C.malloc(8)
	C.RandomHash_Simple2((*C.uchar)(in), (*C.uchar)(out))
	// C.test()
	C.free(in)
	C.free(out)

	return nil
}
