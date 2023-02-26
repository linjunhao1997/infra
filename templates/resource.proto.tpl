syntax = "proto3";

package {{ .ServiceName }}.v1;

option go_package = "./v1";

import "google/protobuf/any.proto";


message	CommonIdData {
	int64 id = 1;
}

message CommonPageData {
    int64 currentPage = 1;
    int64 pageSize = 2;
    int64 total = 3;
    repeated google.protobuf.Any list = 4;
}

message CommonIdDataResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    CommonIdData data = 4;
 }

message CommonResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    google.protobuf.Any data = 4;
 }

message CommonPageResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    CommonPageData data = 4;
}

{{- range .Datas}}
message {{ .ResourceName }}Data {
    {{- range $i, $v := .RespQueryProtoFields }}
    {{ .TypeName}} {{ .Name }} = {{ inc $i 1}};
    {{- end}}
}

message Query{{ .ResourceName }}DetailReq {
    int64 id = 1;
}

message Query{{ .ResourceName }}DetailResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    {{ .ResourceName }}Data data = 4;
}

message Query{{ .ResourceName }}ListReq {
    {{- range $i, $v := .ReqQueryProtoFields }}
    {{ .Modifier }} {{ .TypeName}} {{ .Name }} = {{ inc $i 1 }};
    {{- end}}
}

message Query{{ .ResourceName }}ListResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    repeated {{ .ResourceName }}Data data = 4;
}

message Query{{ .ResourceName }}PageReq {
    int64 currentPage = 1;
    int64 pageSize = 2;
    {{- range $i, $v := .ReqQueryProtoFields }}
    {{ .Modifier }} {{ .TypeName}} {{ .Name }} = {{ inc $i 3 }};
    {{- end}}
}

message {{ .ResourceName }}Page {
    int64 currentPage = 1;
    int64 pageSize = 2;
    int64 total = 3;
    repeated {{ .ResourceName }}Data list = 4;
}

message Query{{ .ResourceName }}PageResp {
    uint32 code = 1;
    string msg = 2;
    string err = 3;
    {{ .ResourceName }}Page data = 4;
}

message Create{{ .ResourceName }}Req {
    {{- range $i, $v := .ReqCreateProtoFields }}
    {{ .TypeName}} {{ .Name }} = {{ inc $i 1 }};
    {{- end}}
}


message Update{{ .ResourceName}}Req {
    int64 id = 1;
    {{- range $i, $v := .ReqUpdateProtoFields }}
    {{ .Modifier }} {{ .TypeName}} {{ .Name }} = {{ inc $i 2 }};
    {{- end}}
}

message Upsert{{ .ResourceName }}Req {
    int64 id = 1;
}

message Remove{{ .ResourceName }}Req {
    int64 id = 1;
}

message Delete{{ .ResourceName }}Req {
    int64 id = 1;
}
{{- end}}

service {{.ServiceName}}Service {
{{- range .Datas}}
    rpc Query{{ .ResourceName }}Detail(Query{{ .ResourceName }}DetailReq) returns (Query{{ .ResourceName }}DetailResp);
    
    rpc Query{{ .ResourceName }}List(Query{{ .ResourceName }}ListReq) returns (Query{{ .ResourceName }}ListResp);

    rpc Query{{ .ResourceName }}Page(Query{{ .ResourceName }}PageReq) returns (Query{{ .ResourceName }}PageResp);

    rpc Create{{ .ResourceName}}(Create{{ .ResourceName}}Req) returns (CommonIdDataResp);

    rpc Update{{ .ResourceName}}(Update{{ .ResourceName}}Req) returns (CommonResp);

    rpc Delete{{ .ResourceName}}(Delete{{ .ResourceName}}Req) returns (CommonResp);

{{ end }}
}

