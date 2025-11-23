package home

import (
	"pdflive/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Instance().Model(t.modelType)
}

// с преобразованием
func (t *page) PageModel() HomeModel {
	model, _ := reductor.Instance().Model(t.modelType)
	if mdl, ok := model.(HomeModel); ok {
		return mdl
	}
	return HomeModel{}
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}
