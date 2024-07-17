// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.0--rc1
// source: n2x/protobuf/rpc/v1/opsAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "n2x.dev/x-api-go/grpc/common/empty"
	ops "n2x.dev/x-api-go/grpc/resources/ops"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OpsAPI_CreateProject_FullMethodName  = "/api.OpsAPI/CreateProject"
	OpsAPI_ListProjects_FullMethodName   = "/api.OpsAPI/ListProjects"
	OpsAPI_GetProject_FullMethodName     = "/api.OpsAPI/GetProject"
	OpsAPI_SetProject_FullMethodName     = "/api.OpsAPI/SetProject"
	OpsAPI_DeleteProject_FullMethodName  = "/api.OpsAPI/DeleteProject"
	OpsAPI_CreateWorkflow_FullMethodName = "/api.OpsAPI/CreateWorkflow"
	OpsAPI_ListWorkflows_FullMethodName  = "/api.OpsAPI/ListWorkflows"
	OpsAPI_GetWorkflow_FullMethodName    = "/api.OpsAPI/GetWorkflow"
	OpsAPI_SetWorkflow_FullMethodName    = "/api.OpsAPI/SetWorkflow"
	OpsAPI_DeleteWorkflow_FullMethodName = "/api.OpsAPI/DeleteWorkflow"
	OpsAPI_ActionWorkflow_FullMethodName = "/api.OpsAPI/ActionWorkflow"
	OpsAPI_ListTaskLogs_FullMethodName   = "/api.OpsAPI/ListTaskLogs"
	OpsAPI_GetTaskLog_FullMethodName     = "/api.OpsAPI/GetTaskLog"
	OpsAPI_DeleteTaskLog_FullMethodName  = "/api.OpsAPI/DeleteTaskLog"
)

// OpsAPIClient is the client API for OpsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [o-api] OpsAPI Definition: Ops Resources
type OpsAPIClient interface {
	// projects
	CreateProject(ctx context.Context, in *ops.NewProjectRequest, opts ...grpc.CallOption) (*ops.Project, error)
	ListProjects(ctx context.Context, in *ops.ListProjectsRequest, opts ...grpc.CallOption) (*ops.Projects, error)
	GetProject(ctx context.Context, in *ops.ProjectReq, opts ...grpc.CallOption) (*ops.Project, error)
	SetProject(ctx context.Context, in *ops.Project, opts ...grpc.CallOption) (*ops.Project, error)
	DeleteProject(ctx context.Context, in *ops.ProjectReq, opts ...grpc.CallOption) (*empty.Response, error)
	// workflows
	CreateWorkflow(ctx context.Context, in *ops.Workflow, opts ...grpc.CallOption) (*ops.Workflow, error)
	ListWorkflows(ctx context.Context, in *ops.ListWorkflowsRequest, opts ...grpc.CallOption) (*ops.Workflows, error)
	GetWorkflow(ctx context.Context, in *ops.WorkflowReq, opts ...grpc.CallOption) (*ops.Workflow, error)
	SetWorkflow(ctx context.Context, in *ops.Workflow, opts ...grpc.CallOption) (*ops.Workflow, error)
	DeleteWorkflow(ctx context.Context, in *ops.WorkflowReq, opts ...grpc.CallOption) (*empty.Response, error)
	ActionWorkflow(ctx context.Context, in *ops.WorkflowAction, opts ...grpc.CallOption) (*empty.Response, error)
	// taskLogs
	ListTaskLogs(ctx context.Context, in *ops.ListTaskLogsRequest, opts ...grpc.CallOption) (*ops.TaskLogs, error)
	GetTaskLog(ctx context.Context, in *ops.TaskLogReq, opts ...grpc.CallOption) (*ops.TaskLog, error)
	DeleteTaskLog(ctx context.Context, in *ops.TaskLogReq, opts ...grpc.CallOption) (*empty.Response, error)
}

type opsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOpsAPIClient(cc grpc.ClientConnInterface) OpsAPIClient {
	return &opsAPIClient{cc}
}

