package accounts

import (
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/killingspark/HaDiBar/authStuff"
	"github.com/killingspark/HaDiBar/restapi"
	"github.com/killingspark/HaDiBar/settings"
)

//AccountController is the controller for accounts
type AccountController struct {
	service *AccountService
}

//NewAccountController creates a new AccountController and initializes the service
func NewAccountController() (*AccountController, error) {
	acC := &AccountController{}
	var err error
	acC.service, err = NewAccountService(settings.S.DataDir)
	if err != nil {
		return nil, err
	}
	return acC, nil
}

//GetAccounts gets all existing accounts
func (controller *AccountController) GetAccounts(ctx *gin.Context) {
	if inter, ok := ctx.Get("logininfo"); ok {
		info, ok := inter.(*authStuff.LoginInfo)
		if ok {
			accs, err := controller.service.GetAccounts(info.GroupID)
			if err != nil {
				response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
				fmt.Fprint(ctx.Writer, string(response))
				ctx.Abort()
				return
			}
			response, _ := restapi.NewOkResponse(accs).Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Next()
		} else {
			response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Abort()
			return
		}
	} else {
		response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}
}

//GetAccount returns the account identified by account/:id
func (controller *AccountController) GetAccount(ctx *gin.Context) {
	ID, ok := ctx.GetQuery("id")
	if !ok {
		response, _ := restapi.NewErrorResponse("no id in path").Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}

	acc, err := controller.service.GetAccount(ID)
	if err != nil {
		response, _ := restapi.NewErrorResponse("Error getting account: " + err.Error()).Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}

	response, _ := restapi.NewOkResponse(acc).Marshal()
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//UpdateAccount updates the value of the account identified by accounts/:id with the form-value "value" as diffenrence
func (controller *AccountController) UpdateAccount(ctx *gin.Context) {
	ID, ok := ctx.GetQuery("id")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No ID given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	value, err := strconv.Atoi(ctx.PostForm("value"))
	if err != nil {
		response, _ := restapi.NewErrorResponse("No valid diff value").Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}

	acc, err := controller.service.UpdateAccount(ID, value)
	if err != nil {
		response, _ := restapi.NewErrorResponse(err.Error()).Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}
	response, _ := restapi.NewOkResponse(acc).Marshal()
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//UpdateAccount updates the value of the account identified by accounts/:id with the form-value "value" as diffenrence
func (controller *AccountController) NewAccount(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No name given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	if inter, ok := ctx.Get("logininfo"); ok {
		info, ok := inter.(*authStuff.LoginInfo)
		if ok {
			acc, err := controller.service.CreateAdd(name, info.GroupID)
			if err != nil {
				response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
				fmt.Fprint(ctx.Writer, string(response))
				ctx.Abort()
				return
			}
			response, _ := restapi.NewOkResponse(acc).Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Next()
		} else {
			response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Abort()
			return
		}
	}
}
