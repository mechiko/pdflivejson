package application

import (
	"fmt"
	"hash/crc32"
	"maps"
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
	jsonSrc   []byte
	ViewScale int
	template  *domain.MarkTemplate
	hash      uint32
	pdf       []byte
	params    map[string]string
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper) (*Application, error) {
	model := &Application{
		model:     domain.Application,
		Title:     "Application Title",
		ViewScale: 1,
		pdf:       make([]byte, 0),
		jsonSrc:   make([]byte, 0),
		params:    make(map[string]string),
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
	if len(m.jsonSrc) == 0 {
		m.template = &domain.MarkTemplate{}
	} else {
		if m.template, err = domain.NewMarkTemplate(m.jsonSrc); err != nil {
			m.template = &domain.MarkTemplate{}
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

// устанавливаем шаблон и константы
func (m *Application) SetTemplate(app domain.Apper, tmpl []byte, vars map[string]string) (err error) {
	maps.Copy(m.params, vars)
	m.jsonSrc = make([]byte, len(tmpl))
	copy(m.jsonSrc, tmpl)
	return m.UpdateJson(app)
}

func (m *Application) Bytes() (out []byte, err error) {
	return m.pdf, nil
}

func (m *Application) SetParams(app domain.Apper, params map[string]string) (err error) {
	maps.Copy(m.params, params)
	return m.UpdatePdf(app)
}

func (m *Application) UpdateJson(app domain.Apper) (err error) {
	crc32q := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum(m.jsonSrc, crc32q)
	if m.hash != hash {
		if len(m.jsonSrc) == 0 {
			m.template = &domain.MarkTemplate{}
		} else {
			if m.template, err = domain.NewMarkTemplate(m.jsonSrc); err != nil {
				m.template = &domain.MarkTemplate{}
				_ = m.UpdatePdf(app)
				return fmt.Errorf("%w", err)
			}
		}
		return m.UpdatePdf(app)
	}
	return nil
}

func (m *Application) UpdatePdf(app domain.Apper) (err error) {
	assets, err := assets.New("assets")
	if err != nil {
		return fmt.Errorf("Error assets: %w", err)
	}
	pdfDoc, err := pdfproc.New(app, assets)
	if err != nil {
		return fmt.Errorf("Error assets: %w", err)
	}
	m.pdf, err = pdfDoc.Create(m.template, m.params)
	return err
}

func (m *Application) GetParam(name string) string {
	if name == "" {
		return ""
	}
	if val, exist := m.params[name]; exist {
		return val
	}
	return ""
}

func (m *Application) SetParam(app domain.Apper, name, value string) error {
	if name == "" {
		return nil
	}
	m.params[name] = value
	return m.UpdatePdf(app)
}
