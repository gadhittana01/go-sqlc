package resthttp

import (

	// injector "github.com/gadhittana01/go-sqlc/injector"

	"log"

	"github.com/go-chi/chi"
)

type RouterDependencies struct {
}

func NewRoutes(rd RouterDependencies) *chi.Mux {
	router := chi.NewRouter()

	ah, err := InitializedAuthorHandler()
	if err != nil {
		log.Println(err)
	}

	// dh, err := InitializedDummyHandler("Giri Putra Adhittana")
	// if err != nil {
	// 	log.Println(err)
	// }

	// author
	router.Get("/authors", ah.GetListAuthor)
	// router.Get("/dummy", dh.TestDummy)
	// router.Get("/activity-groups/{id}", ah.GetActivityByID)
	// router.Post("/activity-groups", ah.CreateActivity)
	// router.Patch("/activity-groups/{id}", ah.UpdateActivity)
	// router.Delete("/activity-groups/{id}", ah.DeleteActivity)

	return router
}
