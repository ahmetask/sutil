package sutil

type OrElse func() interface{}

type Optional interface {
	Exist() bool
	Value() interface{}
	OrElse(orElse OrElse) interface{}
}

type Data struct {
	V interface{}
}

func (d *Data) Exist() bool {
	return d.V != nil
}

func (d *Data) Value() interface{} {
	return d.V
}

func (d *Data) OrElse(f OrElse) interface{} {
	if d.Exist() {
		return d.Value()
	}
	return f()
}
