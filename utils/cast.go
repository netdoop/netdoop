package utils

import (
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

var errorCache map[string]int
var errorCacheMutex sync.Mutex
var errorLogger *zap.Logger
var errorCacheOnce sync.Once

func dealError(err error) {
	errorCacheOnce.Do(func() {
		errorLogger = GetLogger()
		errorCache = map[string]int{}
	})
	errorCacheMutex.Lock()
	defer errorCacheMutex.Unlock()
	v, ok := errorCache[err.Error()]
	if !ok {
		v = 1
		errorLogger.Warn("read data field", zap.String("error", err.Error()))
	} else {
		v += 1
	}
	errorCache[err.Error()] = v
}

var globalTagsMap map[string]map[string]string

func registerPnTags(root string, t reflect.Type) map[string]string {
	tags := make(map[string]string)
	if t.Kind() != reflect.Struct {
		return tags
	}

	if globalTagsMap == nil {
		globalTagsMap = make(map[string]map[string]string)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("pn")
		if tag != "" {
			if parts := strings.Split(tag, ","); len(parts) > 1 {
				tag = parts[0]
			}
			tags[tag] = field.Name
		}
	}
	globalTagsMap[root] = tags
	return tags
}

func getFieldNameByPnTag(root string, t reflect.Type, pn string) string {
	if t.Kind() != reflect.Struct {
		return ""
	}
	tags, ok := globalTagsMap[root]
	if !ok {
		tags = registerPnTags(root, t)
	}
	if v, ok := tags[pn]; ok {
		return v
	}
	return pn
}

func CastRefelectValue(value string, targetType reflect.Type) reflect.Value {
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
		ptrValue := reflect.New(targetType)
		ptrValue.Elem().Set(CastRefelectValue(value, targetType))
		return ptrValue
	}

	switch targetType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, _ := cast.ToInt64E(value)
		return reflect.ValueOf(intValue).Convert(targetType)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, _ := cast.ToUint64E(value)
		return reflect.ValueOf(uintValue).Convert(targetType)
	case reflect.Float32, reflect.Float64:
		floatValue, _ := cast.ToFloat64E(value)
		return reflect.ValueOf(floatValue).Convert(targetType)
	case reflect.Bool:
		boolValue, _ := cast.ToBoolE(value)
		return reflect.ValueOf(boolValue).Convert(targetType)
	case reflect.String:
		return reflect.ValueOf(value)
	case reflect.Slice:
		elemType := targetType.Elem()
		elemValue := reflect.New(elemType).Elem()
		elemValue.Set(CastRefelectValue(value, elemType))
		sliceValue := reflect.MakeSlice(targetType, 1, 1)
		sliceValue.Index(0).Set(elemValue)
		return sliceValue
	case reflect.Map:
		mapValue := reflect.MakeMap(targetType)
		return mapValue
	case reflect.Struct:
		if targetType == reflect.TypeOf(time.Time{}) {
			timeValue, _ := cast.ToTimeE(value)
			return reflect.ValueOf(timeValue)
		} else {
			return reflect.Zero(targetType)
		}
	default:
		return reflect.Zero(targetType)
	}
}

func ReadParamList(data map[string]string, obj any) error {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	re := GetRegexp(`\.CellConfig\.([0-9]+)\.`)
	// objTypeName := objValue.Type().Name()
	for key, value := range data {
		// key = strings.ReplaceAll(key, "-", "_")
		_key := re.ReplaceAllString(key, `.CellConfig1.${1}.`)
		keys := strings.Split(_key, ".")
		if len(keys) < 2 {
			continue
		}
		// if err := setField(keys[0], objValue, keys[1:], value); err != nil {
		// 	dealError(err)
		// }
		if err := setFieldValue(objValue, keys[0], keys[1:], value); err != nil {
			dealError(err)
		}
	}
	return nil
}

func setFieldValue(v reflect.Value, root string, keys []string, value string) error {
	// GetLogger().Warn("debug", zap.String("root", root), zap.Strings("keys", keys))
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.Errorf("nil pointer: %s", root)
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return errors.Errorf("value is invalid: %s", root)
	}
	if len(keys) == 0 {
		v.Set(CastRefelectValue(value, v.Type()))
		return nil
	}
	var fieldRoot string
	var fieldValue reflect.Value
	if v.Kind() == reflect.Map {
		fieldRoot = root + ".{i}"
		keyType := v.Type().Key()
		keyValue := reflect.New(keyType).Elem()
		keyValue.Set(CastRefelectValue(keys[0], keyType))

		if v.IsNil() {
			fieldMap := reflect.MakeMap(v.Type())
			v.Set(fieldMap)
		}

		elemType := v.Type().Elem()
		fieldValue = v.MapIndex(CastRefelectValue(keys[0], keyType))
		if !fieldValue.IsValid() {
			fieldValue = reflect.New(elemType).Elem()
			fieldValue.Set(CastRefelectValue("", elemType))
		}
		if err := setFieldValue(fieldValue, fieldRoot, keys[1:], value); err != nil {
			return err
		}
		v.SetMapIndex(keyValue, fieldValue)
		return nil
	} else {
		fieldRoot = root + "." + keys[0]
		fieldName := getFieldNameByPnTag(root, v.Type(), keys[0])
		if fieldName == "" {
			GetLogger().Warn("debug!!!!!", zap.String("root", root), zap.Strings("keys", keys))
			return nil
		}
		fieldValue = v.FieldByName(fieldName)
	}
	switch fieldValue.Kind() {
	case reflect.Ptr:
		if fieldValue.IsNil() {
			if !fieldValue.CanSet() {
				return errors.Errorf("can not set ptr, %v.%v", root, strings.Join(keys, "."))
			}
			fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
		}
		return setFieldValue(fieldValue, fieldRoot, keys[1:], value)
	case reflect.Struct:
		if !fieldValue.IsValid() {
			if !fieldValue.CanSet() {
				return errors.Errorf("can not set struct, %v.%v", root, strings.Join(keys, "."))
			}
			fieldValue.Set(reflect.New(fieldValue.Type()))
		}
		return setFieldValue(fieldValue, fieldRoot, keys[1:], value)
	case reflect.Map:
		if !fieldValue.IsValid() || fieldValue.IsNil() {
			if !fieldValue.CanSet() {
				return errors.Errorf("can not set map, %v.%v", root, strings.Join(keys, "."))
			}
			fieldValue.Set(reflect.MakeMap(fieldValue.Type()))
		}
		return setFieldValue(fieldValue, fieldRoot, keys[1:], value)
	default:
		return setFieldValue(fieldValue, fieldRoot, keys[1:], value)
		// return errors.Errorf("unsupported field type: %v, %s", fieldValue.Kind(), strings.Join(keys, "."))
	}
}
