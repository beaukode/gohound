package backend

import "github.com/beaukode/gohound/app"

// Interface for backends
type Interface interface {
	GetNextTodo(count int) ([]app.ProbeInfo, error)
	Update(probe app.ProbeInfo)
}
