package easy_error

import "errors"
import "os"

var ErrUnknown = errors.New("Unknown error")
func Apply(r interface{}) (err error) {
	if r != nil {
		ee, ok := r.(error)
		if !ok {
			err = ErrUnknown
		} else {
			err = ee
		}
	}
	return
}

func AssertNil(err error) {
	if err != nil {
		panic(err)
	}
}

func Wrap(v interface{}, err error) (r interface{}) {
	AssertNil(err)
	return v
}

func WrapFile(v *os.File, err error) (r *os.File) {
	AssertNil(err)
	return v
}

func WrapBytes(v []byte, err error) (r []byte) {
	AssertNil(err)
	return v
}