func (c *opsAPIClient) CreateProject(ctx context.Context, in *ops.NewProjectRequest, opts ...grpc.CallOption) (*ops.Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Project)
	err := c.cc.Invoke(ctx, OpsAPI_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) ListProjects(ctx context.Context, in *ops.ListProjectsRequest, opts ...grpc.CallOption) (*ops.Projects, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Projects)
	err := c.cc.Invoke(ctx, OpsAPI_ListProjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) GetProject(ctx context.Context, in *ops.ProjectReq, opts ...grpc.CallOption) (*ops.Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Project)
	err := c.cc.Invoke(ctx, OpsAPI_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) SetProject(ctx context.Context, in *ops.Project, opts ...grpc.CallOption) (*ops.Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Project)
	err := c.cc.Invoke(ctx, OpsAPI_SetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) DeleteProject(ctx context.Context, in *ops.ProjectReq, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, OpsAPI_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) CreateWorkflow(ctx context.Context, in *ops.Workflow, opts ...grpc.CallOption) (*ops.Workflow, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Workflow)
	err := c.cc.Invoke(ctx, OpsAPI_CreateWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) ListWorkflows(ctx context.Context, in *ops.ListWorkflowsRequest, opts ...grpc.CallOption) (*ops.Workflows, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Workflows)
	err := c.cc.Invoke(ctx, OpsAPI_ListWorkflows_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) GetWorkflow(ctx context.Context, in *ops.WorkflowReq, opts ...grpc.CallOption) (*ops.Workflow, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Workflow)
	err := c.cc.Invoke(ctx, OpsAPI_GetWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) SetWorkflow(ctx context.Context, in *ops.Workflow, opts ...grpc.CallOption) (*ops.Workflow, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.Workflow)
	err := c.cc.Invoke(ctx, OpsAPI_SetWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) DeleteWorkflow(ctx context.Context, in *ops.WorkflowReq, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, OpsAPI_DeleteWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) ActionWorkflow(ctx context.Context, in *ops.WorkflowAction, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, OpsAPI_ActionWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) ListTaskLogs(ctx context.Context, in *ops.ListTaskLogsRequest, opts ...grpc.CallOption) (*ops.TaskLogs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.TaskLogs)
	err := c.cc.Invoke(ctx, OpsAPI_ListTaskLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) GetTaskLog(ctx context.Context, in *ops.TaskLogReq, opts ...grpc.CallOption) (*ops.TaskLog, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ops.TaskLog)
	err := c.cc.Invoke(ctx, OpsAPI_GetTaskLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsAPIClient) DeleteTaskLog(ctx context.Context, in *ops.TaskLogReq, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, OpsAPI_DeleteTaskLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpsAPIServer is the server API for OpsAPI service.
// All implementations must embed UnimplementedOpsAPIServer
// for forward compatibility
//
// [o-api] OpsAPI Definition: Ops Resources
type OpsAPIServer interface {
	// projects
	CreateProject(context.Context, *ops.NewProjectRequest) (*ops.Project, error)
	ListProjects(context.Context, *ops.ListProjectsRequest) (*ops.Projects, error)
	GetProject(context.Context, *ops.ProjectReq) (*ops.Project, error)
	SetProject(context.Context, *ops.Project) (*ops.Project, error)
	DeleteProject(context.Context, *ops.ProjectReq) (*empty.Response, error)
	// workflows
	CreateWorkflow(context.Context, *ops.Workflow) (*ops.Workflow, error)
	ListWorkflows(context.Context, *ops.ListWorkflowsRequest) (*ops.Workflows, error)
	GetWorkflow(context.Context, *ops.WorkflowReq) (*ops.Workflow, error)
	SetWorkflow(context.Context, *ops.Workflow) (*ops.Workflow, error)
	DeleteWorkflow(context.Context, *ops.WorkflowReq) (*empty.Response, error)
	ActionWorkflow(context.Context, *ops.WorkflowAction) (*empty.Response, error)
	// taskLogs
	ListTaskLogs(context.Context, *ops.ListTaskLogsRequest) (*ops.TaskLogs, error)
	GetTaskLog(context.Context, *ops.TaskLogReq) (*ops.TaskLog, error)
	DeleteTaskLog(context.Context, *ops.TaskLogReq) (*empty.Response, error)
	mustEmbedUnimplementedOpsAPIServer()
}

// UnimplementedOpsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedOpsAPIServer struct {
}

