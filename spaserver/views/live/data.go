package live

import (
	"pdflive/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Instance().Model(t.modelType)
}

// с преобразованием
func (t *page) PageModel() LiveModel {
	model, _ := reductor.Instance().Model(t.modelType)
	if mdl, ok := model.(LiveModel); ok {
		return mdl
	}
	return LiveModel{}
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}
