package types

const (
	FieldTagSettingSep = ";"

	FieldTagTypeJson   = "json"
	FieldTagTypeGoZero = "gozero"
	FieldTagTypeGorm   = "gorm"

	FieldTagTypeGoZeroSettingCreateReq     = "createreq"
	FieldTagTypeGoZeroSettingCreateReqType = "createreqtype"
	FieldTagTypeGoZeroSettingUpdateReq     = "updatereq"
	FieldTagTypeGoZeroSettingQueryReq      = "queryreq"
	FieldTagTypeGoZeroSettingQueryReqType  = "queryreqtype"
	FieldTagTypeGoZeroSettingQueryReqScope = "queryreqscope"
	FieldTagTypeGoZeroSettingQueryReqMulti = "queryreqmulti"
	FieldTagTypeGoZeroSettingQueryReqSort  = "queryreqsort"
	FieldTagTypeGoZeroSettingUpdateReqType = "updatereqtype"
	FieldTagTypeGoZeroSettingQueryResp     = "queryresp"
	FieldTagTypeGoZeroSettingQueryRespType = "queryresptype"
)

var FieldTagType = &fieldTagType{
	Json:   FieldTagTypeJson,
	GoZero: FieldTagTypeGoZero,
	Gorm:   FieldTagTypeGorm,
}

type fieldTagType struct {
	Json   string
	GoZero string
	Gorm   string
}
