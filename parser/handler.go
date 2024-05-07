package main

import(
	// "fmt"
	"net/http"
	// "encoding/json"

	"github.com/labstack/echo/v4"
)

type ParserHandler struct{
	ethParser Parser
}

func NewParserHandler(ethParser Parser) *ParserHandler{
	return &ParserHandler{
		ethParser: ethParser,
	}
}

func(p *ParserHandler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Alive!")
}

// localhost:8080/get_current_block
func(p *ParserHandler) GetCurrentBlock(c echo.Context) error {
	// return c.JSON(http.StatusInternalServerError, err.Error())
	res := p.ethParser.GetCurrentBlock()

  	return c.JSON(http.StatusOK, res)
}

// localhost:8080/subscribe?address=0xaA247c0D81B83812e1ABf8bAB078E4540D87e3fB
func(p *ParserHandler) Subscribe(c echo.Context) error {
	address := c.QueryParam("address")

	res := p.ethParser.Subscribe(address)
	return c.JSON(http.StatusOK, res)
}

// localhost:8080/get_transactions?address=0xaA247c0D81B83812e1ABf8bAB078E4540D87e3fB
func(p *ParserHandler) GetTransactions(c echo.Context) error {
	address := c.QueryParam("address")
	res := p.ethParser.GetTransactions(address)
	return c.JSON(http.StatusOK, res)
}
