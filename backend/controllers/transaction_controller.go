package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akhamatvarokah/miniWallet/backend/auth"
	"github.com/akhamatvarokah/miniWallet/backend/models"
	"github.com/gin-gonic/gin"
)

type ParamTopUp struct {
	Amount uint64 `json:"amount"`
}

type ParamTransfer struct {
	Amount uint64 `json:"amount"`
	UserID uint64 `json:"user_id"`
}

func (server *Server) TopUp(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	rq := ParamTopUp{}
	err = json.Unmarshal(body, &rq)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	b, err := server.Trasaction(rq.Amount, uid, 0, "debit")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": b,
	})
}

func (server *Server) Transfer(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	rq := ParamTransfer{}
	err = json.Unmarshal(body, &rq)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user := models.User{}
	_, err = user.FindUserByID(server.DB, uint32(rq.UserID))
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	b, err := server.Trasaction(rq.Amount, uid, 0, "credit")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	_, err = server.Trasaction(rq.Amount, rq.UserID, rq.Amount, "debit")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": b,
	})
}

func (server *Server) History(c *gin.Context) {
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	fmt.Println(uid)

	ubh := models.UserBalanceHistory{}
	ubhs, err := ubh.FindAllHistoryUser(server.DB)
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": ubhs,
	})
}

func (server *Server) Trasaction(amount, userId, addBalance uint64, t string) (*models.UserBalance, error) {
	ub := models.UserBalance{}
	ub.UserId = userId
	cb, err := ub.GetBalance(server.DB)
	if err != nil {
		return nil, err
	}

	if t == "debit" {
		ub.Balance = cb.Balance + amount
	} else {
		if cb.Balance < amount {
			return nil, errors.New("Your Balance not enough")
		}

		ub.Balance = cb.Balance - amount
	}

	ub.BalanceAchive += addBalance
	balance, err := ub.UpdateBalance(server.DB)
	if err != nil {
		return nil, err
	}

	ubh := models.UserBalanceHistory{
		BalanceBefore: cb.Balance,
		BalanceAfter:  balance.Balance,
		Type:          t,
		UserBalanceId: balance.ID,
		Author:        "Admin",
		Ip:            ":1",
	}

	_, err = ubh.CreateBalanceHistory(server.DB)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
