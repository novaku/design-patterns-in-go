package framework

type Server interface {
	HandleRequest(string, string) (int, string)
}