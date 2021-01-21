package rest

type DataObject map[string]interface{}

type ApiResponse struct {
	StatusCode int
	Data       DataObject
}
