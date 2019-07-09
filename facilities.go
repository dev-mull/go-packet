package packet

import (
	"encoding/json"
	"io/ioutil"
)

type facilitiesResponse struct {
	Facilities []*Facility `json:"facilities"`
}
type Facility struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Code     string            `json:"code"`
	Features []string          `json:"features"`
	Address  map[string]string `json:"address"`
	IPRanges []string          `json:"ip_ranges"`
}

func (a *API) GetFacilites() ([]*Facility, error) {
	response, err := a.httpRequest("GET", "/projects/"+a.Config.ProjectID+"/facilities", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var facilities *facilitiesResponse
	err = json.Unmarshal(body, &facilities)

	return facilities.Facilities, err

}
