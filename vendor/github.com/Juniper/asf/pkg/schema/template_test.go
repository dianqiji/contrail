package schema

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	schemaPath         = "test_data/simple/schema"
	templateConfigPath = "test_data/templates/template_config.yaml"

	allOutputPath       = "test_output/gen_all.yml"
	hogeGoOutputPath    = "test_output/gen_hoge.go"
	hogeProtoOutputPath = "test_output/gen_hoge.proto"
	hogeSQLOutputPath   = "test_output/gen_hoge.sql"
)

func TestGenerateFiles(t *testing.T) {
	for _, tt := range []struct {
		name               string
		generateConfig     GenerateConfig
		expectedOutputPath string
		expectedContent    string
	}{{
		name: "adds generation comment given YAML file",
		generateConfig: GenerateConfig{
			TemplateConfigs: loadTemplateConfigs(t),
		},
		expectedOutputPath: allOutputPath,
		expectedContent:    "^# Code generated by contrailschema tool .* DO NOT EDIT.",
	}, {
		name: "adds generation comment given Go file",
		generateConfig: GenerateConfig{
			TemplateConfigs: loadTemplateConfigs(t),
		},
		expectedOutputPath: hogeGoOutputPath,
		expectedContent:    "^// Code generated by contrailschema tool .* DO NOT EDIT.",
	}, {
		name: "adds generation comment given Proto file",
		generateConfig: GenerateConfig{
			TemplateConfigs: loadTemplateConfigs(t),
		},
		expectedOutputPath: hogeProtoOutputPath,
		expectedContent:    "^// Code generated by contrailschema tool .* DO NOT EDIT.",
	}, {
		name: "adds generation comment given SQL file",
		generateConfig: GenerateConfig{
			TemplateConfigs: loadTemplateConfigs(t),
		},
		expectedOutputPath: hogeSQLOutputPath,
		expectedContent:    "^-- Code generated by contrailschema tool .* DO NOT EDIT.",
	}, {
		name: "fills given models import path",
		generateConfig: GenerateConfig{
			TemplateConfigs:  loadTemplateConfigs(t),
			ModelsImportPath: "github.com/custom/models",
		},
		expectedOutputPath: hogeGoOutputPath,
		expectedContent:    "// import models \"github.com/custom/models\"",
	}, {
		name: "fills given services import path",
		generateConfig: GenerateConfig{
			TemplateConfigs:    loadTemplateConfigs(t),
			ServicesImportPath: "github.com/custom/services",
		},
		expectedOutputPath: hogeGoOutputPath,
		expectedContent:    "// import services \"github.com/custom/services\"",
	}} {
		t.Run(tt.name, func(t *testing.T) {
			err := GenerateFiles(makeAPI(t), &tt.generateConfig)

			assert.NoError(t, err)
			assert.Regexp(t, tt.expectedContent, loadFile(t, tt.expectedOutputPath))
		})
	}
}

func makeAPI(t *testing.T) *API {
	api, err := MakeAPI([]string{schemaPath}, nil, false)
	require.NoError(t, err)
	return api
}

func loadTemplateConfigs(t *testing.T) []TemplateConfig {
	c, err := LoadTemplateConfigs(templateConfigPath)
	require.NoError(t, err)
	return c
}

func loadFile(t *testing.T, path string) string {
	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	return string(data)
}

func TestResolveOutputPath(t *testing.T) {
	for _, tt := range []struct {
		templateConfig     TemplateConfig
		expectedOutputPath string
	}{{
		templateConfig: TemplateConfig{
			TemplatePath: "/absolute/file.go.tmpl",
		},
		expectedOutputPath: "/absolute/gen_file.go",
	}, {
		templateConfig: TemplateConfig{
			TemplatePath: "relative/file.go.tmpl",
		},
		expectedOutputPath: "relative/gen_file.go",
	}, {
		templateConfig: TemplateConfig{
			TemplatePath: "file.go.tmpl",
		},
		expectedOutputPath: "gen_file.go",
	}, {
		templateConfig: TemplateConfig{
			TemplatePath: "not_a_template.go",
		},
		expectedOutputPath: "gen_not_a_template.go",
	}, {
		templateConfig: TemplateConfig{
			TemplatePath: "file.go.tmpl",
			OutputDir:    "/absolute_output_dir",
		},
		expectedOutputPath: "/absolute_output_dir/gen_file.go",
	}, {
		templateConfig: TemplateConfig{
			TemplatePath: "file.go.tmpl",
			OutputDir:    "relative_output_dir",
		},
		expectedOutputPath: "relative_output_dir/gen_file.go",
	}} {
		t.Run(tt.templateConfig.TemplatePath+tt.templateConfig.OutputDir, func(t *testing.T) {
			resolveOutputPath(&tt.templateConfig)
			assert.Equal(t, tt.expectedOutputPath, tt.templateConfig.OutputPath)
		})
	}
}
