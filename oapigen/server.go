package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yevhenii-babich/go-cources/oapigen/service/models"
)

type Server struct {
}

func (s Server) GetCouch(ctx echo.Context, params models.GetCouchParams) error {
	panic("implement me")
}

func (s Server) PostCouch(ctx echo.Context) error {
	panic("implement me")
}

func (s Server) DeleteCouchId(ctx echo.Context, id float32) error {
	panic("implement me")
}

func (s Server) GetCouchId(ctx echo.Context, id float32) error {
	panic("implement me")
}

func (s Server) PutCouchId(ctx echo.Context, id float32) error {
	panic("implement me")
}

func (s Server) GetStatus(ctx echo.Context) error {
	panic("implement me")
}

func (s Server) GetMIdBuzz(ctx echo.Context, mId string, params models.GetMIdBuzzParams) error {
	panic("implement me")
}

func (s Server) PostMIdBuzz(ctx echo.Context, mId string) error {
	panic("implement me")
}

func (s Server) DeleteMIdBuzzId(ctx echo.Context, mId string, id float32) error {
	panic("implement me")
}

func (s Server) GetMIdBuzzId(ctx echo.Context, mId string, id float32) error {
	panic("implement me")
}

func (s Server) PutMIdBuzzId(ctx echo.Context, mId string, id float32) error {
	panic("implement me")
}
