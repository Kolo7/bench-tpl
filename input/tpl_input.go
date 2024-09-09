package input

import (
	"fmt"
	"html/template"

	"github.com/kolo7/bench-tpl/config"
)

type TplFileInput struct {
	cfg *config.Config

	inputDir string
	tpl      *template.Template
}

func NewTplFileInput(cfg *config.Config) *TplFileInput {

	return &TplFileInput{
		cfg:      cfg,
		inputDir: cfg.Input.Dir,
	}
}

// tplName 模板文件名
func (t *TplFileInput) LoadTemplate(tplName string) (*template.Template, error) {
	if t.tpl == nil {
		var err error
		// 加载目录
		// glob 模式匹配文件
		dir := fmt.Sprintf("%s/*.tpl", t.inputDir)
		t.tpl, err = template.ParseGlob(dir)
		if err != nil {
			return nil, err
		}
	}
	if tplName == "" {
		return t.tpl, nil
	}
	return t.tpl.Lookup(tplName), nil
}
