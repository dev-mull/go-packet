package packet

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGetFacilities(t *testing.T) {
	data := `{
	"facilities": [
		{
			"id": "9329c822-8e2f-4fd3-82a2-d20214548493",
			"name": "Atlanta, GA",
			"code": "atl2",
			"features": [
				"baremetal",
				"layer_2",
				"backend_transfer"
			],
			"address": {
				"href": "#32780f27-addf-42cb-b8f4-237f02213d65"
			},
			"ip_ranges": ["1.1.1.1/22"]
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
	facilities,err := a.GetFacilites()
	if err != nil {
		t.Error(err)
	}
	if len(facilities) != 1 {
		t.Error("Unable to load facilities from example payload",len(facilities))
	}
}

