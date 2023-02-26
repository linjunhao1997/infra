package types

const (
	DeletedTypeNotDeleted = 0
	DeletedTypeIsDeleted  = 1
)

var DeletedType = &deletedType{
	NotDeleted: DeletedTypeNotDeleted,
	IsDeleted:  DeletedTypeNotDeleted,
}

type deletedType struct {
	IsDeleted  int64
	NotDeleted int64
}
