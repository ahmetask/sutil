package sutil

type path struct {
	json  bool
	value string
}

type SUtil struct {
	data  interface{}
	value interface{}
	path  path
}

type ISUtil interface {
	WithValue(v interface{}) ISUtil
	WithPath(p string, json bool) ISUtil
	Set() (bool, error)
}

func New(data interface{}) ISUtil {
	return &SUtil{
		data: data,
	}
}
func (s *SUtil) WithValue(v interface{}) ISUtil {
	s.value = v
	return s
}

func (s *SUtil) WithPath(p string, json bool) ISUtil {
	s.path = path{value: p, json: json}
	return s
}

func (s *SUtil) Set() (bool, error) {


	return true, nil
}
