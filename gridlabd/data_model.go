package gridlabd

import "github.com/mitchellh/mapstructure"

type dataModel interface {
	className() string
}

func AsDict(model dataModel) (map[string]any, error) {
	var out map[string]interface{}
	err := mapstructure.Decode(model, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
