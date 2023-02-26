package templates

import (
	"infra/types"
	"infra/utils"
	"io"
	"os"
	"reflect"
	"strings"
	"text/template"
)

func GenResourceGoZeroApi(tplFilePath, targetDirPath, serviceName string, datas ...interface{}) error {
	bytes, err := os.ReadFile(tplFilePath)
	if err != nil && err != io.EOF {
		return err
	}

	funcMap := template.FuncMap{}
	funcMap["toLower"] = func(str string) string {
		return strings.ToLower(str)
	}
	content := string(bytes)
	tpl, err := template.New(tplFilePath).Funcs(funcMap).Parse(content)
	if err != nil {
		return err
	}

	for _, data := range datas {
		fileName := "gateway" + "_" + strings.ToLower(strings.ReplaceAll(reflect.TypeOf(data).String(), ".", "_")+"_rest.api")
		file, err := os.OpenFile(targetDirPath+"/"+fileName, os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			return err
		}
		if err := tpl.Execute(file, NewGoZeroApiData(data, serviceName)); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}

type JsonField struct {
	Name     string
	TypeName string
	Tag      string
}

type genGoZeroApiData struct {
	ServiceName         string
	ResourceName        string
	ModuleName          string
	ReqUpdateJsonFields []*JsonField
	ReqCreateJsonFields []*JsonField
	ReqQueryJsonFields  []*JsonField
	RespQueryJsonFields []*JsonField
}

type genGoZeroProtoData struct {
	ServiceName          string
	ResourceName         string
	ReqUpdateProtoFields []*ProtoField
	ReqCreateProtoFields []*ProtoField
	ReqQueryProtoFields  []*ProtoField
	RespQueryProtoFields []*ProtoField
}

type genGoZeroProtoDatas struct {
	Datas       []*genGoZeroProtoData
	ServiceName string
}

func NewGoZeroApiData(in interface{}, serviceName string) *genGoZeroApiData {
	return &genGoZeroApiData{
		ServiceName:         serviceName,
		ResourceName:        getResourceName(in),
		ModuleName:          getModuleName(in),
		ReqUpdateJsonFields: getUpdateReqJsonFields(in),
		ReqCreateJsonFields: getCreateReqJsonFields(in),
		ReqQueryJsonFields:  getQueryReqJsonFields(in),
		RespQueryJsonFields: getQueryRespJsonFields(in),
	}
}

func NewGoZeroProtoData(in interface{}, serviceName string) *genGoZeroProtoData {
	return &genGoZeroProtoData{
		ServiceName:          serviceName,
		ResourceName:         getResourceName(in),
		ReqUpdateProtoFields: getUpdateReqProtoFields(in),
		ReqCreateProtoFields: getCreateReqProtoFields(in),
		ReqQueryProtoFields:  getQueryReqProtoFields(in),
		RespQueryProtoFields: getQueryRespProtoFields(in),
	}
}

func getResourceName(data interface{}) string {
	gType := reflect.TypeOf(data)
	return strings.Split(gType.String(), ".")[1]
}

func getModuleName(data interface{}) string {
	gType := reflect.TypeOf(data)
	return strings.Split(gType.String(), ".")[0]
}

func getCreateReqJsonFields(data interface{}) []*JsonField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*JsonField, 0, count)

	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		json := f.Tag.Get(types.FieldTagType.Json)
		if len(json) == 0 {
			json = strings.ToUpper(f.Name[:1]) + f.Name[1:]
		}
		field := &JsonField{
			Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:],
			TypeName: f.Type.String(),
			Tag:      "`" + types.FieldTagType.Json + ":" + `"` + json + `"` + "`",
		}
		if settings.Contains(types.FieldTagTypeGoZeroSettingCreateReqType) {
			field.TypeName = settings.GetSettingValue(types.FieldTagTypeGoZeroSettingCreateReqType, field.TypeName)
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingCreateReq) {
			res = append(res, field)
		}
	}

	return res
}

