package application

import (
	"fmt"
	"hash/crc32"
	"pdflive/assets"
	"pdflive/config"
	"pdflive/domain"
	"pdflive/pdfproc"
)

type Application struct {
	model     domain.Model
	Title     string
	Debug     bool
	License   string
	JsonSrc   []byte
	ViewScale int
	Template  *domain.MarkTemplate
	Hash      uint32
	Pdf       []byte
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper) (*Application, error) {
	model := &Application{
		model:     domain.Application,
		Title:     "Application Title",
		ViewScale: 1,
		Pdf:       make([]byte, 0),
		JsonSrc:   make([]byte, 0),
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
func (m *Application) SetTemplate(app domain.Apper, tmpl []byte) (err error) {
	crc32q := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum(tmpl, crc32q)
	if m.Hash != hash {
		if len(tmpl) == 0 {
			m.Template = &domain.MarkTemplate{}
		} else {
			if m.Template, err = domain.NewMarkTemplate(tmpl); err != nil {
				m.Template = &domain.MarkTemplate{}
				return fmt.Errorf("%w", err)
			}
		}
		assets, err := assets.New("assets")
		if err != nil {
			return fmt.Errorf("Error assets: %w", err)
		}
		pdfDoc, err := pdfproc.New(app, assets)
		if err != nil {
			return fmt.Errorf("Error assets: %w", err)
		}
		m.Pdf, err = pdfDoc.Create(m.Template)
		return err
	}
	return nil
}

func (m *Application) Bytes() (out []byte, err error) {
	return m.Pdf, nil
}
