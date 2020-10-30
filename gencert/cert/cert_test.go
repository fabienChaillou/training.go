package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")

	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}

	if c == nil {
		t.Error("Cert should be a valid reference. got=nill")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")

	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "1dc55b0ce18d75f39ff9b2364794072f8a71adca6ffd66b44b219aa446be22e237688a236114dec956b9b96eec44f2aead5793e40a0c7c94ea4"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")

	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "1dc55b0ce18d75f39ff9b2364794072f8a71adca6ffd66b44b219aa446be22e237688a236114dec956b9b96eec44f2aead5793e40a0c7c94ea4"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long name name (course=%s)", name)
	}
}
