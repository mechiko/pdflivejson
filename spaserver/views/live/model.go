package live

import (
	"fmt"
	"pdflive/domain"
	"pdflive/domain/models/application"
	"pdflive/reductor"
)

// инициализируем модель вида
func (t *page) InitData(app domain.Apper) (interface{}, error) {
	model, err := reductor.Model[*application.Application](domain.Application)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return model, nil
}
