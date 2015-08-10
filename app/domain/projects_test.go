package domain

import "testing"

func Test_Projects_Init(t *testing.T) {
	projects := &ProjectList{}
	projects.init("../../public/resources/projects.yml")
	if len(projects.Items) == 0 {
		t.Error("There are no items in the project.")
	}
}
