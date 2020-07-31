// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: qrcode/cmd/grpcapps/pob/device.proto

/*
Package pob is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pob

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_ManageDevice_ManagedDevice_0(ctx context.Context, marshaler runtime.Marshaler, client ManageDeviceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeviceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ManagedDevice(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ManageDevice_ManagedDevice_0(ctx context.Context, marshaler runtime.Marshaler, server ManageDeviceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeviceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ManagedDevice(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterManageDeviceHandlerServer registers the http handlers for service ManageDevice to "mux".
// UnaryRPC     :call ManageDeviceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterManageDeviceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ManageDeviceServer) error {

	mux.Handle("POST", pattern_ManageDevice_ManagedDevice_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ManageDevice_ManagedDevice_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ManageDevice_ManagedDevice_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterManageDeviceHandlerFromEndpoint is same as RegisterManageDeviceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterManageDeviceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterManageDeviceHandler(ctx, mux, conn)
}

// RegisterManageDeviceHandler registers the http handlers for service ManageDevice to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterManageDeviceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterManageDeviceHandlerClient(ctx, mux, NewManageDeviceClient(conn))
}

// RegisterManageDeviceHandlerClient registers the http handlers for service ManageDevice
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ManageDeviceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ManageDeviceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ManageDeviceClient" to call the correct interceptors.
func RegisterManageDeviceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ManageDeviceClient) error {

	mux.Handle("POST", pattern_ManageDevice_ManagedDevice_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ManageDevice_ManagedDevice_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ManageDevice_ManagedDevice_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ManageDevice_ManagedDevice_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "managedDevices"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_ManageDevice_ManagedDevice_0 = runtime.ForwardResponseMessage
)
