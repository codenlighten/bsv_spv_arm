package contacts

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// oldAccept will accept contact request
// Accept contact godoc
// @Summary		Accept contact - Use (POST) /api/v1/invitations/{paymail} instead.
// @Description	This endpoint has been deprecated. Use (POST) /api/v1/invitations/{paymail} instead.
// @Tags		Contact
// @Produce		json
// @Param		paymail path string true "Paymail address of the contact that the user would like to accept"
// @Success		200
// @Failure		404	"Contact not found"
// @Failure		422	"Contact status not awaiting"
// @Failure		500	"Internal server error"
// @DeprecatedRouter  /v1/contact/accepted/{paymail} [patch]
// @Security	x-auth-xpub
func oldAccept(c *gin.Context, userContext *reqctx.UserContext) {
	acceptInvitations(c, userContext)
}

// acceptInvitations will accept contact request
// Accept contact invitation godoc
// @Summary		Accept contact invitation
// @Description	Accept contact invitation. For contact with status "awaiting" change status to "unconfirmed"
// @Tags		Contacts
// @Produce		json
// @Param		paymail path string true "Paymail address of the contact that the user would like to accept"
// @Success		200
// @Failure		404	"Contact not found"
// @Failure		422	"Contact status not awaiting"
// @Failure		500	"Internal server error"
// @Router		/api/v1/invitations/{paymail}/contacts [post]
// @Security	x-auth-xpub
func acceptInvitations(c *gin.Context, userContext *reqctx.UserContext) {
	paymail := c.Param("paymail")

	err := reqctx.Engine(c).AcceptContact(c, userContext.GetXPubID(), paymail)
	if err != nil {
		spverrors.ErrorResponse(c, err, reqctx.Logger(c))
		return
	}

	c.Status(http.StatusOK)
}
