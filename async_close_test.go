package pgconn

import (
	"context"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsyncCloseOperationError(t *testing.T) {
	t.Parallel()

	pgConn, err := Connect(context.Background(), os.Getenv("PGX_TEST_CONN_STRING"))
	require.NoError(t, err)
	defer pgConn.Close(context.Background())

	// making connection to get busy
	_ = pgConn.Exec(context.Background(), "select 1")
	// not timeout error

	pgConn.asyncClose(syscall.EPIPE)

	// connection should be in closed state
	require.Equal(t, int(pgConn.status), connStatusClosed)
}

func TestAsyncCloseTimeoutError(t *testing.T) {
	t.Parallel()

	pgConn, err := Connect(context.Background(), os.Getenv("PGX_TEST_CONN_STRING"))
	require.NoError(t, err)
	defer pgConn.Close(context.Background())

	// making connection to get busy
	_ = pgConn.Exec(context.Background(), "select 1")

	// timeout error
	pgConn.asyncClose(os.ErrDeadlineExceeded)

	// connection should be in cleanup state
	require.Equal(t, int(pgConn.status), connStatusNeedCleanup)
}
