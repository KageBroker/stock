package context_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	internalctx "github.com/kagebroker/stock/internal/context"
	internallogger "github.com/kagebroker/stock/internal/logger"
)

func Test_Logger(t *testing.T) {
	t.Run("Add successfully a logger and retrieve it", func(t *testing.T) {
		ctx := context.Background()
		logger, _ := internallogger.NewLogger()
		ctx = internalctx.AddLoggerToContex(ctx, logger)
		assert.NotNil(t, ctx)
		loggerToTest, err := internalctx.GetLoggerFromContext(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, loggerToTest)
	})

	t.Run("Add no logger found", func(t *testing.T) {
		ctx := context.Background()
		loggerToTest, err := internalctx.GetLoggerFromContext(ctx)

		assert.Error(t, err)
		assert.Nil(t, loggerToTest)
	})
}
