package live

import (
	"pdflive/domain"
	"pdflive/domain/models/application"
	"pdflive/reductor"
)

// отходим от шаблона и модель у нас всего приложения
func (t *page) PageData() (interface{}, error) {
	return reductor.Instance().Model(domain.Application)
}

// с преобразованием
func (t *page) PageModel() *application.Application {
	model, _ := reductor.Instance().Model(domain.Application)
	if mdl, ok := model.(*application.Application); ok {
		return mdl
	}
	return nil
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}
