package search

import "github.com/kedric/lambdarouter"

func SetupRoutes(router *lambdarouter.TreeMux) {
	router.GET("/", NotFound)
	router.GET("/search/location/cords/:lat/:long", SearchLocationByCords)
}
