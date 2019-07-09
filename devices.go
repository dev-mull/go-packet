package packet

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type devicesResponse struct {
	Devices []*Device `json:"devices"`
}
type Device struct {
	Id           string    `json:"id"`
	ShortId      string    `json:"short_id"`
	Hostname     string    `json:"hostname"`
	Description  *string   `json:"description"`
	State        string    `json:"state"`
	Tags         []string  `json:"tags"`
	ImageUrl     *string   `json:"imageUrl"`
	BillingCycle string    `json:"billing_cycle"`
	User         string    `json:"user"`
	Iqn          string    `json:"iqn"`
	Locked       bool      `json:"locked"`
	OS           *os       `json:"operating_system"`
	Facility     *Facility `json:"facility"`
	Plan         *Plan     `json:"plan"`
}

type NewDevice struct {
	Facility    string  `json:"facility"`         // required
	Plan        string  `json:"plan"`             //required
	OS          string  `json:"operating_system"` //required
	Hostname    *string `json:"hostname,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (a *API) GetDevices() ([]*Device, error) {
	response, err := a.httpRequest("GET", "/projects/"+a.Config.ProjectID+"/devices", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var devices *devicesResponse
	err = json.Unmarshal(body, &devices)

	if err != nil {
		return nil, err
	}
	return devices.Devices, err

}

func (a *API) DeleteDevice(id string) error {
	if id == "" {
		return errors.New("Invalid device id")
	}
	response, err := a.httpRequest("DELETE", "/devices/"+id, nil)
	if err != nil {
		return err
	}
	//Why use 204?
	if response.StatusCode == http.StatusNoContent {
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	errs := &APIErrors{}
	json.Unmarshal(body, &errs)
	err = errors.New(strings.Join(errs.Errors, ","))
	return err
}

func (a *API) DeviceAction(id string, action string) error {
	if id == "" {
		return errors.New("Invalid device id")
	}
	response, err := a.httpRequest("POST", "/devices/"+id+"/actions?type="+action, nil)
	if err != nil {
		return err
	}
	if response.StatusCode == http.StatusAccepted {
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	errs := &APIErrors{}
	json.Unmarshal(body, &errs)
	err = errors.New(strings.Join(errs.Errors, ","))
	return err
}

func (a *API) CreateDevice(d *NewDevice) (*Device, error) {
	if d.Facility == "" {
		return nil, errors.New("Facility is missing")
	}
	if d.OS == "" {
		return nil, errors.New("OS is missing")
	}
	if d.Plan == "" {
		return nil, errors.New("Plan is missing")
	}

	txt, _ := json.Marshal(d)
	response, err := a.httpRequest("POST", "/projects/"+a.Config.ProjectID+"/devices", bytes.NewBuffer(txt))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusCreated {
		device := &Device{}
		err = json.Unmarshal(body, &device)
		return device, nil
	} else {
		errs := &APIErrors{}
		json.Unmarshal(body, &errs)
		err = errors.New(strings.Join(errs.Errors, ","))
	}

	return nil, err
}