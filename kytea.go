package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import (
	"errors"
	"unsafe"
)

type KyTea struct {
	kytea *C.kytea_t
}

func New() (KyTea, error) {
	kytea := C.kytea_new()
	if kytea == nil {
		return KyTea{}, errors.New("kytae is not created")
	}
	return KyTea{kytea: kytea}, nil
}

func (k KyTea) Destory() {
	C.kytea_destroy(k.kytea)
}

func (k KyTea) ReadModel(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.kytea_read_model(k.kytea, cpath)
}

func (k KyTea) StringUtil() StringUtil {
	return StringUtil{
		util: C.kytea_get_string_util(k.kytea),
	}
}

func (k KyTea) GetWS(sentence string) []string {
	if sentence == "" {
		return []string{}
	}

	s := *(**C.char)(unsafe.Pointer(&sentence))
	vec := C.kytea_get_ws(k.kytea, s, C.size_t(len(sentence)))
	defer C.kytea_vector_string_destory(vec)

	size := int(C.kytea_vector_string_size(vec))
	ret := make([]string, size)
	for i := 0; i < size; i++ {
		ret[i] = C.GoString(C.kytea_vector_string_at(vec, C.int(i)))
	}
	return ret
}
