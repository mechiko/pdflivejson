package application

import (
	"fmt"
	"pdflive/config"
	"pdflive/domain"
)

type Application struct {
	model    domain.Model
	Title    string
	Debug    bool
	License  string
	JsonSrc  []byte
	Template *domain.MarkTemplate
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper) (*Application, error) {
	model := &Application{
		model: domain.Application,
		Title: "Application Title",
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *Application) SyncToStore(app domain.Apper) (err error) {
	return nil
}

// читаем состояние приложения
func (m *Application) ReadState(app domain.Apper) (err error) {
	m.Debug = config.Mode == "development"
	m.License = app.Options().Application.License
	if len(m.JsonSrc) == 0 {
		m.Template = &domain.MarkTemplate{}
	} else {
		if m.Template, err = domain.NewMarkTemplate(m.JsonSrc); err != nil {
			m.Template = &domain.MarkTemplate{}
			return fmt.Errorf("%w", err)
		}

	}
	return nil
}

func (a *Application) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *a
	return &dst, nil
}

func (a *Application) Model() domain.Model {
	return a.model
}

func (m *Application) Save(app domain.Apper) (err error) {
	if err := app.SaveOptions(); err != nil {
		return fmt.Errorf("application: save options failed: %w", err)
	}
	return nil
}

// читаем состояние приложения
func (m *Application) SetTemplate(tmpl []byte) (err error) {
	if len(tmpl) == 0 {
		m.Template = &domain.MarkTemplate{}
	} else {
		if m.Template, err = domain.NewMarkTemplate(tmpl); err != nil {
			m.Template = &domain.MarkTemplate{}
			return fmt.Errorf("%w", err)
		}

	}
	return nil
}
