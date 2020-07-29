package util

import (
	"errors"
	. "reflect"
	"strconv"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules) //通过make函数进行map的初始化

//注册自定义规则
func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册，无需再次注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}
func Verify(st interface{}, roleMap Rules) (err error) {

}

//长度和数值的比较，根据类型自动校验
func CompareVerify(value Value, verifyStr string) bool {
	switch value.Kind() {
	case Int, Int8, Int16, Int32, Int64:
		return Compare(value.Int(), verifyStr)
	case Float32, Float64:
		return Compare(value.Float(), verifyStr)
	case Uint, Uint8, Uint16, Uint32, Uint64:
		return Compare(value.Float(), verifyStr)
	case String, Slice, Array:
		return Compare(value.Len(), verifyStr)
	default:
		return false
	}
}

//比较具体的值
func Compare(value interface{}, str string) bool {
	arry := strings.Split(str, "=")
	Val := ValueOf(value)
	switch Val.Kind() {
	case Int, Int8, Int16, Int32, Int64:
		Vint, err := strconv.ParseInt(arry[1], 10, 64)
		if err != nil {
			return false
		}
		switch arry[0] {
		case "lt":
			return Val.Int() < Vint
		case "gt":
			return Val.Int() <= Vint
		case "eq":
			return Val.Int() == Vint
		case "ne":
			return Val.Int() != Vint
		case "gt":
			return Val.Int() > Vint
		case "ge":
			return Val.Int() >= Vint
		default:
			return false
		}
	case Uint, Uint8, Uint16, Uint32, Uint64:
		Vuint, err := strconv.ParseUint(arry[1], 10, 64)
		if err != nil {
			return false
		}
		switch arry[0] {
		case "lt":
			return Val.Uint() < Vuint
		case "gt":
			return Val.Uint() <= Vuint
		case "eq":
			return Val.Uint() == Vuint
		case "ne":
			return Val.Uint() != Vuint
		case "gt":
			return Val.Uint() > Vuint
		case "ge":
			return Val.Uint() >= Vuint
		default:
			return false
		}
	case Float32, Float64:
		VFloat, err := strconv.ParseFloat(arry[1], 64)
		if err != nil {
			return false
		}
		switch arry[0] {
		case "lt":
			return Val.Float() < VFloat
		case "gt":
			return Val.Float() <= VFloat
		case "eq":
			return Val.Float() == VFloat
		case "ne":
			return Val.Float() != VFloat
		case "gt":
			return Val.Float() > VFloat
		case "ge":
			return Val.Float() >= VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func IsNotEmpty(value Value) bool {
	switch value.Kind() {
	case String:
		return value.Len() == 0
	case Bool:
		return !value.Bool()
	case Int, Int8, Int16, Int32, Int64:
		return value.Int() == 0
	case Uint, Uint8, Uint16, Uint32, Uint64:
		return value.Uint() == 0
	case Float32, Float64:
		return value.Float() == 0
	case Interface, Ptr:
		return value.IsNil()
	}
	return DeepEqual(value.Interface(), Zero(value.Type()).Interface())
}

//不为空
func NotEmpty() string {
	return "notEmpty"
}

//小于入参（<）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Lt(str string) string {
	return "lt=" + str
}

//小于等于入参（<=）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Le(str string) string {
	return "le=" + str
}

//等于入参（=）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Eq(str string) string {
	return "eq=" + str
}

//不等于入参（!=）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Ne(str string) string {
	return "ne=" + str
}

//大于入参（>）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Gt(str string) string {
	return "gt=" + str
}

//大于等于入参（>=）,如果为string array Slice 则为长度比较，如果是int unit float 则为数值比较
func Ge(str string) string {
	return "ge=" + str
}
