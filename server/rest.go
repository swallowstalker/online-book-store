package server

import (
	"github.com/julienschmidt/httprouter"
)

type Rest struct {
	*httprouter.Router
}
