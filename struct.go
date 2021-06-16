package sutil

import (
	"errors"
	"reflect"
	"strings"
)

type path struct {
	json  bool
	value string
}

type SUtil struct {
	data    interface{}
	value   interface{}
	res     interface{}
	path    path
	success bool
	err     error
}

type ISUtil interface {
	WithValue(v interface{}) ISUtil
	WithPath(p string, json bool) ISUtil
	Set() (bool, error)
	Get() Optional
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

func (s *SUtil) reflectValue(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}

func (s *SUtil) result(err error) {
	if err != nil {
		s.success = false
		s.err = err
	} else {
		s.success = true
		s.err = nil
	}

}
func (s *SUtil) set(r reflect.Value, v interface{}) {
	if !r.IsValid() {
		s.result(errors.New("invalid value"))
		return
	}
	if r.CanSet() {
		if v == nil {
			r.Set(reflect.Zero(r.Type()))
			s.result(nil)
			return
		}
		vPrime := s.reflectValue(v)
		if r.Kind() == vPrime.Kind() {
			r.Set(vPrime)
			s.result(nil)
		}
	} else {
		s.result(errors.New("value can not be set"))
	}
}

func (s *SUtil) setR(o interface{}, v interface{}, base string) {
	rv := s.reflectValue(o)

	if rv.Kind() == reflect.Invalid {
		return
	} else if rv.Kind() != reflect.Struct {
		if s.path.value == strings.TrimPrefix(base, ".") {
			s.set(rv, v)
		}
		return
	}

	reflectType := rv.Type()

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)

		p := ""
		if s.path.json {
			p = base + "." + field.Tag.Get("json")
		} else {
			p = base + "." + field.Name
		}

		fieldI := rv.Field(i)
		pathMatch := s.path.value == strings.TrimPrefix(p, ".")
		if fieldI.Kind() == reflect.Struct || fieldI.Kind() == reflect.Ptr {
			if pathMatch {
				if fieldI.CanSet() {
					if fieldI.Kind() == reflect.Ptr && v != nil {
						fieldI = fieldI.Elem()
					}
					s.set(fieldI, v)
				}
			} else {
				s.setR(fieldI.Interface(), v, p)
			}
		} else {
			if pathMatch {
				s.set(fieldI, v)
			}
		}
	}

	return
}

func (s *SUtil) getR(o interface{}, base string) {
	rv := s.reflectValue(o)

	if rv.Kind() == reflect.Invalid {
		return
	} else if rv.Kind() != reflect.Struct {
		if s.path.value == strings.TrimPrefix(base, ".") {
			s.res = rv.Interface()
		}
		return
	}

	reflectType := rv.Type()

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)

		p := ""
		if s.path.json {
			p = base + "." + field.Tag.Get("json")
		} else {
			p = base + "." + field.Name
		}

		fieldI := rv.Field(i)
		pathMatch := s.path.value == strings.TrimPrefix(p, ".")
		if fieldI.Kind() == reflect.Struct || fieldI.Kind() == reflect.Ptr {
			if pathMatch {
				if fieldI.Kind() == reflect.Ptr {
					fieldI = fieldI.Elem()
				}
				s.res = fieldI.Interface()
			} else {
				s.getR(fieldI.Interface(), p)
			}
		} else {
			if pathMatch {
				s.res = fieldI.Interface()
			}
		}
	}

	return
}

func (s *SUtil) Set() (bool, error) {
	s.setR(s.data, s.value, "")
	return s.success, s.err
}

func (s *SUtil) Get() Optional {

	s.getR(s.data, "")
	return &Data{V: s.res}
}
