package live

import (
	"fmt"
	"pdflive/domain"
	"pdflive/reductor"
)

type Cis struct {
	Cis      string
	Status   string
	StatusEx string
}

type CisSlice []*Cis

type LiveModel struct {
	model  domain.Model
	Title  string
	State  int
	File   string
	Errors []string // массив ошибок
}

var _ domain.Modeler = (*LiveModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*LiveModel, error) {
	model := &LiveModel{
		model: domain.Live,
		Title: "Редактор",
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model %v read state %w", model.model, err)
	}
	return model, nil
}

// инициализируем модель вида
func (t *page) InitData(app domain.Apper) (interface{}, error) {
	model, err := NewModel(app)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	err = reductor.Instance().SetModel(model, false)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *LiveModel) SyncToStore(_ domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *LiveModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (m *LiveModel) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *m
	return &dst, nil
}

func (a *LiveModel) Model() domain.Model {
	return a.model
}

func (a *LiveModel) Save(_ domain.Apper) (err error) {
	return nil
}
