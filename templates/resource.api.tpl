type (
    {{ .ResourceName }}Data {
        {{- range .RespQueryJsonFields }}
        {{ .Name }} {{ .TypeName}} {{.Tag}}
        {{- end}}
    }

    Query{{ .ResourceName }}DetailReq {
        Id int64 `path:"id"`
    }

    Query{{ .ResourceName }}DetailResp {
        Code uint32 `json:"code"`
        Msg string `json:"msg"`
        Err string `json:"err"`
        Data *{{ .ResourceName }}Data `json:"data"`
    }
    
    Query{{ .ResourceName }}ListReq {
        {{- range .ReqQueryJsonFields }}
        {{ .Name }} {{ .TypeName}} {{.Tag}}
        {{- end}}
    }

    Query{{ .ResourceName }}ListResp {
        Code uint32 `json:"code"`
        Msg string `json:"msg"`
        Err string `json:"err"`
        Data []*{{ .ResourceName }}Data `json:"data"`
    }

    Query{{ .ResourceName }}PageReq {
        CurrentPage int64 `json:"currentPage"`
        PageSize int64 `json:"pageSize"`
        {{- range .ReqQueryJsonFields }}
        {{ .Name }} {{ .TypeName}} {{.Tag}}
        {{- end}}
    }

    {{ .ResourceName }}Page {
        CurrentPage int64 `json:"currentPage"`
        PageSize int64 `json:"pageSize"`
        Total int64 `json:"total"`
        List []*{{ .ResourceName }}Data `json:"list"`
    }
    Query{{ .ResourceName }}PageResp {
        Code uint32 `json:"code"`
        Msg string `json:"msg"`
        Err string `json:"err"`
        Data {{ .ResourceName }}Page `json:"data"`
    }

    Create{{ .ResourceName }}Req {
        {{- range .ReqCreateJsonFields }}
        {{ .Name }} {{ .TypeName }} {{ .Tag }}
        {{- end}}
    }


    Update{{ .ResourceName}}Req {
        Id int64 `path:"id"`
        {{- range .ReqUpdateJsonFields }}
        {{ .Name }} {{ .TypeName }} {{ .Tag }}
        {{- end}}
    }

    Upsert{{ .ResourceName }}Req {
        Id int64 `path:"id"`
    }

    Remove{{ .ResourceName }}Req {
        Id int64 `path:"id"`
    }

    Delete{{ .ResourceName }}Req {
        Id int64 `path:"id"`
    }
)

@server(
    group: {{ .ModuleName }}/{{ toLower .ResourceName }}
    prefix: {{ .ServiceName }}/api/v1
)
service gateway-api {
    @handler Query{{ .ResourceName }}Detail
    get /{{ .ModuleName }}/{{ toLower .ResourceName }}/:id (Query{{ .ResourceName }}DetailReq) returns (Query{{ .ResourceName }}DetailResp)

    @handler Query{{ .ResourceName }}List
    get /{{ .ModuleName }}/{{ toLower .ResourceName }}/list (Query{{ .ResourceName }}ListReq) returns (Query{{ .ResourceName }}ListResp)

    @handler Query{{ .ResourceName }}Page
    get /{{ .ModuleName }}/{{ toLower .ResourceName }}/page (Query{{ .ResourceName }}PageReq) returns (Query{{ .ResourceName }}PageResp)

    @handler  Create{{ .ResourceName}}
    post /{{ .ModuleName }}/{{ toLower .ResourceName }} (Create{{ .ResourceName}}Req) returns (CommonResp)

    @handler  Update{{ .ResourceName}}
    patch /{{ .ModuleName }}/{{ toLower .ResourceName }}/:id (Update{{ .ResourceName}}Req) returns (CommonResp)

    @handler Delete{{ .ResourceName}}
    delete /{{ .ModuleName }}/{{ toLower .ResourceName }}/:id (Delete{{ .ResourceName}}Req) returns (CommonResp)
}