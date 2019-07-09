package packet

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGetPlans(t *testing.T) {
	data := `{
	"plans": [
		{
			"id": "d2829bc4-4c00-461c-9125-ec8cefb609f5",
			"slug": "baremetal_2a4",
			"name": "Type 2A4",
			"description": "Our Type 2A4 configuration is a 56-core dual socket ARM server.",
			"line": "baremetal",
			"specs": {
				"cpus": [
					{
						"count": 2,
						"type": "28-core ThunderX2 @ 2.0Ghz"
					}
				],
				"memory": {
					"total": "256GB"
				},
				"drives": [
					{
						"count": 1,
						"size": "240GB",
						"type": "SSD"
					}
				],
				"nics": [
					{
						"count": 2,
						"type": "10Gbps"
					}
				],
				"features": {}
			},
			"legacy": true,
			"deployment_types": [
				"on_demand",
				"spot_market"
			],
			"available_in": [
				{
					"href": "/facilities/e1e9c52e-a0bc-4117-b996-0fc94843ea09"
				}
			],
			"class": "baremetal_2a4",
			"pricing": {
				"hour": 0.5
			}
		}]}`
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		io.WriteString(w, data)
	}
	a,_ := NewAPI("empty")
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	a.Url = server.URL
	a.config.Token = "test"
	a.config.ProjectID = "test"
	plans,err := a.GetPlans()
	if err != nil {
		t.Error(err)
	}
	if len(plans) != 1 {
		t.Error("Unable to load plans from example payload",len(plans))
	}
}
