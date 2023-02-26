package types

const (
	// get post
	RequestActionTypeQuery = "query"
	// delete
	RequestActionTypeRemove = "remove"
	// delete
	RequestActionTypeDelete = "delete"
	// patch
	RequestActionTypeUpdate = "update"
	// post
	RequestActionTypeCreate = "create"
	// put
	RequestActionTypeUpsert = "upsert"
	// post
	RequestActionTypeExecute = "execute"
)

// RequestActionType 请求动作类型
var RequestActionType = &requestActionType{
	Query:   RequestActionTypeQuery,
	Remove:  RequestActionTypeRemove,
	Delete:  RequestActionTypeDelete,
	Update:  RequestActionTypeUpdate,
	Create:  RequestActionTypeCreate,
	Upsert:  RequestActionTypeUpsert,
	Execute: RequestActionTypeExecute,
}

type requestActionType struct {
	Query   string
	Remove  string
	Delete  string
	Update  string
	Create  string
	Upsert  string
	Execute string
}
