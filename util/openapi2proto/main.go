package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Field struct {
	Name   string
	Type   string
	Id     int
	Source *openapi3.SchemaRef
}

type Message struct {
	Name   string
	Fields []Field
}

type Enum struct {
	Name   string
	Values []interface{}
}

func getType(p *openapi3.SchemaRef) string {
	switch p.Value.Type {
	case "integer":
		return "int32"
	case "boolean":
		return "bool"
	case "number":
		return "double"
	case "object":
		if p.Ref != "" {
			t := strings.Split(p.Ref, "/")
			return t[len(t)-1]
		}
		if p.Value.AdditionalProperties != nil {
			return fmt.Sprintf("map<string,%s>", getType(p.Value.AdditionalProperties))
		}
		return fmt.Sprintf("map<string,string>")
		//return fmt.Sprintf("%#v", p.Value)
	case "array":
		if p.Value.Items.Ref != "" {
			t := strings.Split(p.Value.Items.Ref, "/")
			return "repeated " + t[len(t)-1]
		}
		return "repeated " + getType(p.Value.Items)
	default:
		if p.Ref != "" {
			t := strings.Split(p.Ref, "/")
			return t[len(t)-1]
		}
		return p.Value.Type
	}
}

func getFields(schema *openapi3.SchemaRef) []Field {
	fields := []Field{}
	for pname, pschema := range schema.Value.Properties {
		//	fmt.Printf("\t%s %#v\n", pname, pschema)
		f := Field{Name: pname, Type: getType(pschema), Source: schema}
		fields = append(fields, f)
	}
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	for i := range fields {
		fields[i].Id = i + 1
	}
	return fields
}

func parseMessageSchema(name string, schema *openapi3.SchemaRef) (Message, error) {
	if schema.Value.Properties != nil {
		//fmt.Printf("%s %#v\n", name, schema.Value.Properties)
		fields := getFields(schema)
		m := Message{Name: name, Fields: fields}
		return m, nil
	} else if schema.Value.AllOf != nil {
		//fmt.Printf("All of %s %#v\n", name, schema.Value.AllOf)
		fieldMap := map[string]Field{}
		for i := range schema.Value.AllOf {
			newFields := getFields(schema.Value.AllOf[i])
			//fmt.Printf("\t%#v\n", newFields)
			for _, f := range newFields {
				if x, ok := fieldMap[f.Name]; !ok {
					fieldMap[f.Name] = f
				} else {
					//if same field from two sources, pick local one
					if x.Source.Ref != "" && f.Source.Ref == "" {
						fieldMap[f.Name] = f
					}
				}
			}
		}
		fields := []Field{}
		for _, v := range fieldMap {
			fields = append(fields, v)
		}
		//fmt.Printf("Fields: %#vs\n", fields)
		sort.SliceStable(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
		for i := range fields {
			fields[i].Id = i + 1
		}
		m := Message{Name: name, Fields: fields}
		return m, nil
	} else if schema.Value.Enum != nil {
		return Message{}, fmt.Errorf("Message is Enum")
	} else {
		m := Message{Name: name, Fields: []Field{}}
		return m, nil
	}
}

func parseMessageEnum(name string, schema *openapi3.SchemaRef) (Enum, error) {
	e := Enum{Name: name, Values: schema.Value.Enum}
	return e, nil
}

func main() {
	flag.Parse()

	input := flag.Arg(0)

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	doc, err := loader.LoadFromFile(input)

	messages := []Message{}
	enums := []Enum{}

	if err != nil {
		fmt.Printf("Parsing Error: %s\n", err)
	} else {
		//fmt.Printf("%#v\n", doc.Components.Parameters)
		for name, schema := range doc.Components.Schemas {
			if schema.Value.Enum != nil {
				if e, err := parseMessageEnum(name, schema); err == nil {
					enums = append(enums, e)
				}
			} else {
				if e, err := parseMessageSchema(name, schema); err == nil {
					messages = append(messages, e)
				}
			}
		}

		for path, req := range doc.Paths {
			if req.Get != nil {
				for _, param := range req.Get.Parameters {
					if param.Value.Schema.Value.Enum != nil {
						if m, err := parseMessageEnum(param.Value.Name, param.Value.Schema); err == nil {
							fmt.Printf("enum: %#v\n", m)
						} else {
							fmt.Printf("%s %#v\n", path, param.Value.Schema.Value)
						}
					} else {
						if m, err := parseMessageSchema(param.Value.Name, param.Value.Schema); err == nil {
							fmt.Printf("message: %#v\n", m)
						} else {
							fmt.Printf("%s %#v\n", path, param.Value.Schema.Value)
						}
					}
				}
			}
		}
	}

	sort.SliceStable(messages, func(i, j int) bool { return messages[i].Name < messages[j].Name })

	tmpl, err := template.New("proto").Parse(`
syntax = "proto3";

option go_package = "github.com/ohsu-comp-bio/funnel/tes";

package tes;

{{range $i, $enum := .enums}}
enum {{$enum.Name}} { {{range $j, $value := $enum.Values}}
	{{$value}} = {{$j}};{{end}}
}
{{end}}
{{range $i, $message := .messages}}
message {{$message.Name}} { {{range $j, $field := $message.Fields}}
	{{$field.Type}} {{$field.Name}} = {{$field.Id}};{{end}}
}
{{end}}

`)

	_ = tmpl
	if err != nil {
		fmt.Printf("Template Error: %s\n", err)
	} else {
		//tmpl.Execute(os.Stdout, map[string]interface{}{"messages": messages, "enums": enums})
	}
}
