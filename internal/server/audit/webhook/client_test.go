package webhook

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.flipt.io/flipt/internal/server/audit"
	"go.uber.org/zap"
)

type dummyRetrier struct{}

func (d *dummyRetrier) RequestRetry(_ context.Context, _ []byte, _ audit.RequestCreator) error {
	return nil
}

func TestWebhookClient(t *testing.T) {
	whclient := &webhookClient{
		logger:             zap.NewNop(),
		url:                "https://flipt-webhook.io/webhook",
		maxBackoffDuration: 15 * time.Second,
		retryableClient:    &dummyRetrier{},
	}

	err := whclient.SendAudit(context.TODO(), audit.Event{
		Type:   audit.FlagType,
		Action: audit.Create,
	})

	require.NoError(t, err)
}