func (UnimplementedOpsAPIServer) CreateProject(context.Context, *ops.NewProjectRequest) (*ops.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedOpsAPIServer) ListProjects(context.Context, *ops.ListProjectsRequest) (*ops.Projects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedOpsAPIServer) GetProject(context.Context, *ops.ProjectReq) (*ops.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedOpsAPIServer) SetProject(context.Context, *ops.Project) (*ops.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProject not implemented")
}
func (UnimplementedOpsAPIServer) DeleteProject(context.Context, *ops.ProjectReq) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedOpsAPIServer) CreateWorkflow(context.Context, *ops.Workflow) (*ops.Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (UnimplementedOpsAPIServer) ListWorkflows(context.Context, *ops.ListWorkflowsRequest) (*ops.Workflows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflows not implemented")
}
func (UnimplementedOpsAPIServer) GetWorkflow(context.Context, *ops.WorkflowReq) (*ops.Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflow not implemented")
}
func (UnimplementedOpsAPIServer) SetWorkflow(context.Context, *ops.Workflow) (*ops.Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWorkflow not implemented")
}
func (UnimplementedOpsAPIServer) DeleteWorkflow(context.Context, *ops.WorkflowReq) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedOpsAPIServer) ActionWorkflow(context.Context, *ops.WorkflowAction) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionWorkflow not implemented")
}
func (UnimplementedOpsAPIServer) ListTaskLogs(context.Context, *ops.ListTaskLogsRequest) (*ops.TaskLogs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskLogs not implemented")
}
func (UnimplementedOpsAPIServer) GetTaskLog(context.Context, *ops.TaskLogReq) (*ops.TaskLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLog not implemented")
}
func (UnimplementedOpsAPIServer) DeleteTaskLog(context.Context, *ops.TaskLogReq) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskLog not implemented")
}
func (UnimplementedOpsAPIServer) mustEmbedUnimplementedOpsAPIServer() {}

// UnsafeOpsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpsAPIServer will
// result in compilation errors.
type UnsafeOpsAPIServer interface {
	mustEmbedUnimplementedOpsAPIServer()
}

func RegisterOpsAPIServer(s grpc.ServiceRegistrar, srv OpsAPIServer) {
	s.RegisterService(&OpsAPI_ServiceDesc, srv)
}

func _OpsAPI_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.NewProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).CreateProject(ctx, req.(*ops.NewProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).ListProjects(ctx, req.(*ops.ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).GetProject(ctx, req.(*ops.ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_SetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).SetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_SetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).SetProject(ctx, req.(*ops.Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).DeleteProject(ctx, req.(*ops.ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.Workflow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_CreateWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).CreateWorkflow(ctx, req.(*ops.Workflow))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.ListWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_ListWorkflows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).ListWorkflows(ctx, req.(*ops.ListWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_GetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.WorkflowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).GetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_GetWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).GetWorkflow(ctx, req.(*ops.WorkflowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_SetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.Workflow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).SetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_SetWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).SetWorkflow(ctx, req.(*ops.Workflow))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.WorkflowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_DeleteWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).DeleteWorkflow(ctx, req.(*ops.WorkflowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_ActionWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.WorkflowAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).ActionWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_ActionWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).ActionWorkflow(ctx, req.(*ops.WorkflowAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_ListTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.ListTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).ListTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_ListTaskLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).ListTaskLogs(ctx, req.(*ops.ListTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_GetTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.TaskLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).GetTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_GetTaskLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).GetTaskLog(ctx, req.(*ops.TaskLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsAPI_DeleteTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ops.TaskLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsAPIServer).DeleteTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpsAPI_DeleteTaskLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsAPIServer).DeleteTaskLog(ctx, req.(*ops.TaskLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OpsAPI_ServiceDesc is the grpc.ServiceDesc for OpsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.OpsAPI",
	HandlerType: (*OpsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _OpsAPI_CreateProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _OpsAPI_ListProjects_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _OpsAPI_GetProject_Handler,
		},
		{
			MethodName: "SetProject",
			Handler:    _OpsAPI_SetProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _OpsAPI_DeleteProject_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _OpsAPI_CreateWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflows",
			Handler:    _OpsAPI_ListWorkflows_Handler,
		},
		{
			MethodName: "GetWorkflow",
			Handler:    _OpsAPI_GetWorkflow_Handler,
		},
		{
			MethodName: "SetWorkflow",
			Handler:    _OpsAPI_SetWorkflow_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _OpsAPI_DeleteWorkflow_Handler,
		},
		{
			MethodName: "ActionWorkflow",
			Handler:    _OpsAPI_ActionWorkflow_Handler,
		},
		{
			MethodName: "ListTaskLogs",
			Handler:    _OpsAPI_ListTaskLogs_Handler,
		},
		{
			MethodName: "GetTaskLog",
			Handler:    _OpsAPI_GetTaskLog_Handler,
		},
		{
			MethodName: "DeleteTaskLog",
			Handler:    _OpsAPI_DeleteTaskLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n2x/protobuf/rpc/v1/opsAPI.proto",
}
