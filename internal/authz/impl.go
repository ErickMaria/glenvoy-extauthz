package authz

import (
	"context"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

type ImplAuthorizationServer struct{}

// inject a header
func (a *ImplAuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {

	// Autorization logic

	access, ok := req.Attributes.Request.Http.Headers[config.AppConfig.AppKeys.Headers[0]]
	client, ok := req.Attributes.Request.Http.Headers[config.AppConfig.AppKeys.Headers[1]]

	if !ok && token != "" {
	}
	return unauthorized()
}

func authorized() (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(rpc.OK),
		},
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Header: &core.HeaderValue{
							Key:   "x-some-additional-header",
							Value: "some-value",
						},
					},
				},
			},
		},
	}, nil
}

func unauthorized() (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
			},
		},
	}, nil
}
