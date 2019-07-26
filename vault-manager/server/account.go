package server

import (
	"errors"

	"how-to-build-a-secure-bitcoin-wallet/vault-manager/model"

	"how-to-build-a-secure-bitcoin-wallet/vault-manager/repo"

	"github.com/gin-gonic/gin"
)

// GetAddressBalances action
func (srv *server) GetAddressBalances(c *gin.Context) {
	// get the account by the account id from the x-api-key header
	account, err := repo.GetAccountByAPIKey(c.GetHeader("x-api-key")) // you can define this in the model package
	// abort if no api key provided or account is not found
	if err != nil || account.ID == 0 {
		c.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	addressID := c.PostForm("address_id")
	// load balances from the database
	addressBalances := make([]model.AddressBalance, 0, 0)
	db := srv.repo.Conn.Where("address_id = ?", addressID).Find(&addressBalances)
	if db.Error != nil {
		c.AbortWithError(500, db.Error)
		return
	}
	// display them in an easy to manage format
	balances := make(map[string]interface{}, 0)
	for _, balance := range addressBalances {
		balances[balance.Coin] = map[string]interface{}{
			"balance":       balance.Balance,
			"lockedBalance": balance.LockedBalance,
		}
	}
	c.JSON(200, balances)
}

func (srv *server) HandleDepositTransaction(symbol, txid, amount, to string) error {
	address, err := model.GetAddressByPublicKey(to) // @todo add method in the model
	if err != nil {
		return err
	}

	// add address event
	addressEvent := model.NewAddressEvent(address.ID, "deposit", symbol, amount, txid, "", "0", symbol)
	if err := srv.repo.Conn.Create(&addressEvent).Error; err != nil {
		return err
	}

	// reflect the latest event on the balance
	// this can either be done here or in a separate section in order to make the
	// event generation and the balance update completely independent
	srv.repo.ApplyEventOnBalance(addressEvent)

	return nil
}
