package packet

import (
	"encoding/json"
	"io/ioutil"
)

type planResponse struct {
	Plans []*Plan `json:"plans"`
}

type Plan struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Line        string `json:"line"`
}
type planSpecs struct {
	Cpus            []*planSpecsMisc   `json:"specs"`
	Memory          *planSpecMemory    `json:"memory"` // could have been a map,but struct is more memory efficient
	Drives          []*planSpecsMisc   `json:"drives"`
	Nics            []*planSpecsMisc   `json:"nics"`
	Features        map[string]bool    `json:"features"`
	Legacy          bool               `json:"legacy"`
	DeploymentTypes []string           `json:"deployment_types"`
	Class           string             `json:"class"`
	Pricing         map[string]float64 `json:"pricing"`
}

type planSpecsMisc struct {
	Count int    `json:"count"`
	Size  string `json:"size"`
	Type  string `json:"type"`
}

type planSpecMemory struct {
	Total string `json:"total"`
}

func (a *API) GetPlans() ([]*Plan, error) {
	response, err := a.httpRequest("GET", "/projects/"+a.Config.ProjectID+"/plans", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var plans *planResponse
	err = json.Unmarshal(body, &plans)

	return plans.Plans, err

}
