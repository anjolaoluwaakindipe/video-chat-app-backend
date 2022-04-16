package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAllRoutes),
	fx.Provide(NewUserRouter),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewRoomRouter),
)

type AllRouters []Router

type Router interface {
	SetUp()
}

func NewAllRoutes(userRouter *UserRouter, authRouter *AuthRouter, roomRouter *RoomRouter) *AllRouters {
	return &AllRouters{
		userRouter,
		authRouter,
		roomRouter,
	}
}

func (ar *AllRouters) SetUp() {
	for _, router := range *ar {
		router.SetUp()
	}
}
