package sql2struct

import (
	"fmt"
	"gcli/pkg/word"
	"os"
	"text/template"
)

const structTpl = `type {{.TableName|ToCamelCase}} struct {
{{range .Columns}} {{$length:=len .Comment}}{{if gt $lenght 0}}//{{.Comment}}{{else}}//{{.Name}}{{end}}
{{$typeLen:=len .Type}} {{if gt $typeLen 0}}{{.Name|ToCamelCase}}
{{.Type}} {{.Tag}}{{else}}{{.Name}}{{end}}
{{end}}}
func (model {{.TableName|ToCamelCase}}) TableName()string{
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssenblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"join:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DateType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplcolumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplcolumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
