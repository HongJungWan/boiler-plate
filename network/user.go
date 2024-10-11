package network

import (
	"boiler-plate/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다.")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다.")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")

	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공",
		},
		User: nil,
	})

}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다.")
}
