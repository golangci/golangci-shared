package events

import (
	"context"
	"os"

	"github.com/golangci/golangci-shared/pkg/runmode"
	"github.com/stvp/rollbar"
)

func trackError(ctx context.Context, err error, level string) {
	if !runmode.IsProduction() {
		return
	}

	fields := []*rollbar.Field{}

	contextProps := GetContext(ctx)
	if contextProps != nil {
		fields = append(fields, &rollbar.Field{
			Name: "props",
			Data: contextProps,
		})
	}

	rollbar.Error(level, err, fields...)
}

func init() {
	rollbar.Token = os.Getenv("ROLLBAR_API_TOKEN")
	if runmode.IsProduction() {
		rollbar.Environment = "production"
	} else {
		rollbar.Environment = "development"
	}
}
