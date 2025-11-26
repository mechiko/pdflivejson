package pdfproc

import "fmt"

// значения переменных для шаблона
// индекс страницы
// партия и другие
type Params struct {
	value map[string]string
}

func NewParams() *Params {
	out := &Params{
		value: map[string]string{},
	}
	return out
}

func (v *Params) Add(name string, value string) error {
	if name == "" {
		return fmt.Errorf("is empty key")
	}
	v.value[name] = value
	return nil
}

func (v *Params) Get(name string) string {
	if name == "" {
		return ""
	}
	if val, exist := v.value[name]; exist {
		return val
	}
	return ""
}
