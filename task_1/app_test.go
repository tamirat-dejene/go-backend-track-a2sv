package task1

import (
	"strings"
	"testing"
)

func TestCalculateGPA(t *testing.T) {
	courses := map[Course]float32{
		{Name: "Math", Weight: 3}:  90,
		{Name: "English", Weight: 2}: 80,
	}
	expected := float32((90*3 + 80*2) / (3 + 2))
	got := CalculateGPA(courses)
	if got != expected {
		t.Errorf("Expected GPA %.2f, got %.2f", expected, got)
	}
}

func TestCalculateGPAZeroWeight(t *testing.T) {
	courses := map[Course]float32{
		{Name: "Physics", Weight: 0}: 100,
	}
	expected := float32(0)
	got := CalculateGPA(courses)
	if got != expected {
		t.Errorf("Expected GPA %.2f for zero weight, got %.2f", expected, got)
	}
}

func TestCourseString(t *testing.T) {
	c := Course{Name: "Chemistry", Weight: 2.5}
	expected := "Chemistry (Weight: 2.50)"
	if c.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, c.String())
	}
}

func TestStudentString(t *testing.T) {
	s := Student{
		Name: "Tamirat",
		Grades: map[Course]float32{
			{Name: "Math", Weight: 3}: 95,
		},
	}
	out := s.String()
	if !strings.Contains(out, "Tamirat") || !strings.Contains(out, "Math") || !strings.Contains(out, "95.00") {
		t.Errorf("Student String() missing expected content. Got:\n%s", out)
	}
}
