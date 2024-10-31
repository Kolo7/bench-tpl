package input

import (
	"fmt"
	"io/fs"
	"text/template"

	"github.com/Kolo7/bench-tpl/config"
)

type TplFileInput struct {
	cfg *config.Config

	inputDir string
	tpl      *template.Template
	fs       fs.FS
}

func NewTplFileInput(cfg *config.Config) *TplFileInput {

	return &TplFileInput{
		cfg:      cfg,
		inputDir: cfg.InputDir,
		fs:       cfg.FS,
	}
}

// tplName 模板文件名
func (t *TplFileInput) LoadTemplate(tpl *template.Template, tplName string) (*template.Template, error) {
	if t.tpl == nil {
		var err error
		// 加载目录
		// glob 模式匹配目录下所有模板文件，不管目录深度
		dir := fmt.Sprintf("%s/**/*.tpl", t.inputDir)
		if tpl == nil {
			tpl = template.New("")
		}

		t.tpl, err = tpl.ParseFS(t.fs, dir)
		if err != nil {
			return nil, err
		}
	}
	if tplName == "" {
		return t.tpl, nil
	}
	return t.tpl.Lookup(tplName), nil
}
