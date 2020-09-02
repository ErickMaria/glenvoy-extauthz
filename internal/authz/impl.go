package authz

import (
	"context"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"
	"github/erickmaria/glooe-envoy-extauthz/internal/entity"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
	"github/erickmaria/glooe-envoy-extauthz/internal/types"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/jinzhu/gorm"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

var (
	domain = entity.Domain{}
	app    = entity.App{}
	token  = entity.Token{}
	ctx    = context.Background()
)

type ImplAuthorizationServer struct {
	DB *gorm.DB
}

// inject a header
func (a *ImplAuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {

	// capturing headers
	host := req.Attributes.Request.Http.GetHost()
	access := req.Attributes.Request.Http.Headers[config.AppConfig.Glenvoy.AppKeys.Headers[0]]
	client := req.Attributes.Request.Http.Headers[config.AppConfig.Glenvoy.AppKeys.Headers[1]]

	if autorizationLogic(ctx, a.DB, host, access, client) {
		return authorized()
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
							Key:   "authorization",
							Value: "true",
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

func autorizationLogic(ctx context.Context, db *gorm.DB, host, access, client string) bool {

	db.Find(&domain, entity.Domain{Host: host})
	if domain.Host == "" {
		logging.Logger(ctx).Infof("Domain %s not exist on database\n", domain.Name)
		return false
	}

	db.Find(&app, entity.App{Code: client, DomainID: domain.ID})
	if app.Code == "" {
		logging.Logger(ctx).Infof("Not Fould Code on App with Domain %s\n", domain.Name)
		return false
	}
	if app.Status == types.REVOKED || app.Status == types.DEACTIVATE {
		logging.Logger(ctx).Infof("App %s on database\n", app.Status)
		return false
	}

	db.Find(&token, entity.Token{Code: access, AppID: app.ID})
	if token.Code == "" {
		logging.Logger(ctx).Infof("Not Fould Code on Token with App %s\n", app.Name)
		return false
	}
	if token.Status == types.REVOKED || token.Status == types.DEACTIVATE {
		logging.Logger(ctx).Infof("Token %s on database\n", token.Status)
		return false
	}

	return true
}
