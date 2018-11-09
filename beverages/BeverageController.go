package beverages

import (
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/killingspark/HaDiBar/authStuff"
	"github.com/killingspark/HaDiBar/permissions"
	"github.com/killingspark/HaDiBar/restapi"
	"github.com/killingspark/HaDiBar/settings"
)

//BeverageController : Controller for the Beverages
type BeverageController struct {
	service *BeverageService
}

//NewBeverageController creates a new BeverageController and initializes its service
func NewBeverageController(perms *permissions.Permissions) (*BeverageController, error) {
	bc := &BeverageController{}
	var err error
	bc.service, err = NewBeverageService(settings.S.DataDir, perms)
	if err != nil {
		return nil, err
	}
	return bc, nil
}

//GetBeverages responds with all existing Beverages
func (controller *BeverageController) GetBeverages(ctx *gin.Context) {
	inter, _ := ctx.Get("logininfo")
	info := inter.(*authStuff.LoginInfo)
	bevs, err := controller.service.GetBeverages(info.Name)
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt get the beverage array").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	response, err := restapi.NewOkResponse(bevs).Marshal()
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt marshal the beverage array").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//GetBeverage responds with the beverage identified by beverage/:id
func (controller *BeverageController) GetBeverage(ctx *gin.Context) {
	inter, _ := ctx.Get("logininfo")
	info := inter.(*authStuff.LoginInfo)
	ID, ok := ctx.GetQuery("id")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No ID given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	bev, err := controller.service.GetBeverage(ID, info.Name)
	if err == nil {
		response, err := restapi.NewOkResponse(bev).Marshal()
		if err != nil {
			errResp, _ := restapi.NewErrorResponse("Couldnt marshal the beverage object").Marshal()
			fmt.Fprint(ctx.Writer, string(errResp))
			ctx.Abort()
			return
		}
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Next()
	} else {
		response, _ := restapi.NewErrorResponse(err.Error()).Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}
}

//NewBeverage creates a new beverage with the given form-values "value" and "name" and returns it
func (controller *BeverageController) NewBeverage(ctx *gin.Context) {
	inter, _ := ctx.Get("logininfo")
	info := inter.(*authStuff.LoginInfo)
	nv, err := strconv.Atoi(ctx.PostForm("value"))
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Invalid value").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	nb, err := controller.service.NewBeverage(info.Name, ctx.PostForm("name"), nv)
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt save new beverage").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}
	response, err := restapi.NewOkResponse(nb).Marshal()
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt marshal the beverage object").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//UpdateBeverage updates the beverage identified by /beverage/:id with the given form-values "value" and "name" and returns it
func (controller *BeverageController) UpdateBeverage(ctx *gin.Context) {
	ID, ok := ctx.GetQuery("id")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No ID given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	nv, err := strconv.Atoi(ctx.PostForm("value"))
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Invalid value").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}
	nn := ctx.PostForm("name")

	var info *authStuff.LoginInfo
	if inter, ok := ctx.Get("logininfo"); ok {
		info, ok = inter.(*authStuff.LoginInfo)
		if !ok {
			response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Abort()
			return
		}
	}

	nb, err := controller.service.UpdateBeverage(ID, info.Name, nn, nv)
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt update beverage").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	response, err := restapi.NewOkResponse(nb).Marshal()
	if err != nil {
		errResp, _ := restapi.NewErrorResponse("Couldnt marshal the beverage object").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//UpdateAccount updates the value of the account identified by accounts/:id with the form-value "value" as diffenrence
func (controller *BeverageController) GivePermissionToUser(ctx *gin.Context) {
	ID, ok := ctx.GetQuery("id")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No ID given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	newowner := ctx.PostForm("newowner")

	var info *authStuff.LoginInfo
	if inter, ok := ctx.Get("logininfo"); ok {
		info, ok = inter.(*authStuff.LoginInfo)
		if !ok {
			response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Abort()
			return
		}
	}

	err := controller.service.GivePermissionToUser(ID, info.Name, newowner, permissions.CRUD)
	if err != nil {
		response, _ := restapi.NewErrorResponse(err.Error()).Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}
	response, _ := restapi.NewOkResponse("").Marshal()
	fmt.Fprint(ctx.Writer, string(response))
	ctx.Next()
}

//DeleteBeverage deletes the beverage identified by /beverage/:id and responds with a YEP/NOPE
func (controller *BeverageController) DeleteBeverage(ctx *gin.Context) {
	ID, ok := ctx.GetQuery("id")
	if !ok {
		errResp, _ := restapi.NewErrorResponse("No ID given").Marshal()
		fmt.Fprint(ctx.Writer, string(errResp))
		ctx.Abort()
		return
	}

	var info *authStuff.LoginInfo
	if inter, ok := ctx.Get("logininfo"); ok {
		info, ok = inter.(*authStuff.LoginInfo)
		if !ok {
			response, _ := restapi.NewErrorResponse("Something went wrong while processing the username").Marshal()
			fmt.Fprint(ctx.Writer, string(response))
			ctx.Abort()
			return
		}
	}

	if err := controller.service.DeleteBeverage(ID, info.Name); err == nil {
		response, _ := restapi.NewOkResponse("").Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Next()
	} else {
		response, _ := restapi.NewErrorResponse(err.Error()).Marshal()
		fmt.Fprint(ctx.Writer, string(response))
		ctx.Abort()
		return
	}
}
