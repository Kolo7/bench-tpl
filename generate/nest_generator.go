package generate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/varmanager"
)

var (
	_ Generator = (*NestGenerator)(nil)
)

type NestGenerator struct {
	cfg        *config.Config
	varManager *varmanager.VarManager

	tpl       *template.Template
	outputDir string
	table     string
}

func NewNestGenerator(cfg *config.Config, table string, varManager *varmanager.VarManager, tpl *template.Template) Generator {
	return &NestGenerator{
		cfg:        cfg,
		outputDir:  cfg.OutputDir,
		varManager: varManager,
		tpl:        tpl,
		table:      table,
	}
}

func (g *NestGenerator) Generate(ctx context.Context) (string, error) {
	for _, nest := range config.DefaultNestConf {
		if err := g.createNest(ctx, nest); err != nil {
			return "", err
		}

	}

	return "", nil
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
func (g *NestGenerator) createNest(ctx context.Context, nest *config.NestConf, prefix ...string) error {
	if nest == nil {
		return nil
	}

	isPackage := false
	if len(nest.Nest) > 0 {
		isPackage = true
	}

	if isPackage {
		packageName := nest.Package
		if packageName == "" {
			packageName = nest.Name
		}
		// 递归创建nest
		for _, n := range nest.Nest {
			if err := g.createNest(ctx, n, append(prefix, packageName)...); err != nil {
				return err
			}
		}
		return nil
	}

	if len(nest.Name) > 0 {
		// 设置包级别变量
		g.varManager.SetPackageVar(prefix...)
		// 调用方法,使用模板创建文件
		// 加载模板
		tpl := g.tpl.Lookup(nest.Name + ".tpl")
		if tpl == nil {
			return nil
		}
		// 创建文件，如果文件已经存在，则不创建
		var (
			f          *os.File
			goFileName string
		)
		outDir := filepath.Join(g.outputDir, filepath.Join(prefix...))
		if nest.PkgUnique {
			goFileName = filepath.Join(outDir, nest.Name+".go")
		} else {
			goFileName = filepath.Join(outDir, g.table+"_"+nest.Name+".go")
		}
		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			// 创建多级目录
			if err := os.MkdirAll(outDir, 0755); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		// 判断文件是否存在，如果存在并且不覆盖，则不创建
		_, err := os.Stat(goFileName)
		if os.IsNotExist(err) {
			fmt.Printf("create file %s\n", goFileName)
		} else if !nest.Override {
			fmt.Printf("skip file %s, already exists\n", goFileName)
			return nil
		} else if nest.Override {
			fmt.Printf("override file %s\n", goFileName)
		} else if err != nil {
			return err
		}
		// 打开文件，覆盖原文件
		f, err = os.OpenFile(goFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}

		defer f.Close()
		if err := tpl.Execute(f, g.varManager.GetTableVar(g.table)); err != nil {
			return err
		}
	}
	return nil
}
