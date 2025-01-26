package application

import (
	"liam-nak-na-api/internal/domain"
	"math"
)

type triangleService struct {
	repo domain.TriangleRepository
}

func NewTriangleService() domain.TriangleService {
	return &triangleService{}
}

func (s *triangleService) ClassifyTriangle(t domain.Triangle) (domain.TriangleClassification, error) {
	if !isValidTriangle(t) {
		return domain.TriangleClassification{
			EnglishName: "Not a triangle",
			ThaiName:    "ไม่ใช่สามเหลี่ยม",
		}, nil
	}

	if isEquilateral(t) {
		return domain.TriangleClassification{
			EnglishName: "Equilateral Triangle",
			ThaiName:    "สามเหลี่ยมด้านเท่า",
		}, nil
	}

	if isIsosceles(t) {
		return domain.TriangleClassification{
			EnglishName: "Isosceles Triangle",
			ThaiName:    "สามเหลี่ยมหน้าจั่ว",
		}, nil
	}

	if isRightAngled(t) {
		return domain.TriangleClassification{
			EnglishName: "Right Angle Triangle",
			ThaiName:    "สามเหลี่ยมมุมฉาก",
		}, nil
	}

	return domain.TriangleClassification{
		EnglishName: "Scalene Triangle",
		ThaiName:    "สามเหลี่ยมด้านไม่เท่า",
	}, nil
}

func isValidTriangle(t domain.Triangle) bool {
	return t.Height+t.Width > t.Base &&
		t.Width+t.Base > t.Height &&
		t.Base+t.Height > t.Width
}

func isEquilateral(t domain.Triangle) bool {
	return t.Height == t.Width && t.Width == t.Base
}

func isIsosceles(t domain.Triangle) bool {
	return t.Height == t.Width || t.Width == t.Base || t.Height == t.Base
}

func isRightAngled(t domain.Triangle) bool {
	sides := []float64{t.Height, t.Width, t.Base}
	for i := 0; i < 3; i++ {
		a := math.Pow(sides[i], 2)
		b := math.Pow(sides[(i+1)%3], 2)
		c := math.Pow(sides[(i+2)%3], 2)
		if math.Abs(a+b-c) < 0.0001 {
			return true
		}
	}
	return false
}
