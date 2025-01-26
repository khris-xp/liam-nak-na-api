package domain

type Triangle struct {
	Height float64
	Width  float64
	Base   float64
}

type TriangleClassification struct {
	EnglishName string
	ThaiName    string
}

type TriangleService interface {
	ClassifyTriangle(t Triangle) (TriangleClassification, error)
}

type TriangleRepository interface {
	// Add repository methods if needed
}
