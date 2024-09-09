package generate

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/varmanager"
)

var (
	_ Generator = (*NestGenerator)(nil)
)

type NestGenerator struct {
	cfg        *config.Config
	varManager *varmanager.VarManager

	tpl       *template.Template
	outputDir string
}

func NewNestGenerator(cfg *config.Config, varManager *varmanager.VarManager, tpl *template.Template) Generator {
	return &NestGenerator{
		cfg:        cfg,
		outputDir:  fmt.Sprintf("%s", cfg.Output.Dir),
		varManager: varManager,
		tpl:        tpl,
	}
}

func (g *NestGenerator) Generate(ctx context.Context) (string, error) {
	return "nest", nil
}

// 按cfg.NestRoot的嵌套结构创建目录层级
// nestRoot:
//
//	name: root
//	nest:
//	  - name: api
//	    package: api
//	    nest:
//	      - name: *.go
//	  - name: internal
//	    package: internal
//	    nest:
//	      - name: model
//	        package: model
//	        nest:
//	          - name: model_base.go
//	          - name: model_*.go
//	          - name: model_*_gen.go
//	      - name: dao
//	        package: dao
//	        nest:
//	          - name: dao_base.go
//	          - name: dao_*.go
//	          - name: dao_*_gen.go
//	      - name: service
//	        package: service
//	        nest:
//	          - name: service_*.go
//	      - name: server
//	        package: server
//	        nest:
//	          - name: http
//	            package: http
//	            nest:
//	              - name: *.go
func (g *NestGenerator) createNest(ctx context.Context, nest *config.NestConf) error {
	if nest == nil {
		return nil
	}

	isPackage := false
	if len(nest.Nest) > 0 {
		isPackage = true
	}

	if isPackage {
		packageName := nest.PackageName
		if packageName == "" {
			packageName = nest.Name
		}

		// 创建package目录
		if err := os.MkdirAll(filepath.Join(g.outputDir), 0755); err != nil {
			return err
		}
		// 递归创建nest
		for _, n := range nest.Nest {
			if err := g.createNest(ctx, n); err != nil {
				return err
			}
		}
		return nil
	}

	if len(nest.Name) > 0 {
		// 调用方法,使用模板创建文件
		// 加载模板
		// err = g.tpl.Lookup(nest.Name)
	}
	return nil
}
