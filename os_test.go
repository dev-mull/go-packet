package packet

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGetOSes(t *testing.T) {
	data := `{
	"operating_systems": [
		{
			"id": "d6901f66-19ee-47cf-8348-bb5379403eb9",
			"slug": "flatcar_stable",
			"name": "Flatcar Linux - Stable",
			"distro": "flatcar",
			"version": "stable",
			"provisionable_on": [
				"c1.small.x86",
				"baremetal_1",
				"c1.xlarge.x86",
				"baremetal_3",
				"c2.medium.x86",
				"m1.xlarge.x86",
				"baremetal_2",
				"m2.xlarge.x86",
				"s1.large.x86",
				"baremetal_s",
				"t1.small.x86",
				"baremetal_0",
				"x1.small.x86",
				"baremetal_1e"
			],
			"preinstallable": false,
			"pricing": {},
			"licensed": false
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
	oses,err := a.GetOSes()
	if err != nil {
		t.Error(err)
	}
	if len(oses) != 1 {
		t.Error("Unable to load oses from example payload",len(oses))
	}
}
