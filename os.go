package packet

import (
	"encoding/json"
	"io/ioutil"
)

type osresponse struct {
	OperatingSystems []*os `json:"operating_systems"`
}
type os struct {
	Id              string                `json:"id"`
	Slug            string                `json:"slug"`
	Name            string                `json:"name"`
	Distro          string                `json:"distro"`
	Version         string                `json:"version"`
	ProvisionableOn []string              `json:"provisionable_on"`
	Preinstallable  bool                  `json:"preinstallable"`
	Pricing         map[string]*osPricing `json:"pricing"`
	Licensed        bool                  `json:"licensed"`
}

type osPricing struct {
	Price      float64 `json:"price"`
	Multiplier string  `json:"multiplier"`
}

func (a *API) GetOSes() ([]*os, error) {
	response, err := a.httpRequest("GET", "/operating-systems", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var oses *osresponse
	err = json.Unmarshal(body, &oses)

	return oses.OperatingSystems, err

}