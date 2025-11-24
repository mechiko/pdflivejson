package pdfproc

import (
	"fmt"
	"pdflive/domain"
	"pdflive/embedded"

	"github.com/johnfercher/maroto/v2"
	"github.com/mechiko/utility"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/repository"

	"github.com/johnfercher/maroto/v2/pkg/props"
)

func (p *pdfProc) Create(tmpl *domain.MarkTemplate) (out []byte, err error) {
	if err := p.BuildMaroto(tmpl); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	err = p.SetVars("party", "party")
	if err != nil {
		return nil, fmt.Errorf(" %w", err)
	}
	err = p.SetVars("idx", fmt.Sprintf("%06d", 100))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	code := `0105000213100066215qiDHO-93lijm`
	codeCis, _ := utility.ParseCisInfo(code)

	if err := p.AddPageByTemplate(tmpl, []*utility.CisInfo{codeCis}); err != nil {
		return nil, fmt.Errorf("add datamatrix KM in page %w", err)
	}

	err = p.DocumentGenerate()
	if err != nil {
		return nil, fmt.Errorf("генерация пдф блока ошибка %w", err)
	}

	return p.PdfBytes()
}

func (p *pdfProc) PdfBytes() (out []byte, err error) {
	if p.document == nil {
		return nil, fmt.Errorf("document is nil")
	}
	out = p.document.GetBytes()
	return out, nil
}

func (p *pdfProc) PdfDocumentSave(fileName string) (err error) {
	if p.document == nil {
		return fmt.Errorf("save document: document is nil")
	}
	err = p.document.Save(fileName)
	if err != nil {
		return fmt.Errorf("save document: %w", err)
	}
	return nil
}

func (p *pdfProc) PdfDocumentReportSave(fileName string) (err error) {
	// запись отчета генерации
	if p.document == nil {
		return fmt.Errorf("save document report: document is nil")
	}
	err = p.document.GetReport().Save(fileName)
	if err != nil {
		return fmt.Errorf("save document report: %w", err)
	}
	return nil
}

func (p *pdfProc) BuildMaroto(tmpl *domain.MarkTemplate) (err error) {
	customFont := "roboto"
	customFonts, err := repository.New().
		AddUTF8FontFromBytes(customFont, fontstyle.Normal, embedded.Regular).
		AddUTF8FontFromBytes(customFont, fontstyle.Italic, embedded.Italic).
		AddUTF8FontFromBytes(customFont, fontstyle.Bold, embedded.Bold).
		AddUTF8FontFromBytes(customFont, fontstyle.BoldItalic, embedded.BoldItalic).
		Load()
	if err != nil {
		return err
	}
	// …custom font setup…
	builder := config.NewBuilder().WithCustomFonts(customFonts).WithCompression(true)
	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).Build()
	cfg.Dimensions.Height = tmpl.PageHeight
	cfg.Dimensions.Width = tmpl.PageWidth
	cfg.Margins.Bottom = tmpl.Margin.Bottom
	cfg.Margins.Top = tmpl.Margin.Top
	cfg.Margins.Left = tmpl.Margin.Left
	cfg.Margins.Right = tmpl.Margin.Right
	cfg.DefaultFont.Size = 4
	p.maroto = maroto.New(cfg)
	p.maroto = maroto.NewMetricsDecorator(p.maroto)
	return nil
}

func (p *pdfProc) DocumentGenerate() (err error) {
	if p.maroto == nil {
		return fmt.Errorf("document generate: maroto is not initialized (call BuildMaroto first)")
	}
	doc, err := p.maroto.Generate()
	if err != nil {
		return fmt.Errorf("document generate: %w", err)
	}
	p.document = doc
	return nil
}

func (p *pdfProc) AddPageByTemplate(tmpl *domain.MarkTemplate, cis []*utility.CisInfo) error {
	if tmpl == nil {
		return fmt.Errorf("add page: template is nil")
	}
	if p.maroto == nil {
		return fmt.Errorf("add page: maroto is not initialized (call BuildMaroto first)")
	}
	pgNew, err := p.Page(tmpl, cis)
	if err != nil {
		return fmt.Errorf("build page: %w", err)
	}
	if pgNew == nil {
		return fmt.Errorf("build page: result is nil")
	}
	p.maroto.AddPages(pgNew)
	return nil
}
