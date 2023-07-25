package mocks

type MockResponse struct {
	Code int
	Obj any
}

func (m *MockResponse) JSON(code int, obj any) {
	m.Code = code
	m.Obj = obj
}