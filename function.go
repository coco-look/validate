package validate

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type CallFunc func(f *Field, args ...string) bool

/**
 * 格式化函数
 */
var formatFunc = map[string]CallFunc{
	"email":     email,
	"cn_mobile": cn_mobile,
	"url":       url,
	"safe_str":  safe_str,
}

/**
 * 表达式比较函数
 */
var expFunc = map[string]CallFunc{
	"gt":         gt,
	"gte":        gte,
	"eq":         eq,
	"lt":         lt,
	"lte":        lte,
	"empty":      empty,
	"o_interval": o_interval,
	"c_interval": c_interval,
	"in":         in,
	"eq_field":   eq_field,
}

func gt(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if compare_val, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			return f.Val.Int() > compare_val
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if compare_val, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			return f.Val.Uint() > compare_val
		}
	case reflect.Float32, reflect.Float64:
		if compare_val, err := strconv.ParseFloat(args[0], 64); err == nil {
			return f.Val.Float() > compare_val
		}
	case reflect.String:
		if compare_val, err := strconv.Atoi(args[0]); err == nil {
			return f.Val.Len() > compare_val
		}
	}
	return false
}

func gte(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if compare_val, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			return f.Val.Int() >= compare_val
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if compare_val, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			return f.Val.Uint() >= compare_val
		}
	case reflect.Float32, reflect.Float64:
		if compare_val, err := strconv.ParseFloat(args[0], 64); err == nil {
			return f.Val.Float() >= compare_val
		}
	case reflect.String:
		if compare_val, err := strconv.Atoi(args[0]); err == nil {
			return f.Val.Len() >= compare_val
		}
	}
	return false
}

/**
 * 适用数字和字符串
 */
func eq(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if compare_val, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			return f.Val.Int() == compare_val
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if compare_val, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			return f.Val.Uint() == compare_val
		}
	case reflect.Float32, reflect.Float64:
		if compare_val, err := strconv.ParseFloat(args[0], 64); err == nil {
			return f.Val.Float() == compare_val
		}
	case reflect.String:
		if compare_val, err := strconv.Atoi(args[0]); err == nil {
			return f.Val.Len() == compare_val
		}
	}
	return false
}

/**
 * 适用数字和字符串
 */
func lt(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if compare_val, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			return f.Val.Int() < compare_val
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if compare_val, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			return f.Val.Uint() < compare_val
		}
	case reflect.Float32, reflect.Float64:
		if compare_val, err := strconv.ParseFloat(args[0], 64); err == nil {
			return f.Val.Float() < compare_val
		}
	case reflect.String:
		if compare_val, err := strconv.Atoi(args[0]); err == nil {
			return f.Val.Len() < compare_val
		}
	}
	return false
}

/**
 * 适用数字和字符串
 */
func lte(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if compare_val, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			return f.Val.Int() <= compare_val
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if compare_val, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			return f.Val.Uint() <= compare_val
		}
	case reflect.Float32, reflect.Float64:
		if compare_val, err := strconv.ParseFloat(args[0], 64); err == nil {
			return f.Val.Float() <= compare_val
		}
	case reflect.String:
		if compare_val, err := strconv.Atoi(args[0]); err == nil {
			return f.Val.Len() <= compare_val
		}
	}
	return false
}

/**
 * 字符串
 */
func empty(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.String:
		if args[0] == "true" {
			return f.Val.Len() == 0
		} else {
			return f.Val.Len() > 0
		}
	}
	return false
}

/**
 * 适用数字和字符串 枚举 in=1,0  in=active,frozen
 */
func in(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare_val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return false
			}
			if f.Val.Int() == compare_val {
				return true
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare_val, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return false
			}
			if f.Val.Uint() == compare_val {
				return true
			}
		}
	case reflect.Float32, reflect.Float64:
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare_val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return false
			}
			if f.Val.Float() == compare_val {
				return true
			}
		}
	case reflect.String:
		inSlice := strings.Split(args[0], ",")
		for _, compare_val := range inSlice {
			if f.Val.String() == compare_val {
				return true
			}
		}
	}
	return false
}

/**
 * 数字：开区间 open_interval=min,max  min<val<max
 */
func o_interval(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseInt(before, 10, 64)
			a, _ := strconv.ParseInt(after, 10, 64)
			val := f.Val.Int()
			return val > b && val < a
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseUint(before, 10, 64)
			a, _ := strconv.ParseUint(after, 10, 64)
			val := f.Val.Uint()
			return val > b && val < a
		}
	case reflect.Float32, reflect.Float64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseFloat(before, 64)
			a, _ := strconv.ParseFloat(after, 64)
			val := f.Val.Float()
			return val > b && val < a
		}
	}
	return false
}

/**
 * 数字：闭区间 closed_interval=min,max  min<=val<=max
 */
func c_interval(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseInt(before, 10, 64)
			a, _ := strconv.ParseInt(after, 10, 64)
			val := f.Val.Int()
			return val >= b && val <= a
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseUint(before, 10, 64)
			a, _ := strconv.ParseUint(after, 10, 64)
			val := f.Val.Uint()
			return val >= b && val <= a
		}
	case reflect.Float32, reflect.Float64:
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseFloat(before, 64)
			a, _ := strconv.ParseFloat(after, 64)
			val := f.Val.Float()
			return val >= b && val <= a
		}
	}
	return false
}

/**
 * 格式化：邮箱
 * 适用类型：字符串
 */
func email(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 格式化：中国手机
 * 适用类型：字符串
 */
func cn_mobile(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 格式化：网址
 * 使用类型：字符串
 */
func url(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 跨字段比较
 */
func eq_field(f *Field, args ...string) bool {
	compare_val := f.RefStruct.FieldByName(args[0])
	switch f.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Val.Int() == compare_val.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return f.Val.Uint() == compare_val.Uint()
	case reflect.Float32, reflect.Float64:
		return f.Val.Float() == compare_val.Float()
	case reflect.String:
		return f.Val.String() == compare_val.String()
	}
	return false
}

/**
 * 安全的字符串
 */
func safe_str(f *Field, args ...string) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^[A-Za-z0-9_]+$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}
