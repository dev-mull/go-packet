package packet

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGetEvents(t *testing.T) {
	data := `{
	"events": [
		{
			"id": "72477be7-cdac-45ac-9629-9f8f5b02fa9a",
			"type": "provisioning.106",
			"body": "Operating system packages installed",
			"state": null,
			"created_at": "2019-07-05T21:22:21Z",
			"relationships": [
				{
					"href": "#497a7f3b-347d-4b99-b286-039dcb9aa406"
				}
			],
			"ip": "104.156.90.27",
			"modified_by": {
				"id": "0ff7f54a-8999-4102-8c74-f64813ebc57c",
				"full_name": "PacketBot ​",
				"avatar_url": "/users/0ff7f54a-8999-4102-8c74-f64813ebc57c/avatars/original?1519335086",
				"avatar_thumb_url": "/users/0ff7f54a-8999-4102-8c74-f64813ebc57c/avatars/thumb?1519335086"
			},
			"interpolated": "Operating system packages installed",
			"href": "/events/72477be7-cdac-45ac-9629-9f8f5b02fa9a"
		}]}`
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		io.WriteString(w, data)
	}
	a, _ := NewAPI("empty")
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	a.Url = server.URL
	a.config.Token = "test"
	a.config.ProjectID = "test"
	events, err := a.GetEvents("device", "test")
	if err != nil {
		t.Error(err)
	}
	if len(events) != 1 {
		t.Error("Unable to load events from example payload", len(events))
	}
}
