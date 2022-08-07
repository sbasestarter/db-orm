package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}
{{end}}

{{range .Tables}}
{{$tableName := .Name}}
{{$clsName := Mapper .Name}}
type {{Mapper .Name}}Helper struct {

}

{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}
func (cols {{$clsName}}Helper) {{Mapper $col.Name}}() string {
	return "{{ $col.Name }}"
}

func (cols {{$clsName}}Helper) {{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }}"
}

func (cols {{$clsName}}Helper) Eq{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} = ?"
}

func (cols {{$clsName}}Helper) Eq{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} = ?"
}

func (cols {{$clsName}}Helper) In{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} in ?"
}

func (cols {{$clsName}}Helper) In{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} in ?"
}

func (cols {{$clsName}}Helper) Ne{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} != ?"
}

func (cols {{$clsName}}Helper) Ne{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} != ?"
}

func (cols {{$clsName}}Helper) Gt{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} > ?"
}

func (cols {{$clsName}}Helper) Gt{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} > ?"
}

func (cols {{$clsName}}Helper) Gte{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} >= ?"
}

func (cols {{$clsName}}Helper) Gte{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} >= ?"
}

func (cols {{$clsName}}Helper) Lt{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} < ?"
}

func (cols {{$clsName}}Helper) Lt{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} < ?"
}

func (cols {{$clsName}}Helper) Lte{{Mapper $col.Name}}() string {
	return "{{ $col.Name }} <= ?"
}

func (cols {{$clsName}}Helper) Lte{{Mapper $col.Name}}WT() string {
	return "{{ $tableName }}.{{ $col.Name }} <= ?"
}
{{end}}

func (cols {{$clsName}}Helper) TableName() string {
	return "{{ .Name }}"
}

// O{{Mapper .Name}} .
var O{{Mapper .Name}} = {{Mapper .Name}}Helper{};

{{end}}
