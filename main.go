package main

import (
	"net/http"
	"todolist/app/infra/jwttoken"
	"todolist/app/usecase/create"
	"todolist/app/usecase/delete"
	"todolist/app/usecase/display"
	"todolist/app/usecase/update"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Run()
	server := gin.Default()

	server.POST("/api/member/todo", listCreate)
	server.DELETE("/api/member/todo/:userid/:listid", listDelete)
	server.PUT("/api/member/todo/:listid", listUpdate)
	server.GET("/api/member/todo", listDisplay)

	server.Run(":8887")
}

type Server struct {
	createUsecase  create.CreateUsecase
	deleteUsecase  delete.DeleteUsecase
	updateUsecase  update.UpdateUsecase
	displayUsecase display.DisplayUsecase
}

func listCreate(l *gin.Context) {
	// token := l.Request.Header.Get("Authorization")
	info := create.CreateInput{}
	l.BindJSON(&info)
	create := new(Server)
	create.createUsecase.Create(&info)
	// if jwttoken.AuthJWT(token, info.Userid) {

	// }
}

func listDelete(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	info := delete.DeleteInput{}
	l.BindJSON(&info)
	if jwttoken.AuthJWT(token, info.Userid) {
		delete := new(Server)
		delete.deleteUsecase.Delete(&info)
		l.JSON(http.StatusOK, gin.H{
			"userid":  info.Userid,
			"message": "listid:" + info.Listid + " " + "delete success!",
		})
	}
}

func listUpdate(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	info := update.UpdateInput{}
	l.BindJSON(&info)
	if jwttoken.AuthJWT(token, info.Userid) {
		update := new(Server)
		update.updateUsecase.Update(&info)
		l.JSON(http.StatusOK, gin.H{
			"userid":  info.Userid,
			"message": "listid:" + info.Listid + " " + "update success!",
		})
	}
}

func listDisplay(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	info := display.DisplayInput{}
	l.BindJSON(&info)
	if jwttoken.AuthJWT(token, info.Userid) {
		display := new(Server)
		output, _ := display.displayUsecase.Display(&info)
		l.JSON(http.StatusOK, output)
	}
}
