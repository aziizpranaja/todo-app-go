package main

import (
	"todo-app-go/initializers"
	"todo-app-go/routes"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main(){
	routes.AllRoute()
}