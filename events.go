package packet


import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//GetEvents Get events for a project or device
func (a *API) GetEvents(eventType string, id string) ([]*Event, error) {
	var url string
	switch eventType {
	case "project":
		url = fmt.Sprintf("/projects/%s/events", a.Config.ProjectID)
	case "device":
		url = fmt.Sprintf("/devices/%s/events", id)
	default:
		return nil, errors.New("I dont know how to get events for " + eventType)
	}
	response, err := a.httpRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		var events *eventBody
		err = json.Unmarshal(body, &events)
		if err != nil {
			return nil, err
		}
		if events == nil {
			return nil,err
		}
		return events.Events, err

	}
	errs := &APIErrors{}
	json.Unmarshal(body, &errs)
	err = errors.New(strings.Join(errs.Errors, ","))
	return nil, err
}

type eventBody struct {
	Events []*Event `json:"events"`
}

type Event struct {
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"created_at"`
	Body         string    `json:"body"`
	Interpolated string    `json:"interpolated"`
}