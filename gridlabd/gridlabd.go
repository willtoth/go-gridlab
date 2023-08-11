package gridlabd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ID        = "id"
	CLASS     = "class"
	ISA       = "isa"
	MODULE    = "module"
	GROUPID   = "groupid"
	RANK      = "rank"
	PARENT    = "parent"
	INSVC     = "insvc"
	NAME      = "name"
	LATITUDE  = "latitude"
	LONGITUDE = "longitude"
	CLOCK     = "clock"
	OUTSVC    = "outsvc"
	FLAGS     = "flags"
)

type Gridlabd struct {
	Hostname string
	Port     int
}

func (g *Gridlabd) buildURL(api string) string {
	return fmt.Sprintf("http://%s:%d/%s", g.Hostname, g.Port, api)
}

func (g *Gridlabd) makeJSONRequest(api string) ([]byte, error) {
	resp, err := http.Get(g.buildURL(api))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gridlabd server returned status: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (g *Gridlabd) makeRequest(api string) error {
	resp, err := http.Get(g.buildURL(api))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK || resp.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	return fmt.Errorf("Gridlabd server returned status: %d", resp.StatusCode)
}

func (g *Gridlabd) PauseAt(pausetime int) error {
	var api = fmt.Sprintf("control/pauseat=%d", pausetime)
	return g.makeRequest(api)
}

func (g *Gridlabd) Pause() error {
	return g.makeRequest("control/pause")
}

func (g *Gridlabd) Resume() error {
	return g.makeRequest("control/resume")
}

func (g *Gridlabd) Stop() error {
	return g.makeRequest("control/stop")
}

func (g *Gridlabd) Shutdown() error {
	return g.makeRequest("control/shuwdown")
}

func (g *Gridlabd) Kill() error {
	return g.makeRequest("control/shuwdown")
}

type findEntry struct {
	Name string `json:"name"`
}

func Find[T dataModel](gld *Gridlabd) ([]*T, error) {
	var tmp T
	jsonData, err := gld.makeJSONRequest(fmt.Sprintf("find/%s=%s", "class", tmp.className()))
	if err != nil {
		return nil, fmt.Errorf("failed to find: %v", err)
	}

	var findResults []*findEntry
	if err := json.Unmarshal(jsonData, &findResults); err != nil {
		return nil, fmt.Errorf("failed to find: %v", err)
	}

	var results []*T

	for _, entry := range findResults {
		data, err := Fetch[T](gld, entry.Name)

		if err != nil {
			return nil, fmt.Errorf("failed to find: %v", err)
		}

		results = append(results, data)
	}

	return results, nil
}

func Fetch[T dataModel](gld *Gridlabd, name string) (*T, error) {
	jsonData, err := gld.makeJSONRequest(fmt.Sprintf("json/%s/*[dict]", name))
	if err != nil {
		return nil, fmt.Errorf("fetch failed: %e", err)
	}

	var inputArray T
	err = json.Unmarshal(jsonData, &inputArray)
	if err != nil {
		return nil, fmt.Errorf("fetch unmarshall failed: %e", err)
	}

	return &inputArray, nil
}
