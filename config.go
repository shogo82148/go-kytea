package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import "unsafe"

type Config struct {
	config *C.kytea_config_t
}

func (c Config) ParseRunCommandLine(args []string) {
	opts := make([]*C.char, 0, len(args))
	for _, arg := range args {
		opt := C.CString(arg)
		defer C.free(unsafe.Pointer(opt))
		opts = append(opts, opt)
	}
	C.kytea_config_parse_run_command_line(c.config, C.int(len(opts)), (**C.char)(&opts[0]))
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
