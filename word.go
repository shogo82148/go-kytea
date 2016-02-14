package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type KyTeaWord struct {
	word *C.kytea_word_t
}

func (w KyTeaWord) Surface(util StringUtil) string {
	s := C.kytea_word_surface(w.word, util.util)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}
