package api

import (
	"database/sql"
	"net/http"

	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreteAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var request CreteAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: request.Owner,
		Currency: request.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	Id int64 `json:"id" uri:"id" binding:"required,min=1"`	//? min simbolizes the smallest possible value

}

// GetAccount gets an account from the 'id' shown in the uri
func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound , errorResponse(err))
			return	
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest  struct {
	PageID int32 `json:"page_id" form:"page_id" binding:"required,min=1"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required,min=5,max=10"`
}

// ListAccount gets an account from the 'id' shown in the uri
func (server *Server) ListAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1)*req.PageSize, //? we divide the accounts in pages of 'Pagesize'
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
