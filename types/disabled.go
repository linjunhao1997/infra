package types

const (
	DisabledTypeNotDisabled = 0
	DisabledTypeIsDisabled  = 1
)

var DisabledType = &disabledType{
	NotDisabled: DisabledTypeNotDisabled,
	IsDisabled:  DisabledTypeIsDisabled,
}

type disabledType struct {
	IsDisabled  int64
	NotDisabled int64
}
