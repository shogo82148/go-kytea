package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import "unsafe"

type Config struct {
	config *C.kytea_config_t
}

func (c Config) ModelFile() string {
	path := C.kytea_config_get_model_file(c.config)
	defer C.kytea_std_string_destroy(path)
	return C.GoString(C.kytea_std_string_cstring(path))
}

func (c Config) SetModelFile(path string) {
	s := C.CString(path)
	defer C.free(unsafe.Pointer(s))
	C.kytea_config_set_model_file(c.config, s)
}
