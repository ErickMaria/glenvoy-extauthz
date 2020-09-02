package logging

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

var (
	appName string
)

type Log struct {

}

func Init(name string){
	appName = name
}

func Logger(ctx context.Context) *zap.SugaredLogger {
	return contextutils.LoggerFrom(contextutils.WithLogger(ctx, appName))
}
