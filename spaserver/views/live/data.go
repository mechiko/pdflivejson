package live

import (
	"pdflive/domain"
	"pdflive/domain/models/application"
	"pdflive/reductor"
)

// отходим от шаблона и модель у нас всего приложения
func (t *page) PageData() (interface{}, error) {
	return reductor.Model[*application.Application](domain.Application)
}

// с преобразованием
func (t *page) PageModel() *application.Application {
	model, _ := reductor.Model[*application.Application](domain.Application)
	return model
}

func (t *page) PageModelUpdate(model *application.Application) error {
	err := reductor.SetModel(model, false)
	return err
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}
