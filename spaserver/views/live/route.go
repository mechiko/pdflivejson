package live

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	// Serve static and media files under /static/ and /uploads/ path.
	t.Echo().GET("/"+t.modelType.String(), t.Index)
	t.Echo().GET("/pdf.pdf", t.Pdf)
	t.Echo().GET("/scaleup", t.ScaleUp)
	t.Echo().GET("/scaledown", t.ScaleDown)
	return nil
}

func (t *page) Index(c echo.Context) error {
	data, err := t.PageData()
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("index", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) Pdf(c echo.Context) error {
	start := time.Now()
	data := t.PageModel()
	t.Logger().Debugf("file bytes since %v", time.Since(start))
	return c.Blob(http.StatusOK, "application/pdf", data.Pdf)
	// contentBytes, err := os.ReadFile("pdf.pdf")
	// if err != nil {
	// 	return t.ServerError(c, err)
	// }
	// t.Logger().Debugf("file bytes since %v", time.Since(start))
	// return c.Blob(http.StatusOK, "application/pdf", contentBytes)
}

func (t *page) ScaleUp(c echo.Context) error {
	data := t.PageModel()
	data.ViewScale += 1
	err := t.PageModelUpdate(data)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("canvas", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) ScaleDown(c echo.Context) error {
	data := t.PageModel()
	data.ViewScale -= 1
	if data.ViewScale < 1 {
		data.ViewScale = 1
	}
	err := t.PageModelUpdate(data)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("canvas", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) Serial(c echo.Context) error {
	serial := c.FormValue("serial")
	data := t.PageModel()
	data.ViewScale -= 1
	if data.ViewScale < 1 {
		data.ViewScale = 1
	}
	err := t.PageModelUpdate(data)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("canvas", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil

}
