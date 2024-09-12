package web

import (
	"context"

	"github.com/a-h/templ"
	"github.com/is1ab/Arvosana/web/template"
)

//go:generate templ generate
//go:generate pnpx tailwindcss -i template/styles.css -o static/styles.css --minify
func Render(ctx context.Context, t templ.Component) (string, error) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	base := template.Base(t)

	err := base.Render(ctx, buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
