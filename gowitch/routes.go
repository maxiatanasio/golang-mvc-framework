package gowitch

type Route struct {
	Method     string
	Path       string
	Controller ControllerInterface
}
