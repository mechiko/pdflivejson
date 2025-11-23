package live

import (
	"fmt"
	"pdflive/config"
	"pdflive/domain"
	"time"

	"github.com/mechiko/utility"
)

type Live struct {
	model        domain.Model
	Title        string
	Export       string
	Browser      utility.Browser
	BrowserList  []string
	Output       string
	Debug        bool
	Host         string
	Port         string
	DbLiteDesc   string
	DbConfigDesc string
	DbZnakDesc   string
	DbA3Desc     string
	License      string
	FsrarID      string
	startTime    time.Time
	endTime      time.Time
	period       string
}

var _ domain.Modeler = (*Live)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper) (*Live, error) {
	model := &Live{
		model:       domain.Application,
		Title:       "Application Title",
		BrowserList: []string{string(utility.Default), string(utility.Chrome), string(utility.Firefox), string(utility.Yandex), string(utility.Edge)},
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *Live) SyncToStore(app domain.Apper) (err error) {
	if err := app.SetOptions("export", m.Export); err != nil {
		return fmt.Errorf("application sync to store: set export failed: %w", err)
	}
	if err := app.SetOptions("browser", string(m.Browser)); err != nil {
		return fmt.Errorf("application sync to store: set browser failed: %w", err)
	}
	return nil
}

// читаем состояние приложения
func (m *Live) ReadState(app domain.Apper) (err error) {
	m.Export = app.Options().Export
	m.Browser = utility.Browser(app.Options().Browser)
	m.Output = app.Options().Output
	m.Host = app.Options().Hostname
	m.Port = app.Options().HostPort
	m.Debug = config.Mode == "development"
	return nil
}

func (a *Live) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *a
	return &dst, nil
}

func (a *Live) Model() domain.Model {
	return a.model
}

func (m *Live) Save(app domain.Apper) (err error) {
	if err := app.SaveOptions(); err != nil {
		return fmt.Errorf("application: save options failed: %w", err)
	}
	return nil
}
