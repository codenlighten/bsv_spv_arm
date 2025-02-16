package transactions

import (
	"net/http"

	"github.com/bitcoin-sv/go-broadcast-client/broadcast"
	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// broadcastCallback will handle a broadcastCallback call from the broadcast api
func broadcastCallback(c *gin.Context) {
	logger := reqctx.Logger(c)
	var resp *broadcast.SubmittedTx

	err := c.Bind(&resp)
	if err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, logger)
		return
	}

	err = reqctx.Engine(c).UpdateTransaction(c.Request.Context(), resp)
	if err != nil {
		logger.Err(err).Msgf("failed to update transaction - tx: %v", resp)
		spverrors.ErrorResponse(c, err, logger)
		return
	}

	// Return response
	c.Status(http.StatusOK)
}
