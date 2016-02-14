package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type Word struct {
	word *C.kytea_word_t
}

func (w Word) Surface(util StringUtil) string {
}
