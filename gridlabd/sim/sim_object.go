package sim

import "github.com/willtoth/go-gridlab/internal/glm"

type SimObject interface {
	ClassName() string
}

func GenerateGlm(so []SimObject) string {
	var result string
	for _, obj := range so {
		blockBuilder := glm.NewBlockBuilder(obj.ClassName(), glm.ObjectType).AddStruct(obj)
		result += blockBuilder.Build() + "\n"
	}
	return result
}
