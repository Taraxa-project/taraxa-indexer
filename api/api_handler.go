package api

import (
	"fmt"
	"net/http"

	"github.com/Taraxa-project/taraxa-indexer/api/models"
	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

// GetAddressDags returns all DAG blocks sent by the selected address
func (a *ApiHandler) GetAddressDags(ctx echo.Context, address models.AddressFilter, params models.GetAddressDagsParams) error {
	fmt.Println("GetAddressDags")
	var data []models.PaginatedResponse_Data_Item
	var response models.PaginatedResponse
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return nil
}

// GetAddressPbfts returns all PBFT blocks produced by the selected address
func (a *ApiHandler) GetAddressPbfts(ctx echo.Context, address models.AddressFilter, params models.GetAddressPbftsParams) error {
	fmt.Println("GetAddressPbfts")
	var data []models.PaginatedResponse_Data_Item
	var response models.PaginatedResponse
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return nil
}

// GetAddressTransactions returns all transactions from and to the selected address
func (a *ApiHandler) GetAddressTransactions(ctx echo.Context, address models.AddressFilter, params models.GetAddressTransactionsParams) error {
	fmt.Println("GetAddressTransactions")
	var data []models.PaginatedResponse_Data_Item
	var response models.PaginatedResponse
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return nil
}

// GetValidators returns all validators for the selected week and the number of PBFT blocks they produced
func (a *ApiHandler) GetValidators(ctx echo.Context, params models.GetValidatorsParams) error {
	fmt.Println("GetValidators")
	var data []models.PaginatedResponse_Data_Item
	var response models.PaginatedResponse
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return nil
}

// GetAddressDagTotal returns total number of DAG blocks sent from the selected address
func (a *ApiHandler) GetAddressDagTotal(ctx echo.Context, address models.AddressFilter) error {
	fmt.Println("GetAddressDagTotal")
	var count models.Count
	count.Total = 0
	ctx.JSON(http.StatusOK, count)
	return nil
}

// GetAddressPbftTotal returns total number of PBFT blocks produced for the selected address
func (a *ApiHandler) GetAddressPbftTotal(ctx echo.Context, address models.AddressFilter) error {
	fmt.Println("GetAddressPbftTotal")
	var count models.Count
	count.Total = 0
	ctx.JSON(http.StatusOK, count)
	return nil
}

// GetValidatorsTotal returns total number of PBFT blocks produced in selected week
func (a *ApiHandler) GetValidatorsTotal(ctx echo.Context, params models.GetValidatorsTotalParams) error {
	fmt.Println("GetValidatorsTotal")
	var count models.Count
	count.Total = 0
	ctx.JSON(http.StatusOK, count)
	return nil
}
