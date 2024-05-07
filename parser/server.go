package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type server struct{
	ethParser    Parser
	e        *echo.Echo
}

func NewServer(ethParser Parser) *server{
	return &server{
		ethParser: ethParser,
		e: echo.New(),
	}
}

func(s *server) Start(){
	fmt.Println("Starting Server.....")

	parserHandler := NewParserHandler(s.ethParser)

	s.e.GET("/health_check", parserHandler.HealthCheck)
	s.e.GET("/get_current_block", parserHandler.GetCurrentBlock)
	s.e.POST("/subscribe", parserHandler.Subscribe)
	s.e.GET("/get_transactions", parserHandler.GetTransactions)

  	s.e.Logger.Fatal(s.e.Start(":8080"))
}


