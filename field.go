package validate

import (
	"fmt"
	"reflect"
)

type Field struct {
	RefStruct reflect.Value
	Name      string
	Val       reflect.Value
	Kind      reflect.Kind
	Tag       string
	State     bool
	Msg       string
}

func NewField(struct_value reflect.Value, field_name string, field_val reflect.Value, field_kind reflect.Kind, field_tag string) *Field {
	return &Field{
		RefStruct: struct_value,
		Name:      field_name,
		Val:       field_val,
		Kind:      field_kind,
		Tag:       field_tag,
		State:     false,
	}
}

//exp:[map[empty:true] map[format:email gt:3]]
func (f *Field) Parse() *Field {
	t := NewTag(f.Tag).Parse()
	exp := t.GetExp()

	for _, part := range exp {
		for k, v := range part {
			if k == "format" {
				if call, ok := formatFunc[v]; ok {
					f.State = call(f)
				}
			} else {
				if call, ok := expFunc[k]; ok {
					f.State = call(f, v)
				}
			}
			// and 条件有false就不满足
			if f.State == false {
				break
			}
		}
		// or条件有true就满足
		if f.State == true {
			break
		}
	}
	if f.State == false {
		if DebugModel {
			f.Msg = fmt.Sprintf("字段:%s 传值:%v 校验:%s", f.Name, f.Val, t.GetMsg())
		} else {
			f.Msg = fmt.Sprintf("字段:%s 校验:%s", f.Name, t.GetMsg())
		}

	}

	return f
}
