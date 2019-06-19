package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"./util"
)

type PersonResource struct {
	// typically reference a DAO (data-access-object)
}

func (p PersonResource) getAll(req *restful.Request, resp *restful.Response) {
	var res = util.SelectAllUser()
	_ = resp.WriteEntity(res)
}


func (p PersonResource) getOne(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	var res = util.SelectOneUser(id)
	_ = resp.WriteEntity(res)
}


func (p PersonResource) Register() {
	ws := new(restful.WebService)
	ws.Path("/students").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(p.getAll))
	ws.Route(ws.GET("/{user-id}").To(p.getOne))

	restful.Add(ws)
}

func main() {
	PersonResource{}.Register()
	log.Fatal(http.ListenAndServe(":3001", nil))
}
