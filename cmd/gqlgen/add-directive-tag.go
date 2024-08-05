//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/vektah/gqlparser/v2/ast"
)

func fieldHook(f *modelgen.Field, fd *ast.FieldDefinition) {
	// @tagディレクティブ
	directive := fd.Directives.ForName("tag")
	if directive != nil {
		// validateタグを追加
		validateTag := directive.Arguments.ForName("validate")
		if validateTag != nil {
			f.Tag += fmt.Sprintf(` validate:"%s"`, validateTag.Value.Raw)
		}
	}
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
			for _, model := range b.Models {
				for i, field := range model.Fields {
					fieldHook(field, cfg.Schema.Types[model.Name].Fields[i])
				}
			}

			return b
		},
	}

	// ディレクティブとしては使用しないのでコードに出力されないように設定
	// https://github.com/99designs/gqlgen/blob/v0.14.0/codegen/config/config.go#L231
	cfg.Directives["tag"] = config.DirectiveConfig{
		SkipRuntime: true,
	}

	err = api.Generate(cfg,
		func(cfg *config.Config, plugins *[]plugin.Plugin) {
			for i, plugin := range *plugins {
				if _, ok := plugin.(*modelgen.Plugin); ok {
					// modelgen.Pluginを置き換える
					(*plugins)[i] = &p
				}
			}
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
