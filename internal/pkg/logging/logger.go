package logging

import (
	"context"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"

	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

func Logger(ctx context.Context) *zap.SugaredLogger {
	return contextutils.LoggerFrom(contextutils.WithLogger(ctx, config.AppConfig.App.Name))
}
