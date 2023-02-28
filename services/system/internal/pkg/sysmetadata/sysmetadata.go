package sysmetadata

import (
	"encoding/json"
	"infra/types"
	"infra/utils"
	"reflect"
	"regexp"
	"time"

	"github.com/jinzhu/copier"
)

const (
	ColumnSysMetadataId          = "id"
	ColumnSysMetadataType        = "type"
	ColumnSysMetadataName        = "name"
	ColumnSysMetadataCode        = "code"
	ColumnSysMetadataValue       = "value"
	ColumnSysMetadataIsOnline    = "is_online"
	ColumnSysMetadataDescription = "description"
	ColumnSysMetadataMtime       = "mtime"
	ColumnSysMetadataModifierUid = "modifier_uid"
	ColumnSysMetadataIsDeleted   = "is_deleted"

	JsonSysMetadataName = "name"
	JsonSysMetadataType = "type"
	JsonSysMetadataCode = "code"
)

var (
	RegexpSysMetadataType = regexp.MustCompile(`^[\\w\\d_\\-]+$`)
	RegexpSysMetadataCode = RegexpSysMetadataType
)

// SysMetadata 系统元数据模型
type SysMetadata struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key" redis:"id"`
	Type        string    `json:"type" gorm:"column:type" redis:"type"`
	Code        string    `json:"code" gorm:"column:code" redis:"code"`
	Value       string    `json:"value" gorm:"column:value" redis:"value"`
	IsOnline    bool      `json:"isOnline" gorm:"column:is_online" redis:"isOnline"`
	Description string    `json:"description" gorm:"column:description" redis:"description"`
	Ctime       time.Time `json:"createTime" gorm:"column:ctime;autoCreateTime:milli" redis:"mtime"`
	Mtime       time.Time `json:"updateTime" gorm:"column:mtime;autoUpdateTIme:milli" redis:"ctime"`
	CreatorUid  string    `json:"creatorUid" gorm:"column:creator_uid" redis:"creatorUid"`
	ModifierUid string    `json:"modifierUid" gorm:"column:modifier_uid" redis:"modifierUid"`
	IsDeleted   bool      `json:"isDeleted" gorm:"column:is_deleted"`
}

// TableName 系统元数据表名
func (metadata *SysMetadata) TableName() string {
	return TABLE_NAME
}

type SysMetadataDataRedisHMap map[string]interface{}

func (metadata *SysMetadata) MarshalBinary() ([]byte, error) {
	return json.Marshal(metadata)
}

func (metadata *SysMetadata) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, metadata)
}

type SysMetadataList []*SysMetadata

// CodeList Code切片
func (metadataList SysMetadataList) CodeList() []string {
	res := make([]string, 0, len(metadataList))
	for _, e := range metadataList {
		res = append(res, e.Code)
	}
	return res
}

// Filter 通用筛选方法
func (metadataList SysMetadataList) Filter(f func(ptr *SysMetadata) bool) SysMetadataList {
	res := SysMetadataList{}
	for _, e := range metadataList {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

// FilterOnline 筛选上线的
func (metadataList SysMetadataList) FilterOnline() SysMetadataList {
	return metadataList.Filter(func(ptr *SysMetadata) bool {
		return ptr.IsOnline
	})
}

// PurgeDeleted 过滤掉已逻辑删除的
func (metadataList SysMetadataList) PurgeDeleted() SysMetadataList {
	return metadataList.Filter(func(ptr *SysMetadata) bool {
		return !ptr.IsDeleted
	})
}

type SysMetadataListGroupByType map[string]SysMetadataList

// GroupByType 通过Type分组
func (metadataList SysMetadataList) GroupByType() SysMetadataListGroupByType {
	res := SysMetadataListGroupByType{}
	for _, e := range metadataList {
		res[e.Type] = append(res[e.Type], e)
	}
	return res
}

type SysMetadataMappingByCode map[string]*SysMetadata

// FilterByCode 筛选映射
func (mapping SysMetadataMappingByCode) FilterByCode(codeList ...string) SysMetadataList {
	res := SysMetadataList{}
	for _, code := range codeList {
		if mapping[code] != nil {
			res = append(res, mapping[code])
		}
	}
	return res
}

// MappingByCode Code映射
func (metadataList SysMetadataList) MappingByCode() SysMetadataMappingByCode {
	res := make(map[string]*SysMetadata, len(metadataList))
	for _, e := range metadataList {
		res[e.Code] = e
	}
	return res
}

type SysMetadataColumnMappingByField map[string]string

// NewColumnMapping 获取数据库字段Map
func (metadata SysMetadata) NewColumnMapping() SysMetadataColumnMappingByField {
	mapping := SysMetadataColumnMappingByField{}
	mt := reflect.TypeOf(SysMetadata{})
	for i := 0; i < mt.NumField(); i++ {
		f := mt.Field(i)
		val := utils.ParseTagSetting(f.Tag.Get(types.FieldTagType.Gorm), ";").GetSettingValue("column", "")
		if len(val) > 0 {
			mapping[f.Name] = val
		}
	}
	return mapping
}

// Dup 拷贝
func (metadata *SysMetadata) Dup() *SysMetadata {
	res := &SysMetadata{}
	_ = copier.CopyWithOption(res, metadata, copier.Option{DeepCopy: true})
	return res
}

// IsRoot 是否为根类型
func (metadata *SysMetadata) IsRoot() bool {
	return metadata.Type == ROOT_TYPE
}
