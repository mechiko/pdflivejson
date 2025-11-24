package live

import (
	"fmt"
	"pdflive/domain"
	"pdflive/domain/models/application"
	"pdflive/reductor"
)

// инициализируем модель вида
func (t *page) InitData(app domain.Apper) (interface{}, error) {
	model, err := reductor.Instance().Model(domain.Application)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	mdl, ok := model.(*application.Application)
	if !ok {
		return nil, fmt.Errorf("%w", err)
	}
	err = reductor.Instance().SetModel(mdl, false)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return model, nil
}
