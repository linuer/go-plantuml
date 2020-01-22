package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClass_HasRelation(t *testing.T) {
	classA := Class{
		Name: "A",
	}
	classC := Class{
		Name: "C",
	}
	classB := Class{
		Fields: Fields{
			Field{
				Type: Type("A"),
			},
		},
	}

	assert.True(t, classB.HasRelation(classA))
	assert.False(t, classB.HasRelation(classC))
}

func TestClasses_ClassByName(t *testing.T) {
	classA := Class{Name: "A"}
	classB := Class{Name: "B"}
	classC := Class{Name: "C"}
	classes := Classes{classA, classB, classC}
	assert.Equal(t, &classA, classes.ClassByName("A"))
	assert.NotEqual(t, classB, classes.ClassByName("A"))
}

func TestClasses_ClassIndexByName(t *testing.T) {
	classA := Class{Name: "A"}
	classB := Class{Name: "B"}
	classC := Class{Name: "C"}
	classes := Classes{classA, classB, classC}
	assert.Equal(t, 2, classes.ClassIndexByName("C"))
	assert.NotEqual(t, 0, classes.ClassIndexByName("C"))
}

func TestClasses_AllPackages(t *testing.T) {
	classA := Class{Package: Package("A")}
	classB := Class{Package: Package("B")}
	classC := Class{Package: Package("C")}
	classes := Classes{classA, classB, classC}
	assert.Equal(t, Packages{Package("A"), Package("B"), Package("C")}, classes.AllPackages())
}