func getUpdateReqJsonFields(data interface{}) []*JsonField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*JsonField, 0, count)
	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		json := f.Tag.Get(types.FieldTagType.Json)
		if len(json) == 0 {
			json = strings.ToUpper(f.Name[:1]) + f.Name[1:]
		}
		fTypeName := f.Type.String()
		if f.Type.Kind() != reflect.Ptr {
			fTypeName = `*` + fTypeName
		}
		field := &JsonField{
			Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:],
			TypeName: fTypeName,
			Tag:      "`" + types.FieldTagType.Json + ":" + `"` + json + `"` + "`",
		}
		if settings.Contains(types.FieldTagTypeGoZeroSettingUpdateReqType) {
			field.TypeName = settings.GetSettingValue(types.FieldTagTypeGoZeroSettingUpdateReqType, field.TypeName)
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingUpdateReq) {
			res = append(res, field)
		}
	}

	return res
}

func getQueryReqJsonFields(data interface{}) []*JsonField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*JsonField, 0, count)
	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		fTypeName := f.Type.String()
		if f.Type.Kind() != reflect.Ptr {
			fTypeName = `*` + fTypeName
		}
		typeName := settings.GetSettingValue(types.FieldTagTypeGoZeroSettingQueryReqType, fTypeName)

		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqMulti) {
			typeName = `[]` + typeName
			field := &JsonField{
				Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
				Tag:      "`" + types.FieldTagType.Json + ":" + `"` + f.Tag.Get(types.FieldTagType.Json) + `"` + "`",
			}
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqScope) {
			start := "start" + strings.ToUpper(f.Tag.Get(types.FieldTagType.Json)[:1]) + f.Tag.Get(types.FieldTagType.Json)[1:]
			startField := &JsonField{
				Name:     "Start" + strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
				Tag:      "`" + types.FieldTagType.Json + ":" + `"` + start + `"` + "`",
			}
			end := "end" + strings.ToUpper(f.Tag.Get(types.FieldTagType.Json)[:1]) + f.Tag.Get(types.FieldTagType.Json)[1:]
			endField := &JsonField{
				Name:     "End" + strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
				Tag:      "`" + types.FieldTagType.Json + ":" + `"` + end + `"` + "`",
			}
			res = append(res, startField, endField)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReq) {
			field := &JsonField{
				Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
				Tag:      "`" + types.FieldTagType.Json + ":" + `"` + f.Tag.Get(types.FieldTagType.Json) + `"` + "`",
			}
			res = append(res, field)
		}

		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqSort) {
			desc := f.Tag.Get(types.FieldTagType.Json) + "Desc"
			field := &JsonField{
				Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:] + "Desc",
				TypeName: `bool`,
				Tag:      "`" + types.FieldTagType.Json + ":" + `"` + desc + `"` + "`",
			}
			res = append(res, field)
		}

	}

	return res
}

func getQueryRespJsonFields(data interface{}) []*JsonField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*JsonField, 0, count)
	for i := 0; i < count; i++ {
		f := gType.Field(i)
		json := f.Tag.Get(types.FieldTagType.Json)
		if len(json) == 0 {
			json = strings.ToUpper(f.Name[:1]) + f.Name[1:]
		}
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		field := &JsonField{
			Name:     strings.ToUpper(f.Name[:1]) + f.Name[1:],
			TypeName: settings.GetSettingValue(types.FieldTagTypeGoZeroSettingQueryRespType, f.Type.String()),
			Tag:      "`" + types.FieldTagType.Json + ":" + `"` + json + `"` + "`",
		}
		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryRespType) {
			field.TypeName = settings.GetSettingValue(types.FieldTagTypeGoZeroSettingQueryRespType, field.TypeName)
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryResp) {
			res = append(res, field)
		}
		if field.TypeName == "time.Time" {
			field.TypeName = "int64"
		}
	}
	return res
}

func GenResourceGoZeroProto(tplFilePath, targetDirPath, serviceName string, datas ...interface{}) error {
	bytes, err := os.ReadFile(tplFilePath)
	if err != nil && err != io.EOF {
		return err
	}

	funcMap := template.FuncMap{}
	funcMap["toLower"] = func(str string) string {
		return strings.ToLower(str)
	}
	funcMap["inc"] = func(i int, num int) int {
		return i + num
	}
	content := string(bytes)
	tpl, err := template.New(tplFilePath).Funcs(funcMap).Parse(content)
	if err != nil {
		return err
	}

	protoDatas := genGoZeroProtoDatas{
		ServiceName: serviceName,
	}
	for _, data := range datas {
		protoDatas.Datas = append(protoDatas.Datas, NewGoZeroProtoData(data, serviceName))
	}

	fileName := serviceName + "_gen.proto"
	file, err := os.OpenFile(targetDirPath+"/"+fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	if err := tpl.Execute(file, protoDatas); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

type ProtoField struct {
	Modifier string
	Name     string
	TypeName string
}

func getCreateReqProtoFields(data interface{}) []*ProtoField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*ProtoField, 0, count)

	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		field := &ProtoField{}
		if settings.Contains(types.FieldTagTypeGoZeroSettingCreateReq) {
			field.Name = strings.ToLower(f.Name[:1]) + f.Name[1:]
			field.TypeName = settings.GetSettingValue(types.FieldTagTypeGoZeroSettingCreateReqType, f.Type.String())
			res = append(res, field)
		}
	}
	return res
}

func getUpdateReqProtoFields(data interface{}) []*ProtoField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*ProtoField, 0, count)

	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		if settings.Contains(types.FieldTagTypeGoZeroSettingUpdateReq) {
			fTypeName := f.Type.String()
			typeName := settings.GetSettingValue(types.FieldTagTypeGoZeroSettingUpdateReqType, fTypeName)

			field := &ProtoField{
				Modifier: "optional",
				Name:     strings.ToLower(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
			}
			res = append(res, field)
		}
	}
	return res
}

func getQueryReqProtoFields(data interface{}) []*ProtoField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*ProtoField, 0, count)
	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		fTypeName := f.Type.String()
		typeName := settings.GetSettingValue(types.FieldTagTypeGoZeroSettingQueryReqType, fTypeName)
		typeName = strings.TrimPrefix(typeName, "*")
		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqMulti) {
			modifier := "repeated"
			field := &ProtoField{
				Modifier: modifier,
				Name:     strings.ToLower(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
			}
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqScope) {
			modifier := "optional"
			startField := &ProtoField{
				Modifier: modifier,
				Name:     "start" + strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
			}
			endField := &ProtoField{
				Modifier: modifier,
				Name:     "end" + strings.ToUpper(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
			}
			res = append(res, startField, endField)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReq) {
			modifier := "optional"
			field := &ProtoField{
				Modifier: modifier,
				Name:     strings.ToLower(f.Name[:1]) + f.Name[1:],
				TypeName: typeName,
			}
			res = append(res, field)
		}

		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryReqSort) {
			field := &ProtoField{
				Modifier: "optional",
				Name:     strings.ToLower(f.Name[:1]) + f.Name[1:] + "Desc",
				TypeName: `bool`,
			}
			res = append(res, field)
		}
	}
	return res
}

func getQueryRespProtoFields(data interface{}) []*ProtoField {
	gType := reflect.TypeOf(data)
	count := gType.NumField()
	res := make([]*ProtoField, 0, count)
	for i := 0; i < count; i++ {
		f := gType.Field(i)
		settings := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.GoZero), types.FieldTagSettingSep)
		field := &ProtoField{
			Name:     strings.ToLower(f.Name[:1]) + f.Name[1:],
			TypeName: settings.GetSettingValue(types.FieldTagTypeGoZeroSettingQueryRespType, f.Type.String()),
		}
		if settings.Contains(types.FieldTagTypeGoZeroSettingQueryRespType) {
			if strings.HasPrefix(field.TypeName, "[]*") {
				field.Modifier = "repeated"
				field.TypeName = strings.Split(field.TypeName, "[]*")[1]
			}

			if strings.HasPrefix(field.TypeName, "[]") {
				field.Modifier = "repeated"
				field.TypeName = strings.Split(field.TypeName, "[]")[1]
			}
			res = append(res, field)
		} else if settings.Contains(types.FieldTagTypeGoZeroSettingQueryResp) {
			res = append(res, field)
		}

		if field.TypeName == "time.Time" {
			field.TypeName = "int64"
		}
	}
	return res
}
