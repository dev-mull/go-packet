package packet

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGetDevices(t *testing.T) {
data := `{
	"devices": [
		{
			"id": "519e59c4-1cf7-41a3-98d3-85d48eb9bdcd",
			"short_id": "519e59c4",
			"hostname": "test1234",
			"description": null,
			"state": "active",
			"tags": [],
			"image_url": null,
			"billing_cycle": "hourly",
			"user": "core",
			"iqn": "iqn.2019-07.net.packet:device.519e59c4",
			"locked": false,
			"bonding_mode": 5,
			"created_at": "2019-07-04T19:16:15Z",
			"updated_at": "2019-07-04T19:20:54Z",
			"ipxe_script_url": null,
			"always_pxe": false,
			"storage": null,
			"customdata": {},
			"operating_system": {
				"id": "d61c3912-8422-4daf-835e-854efa0062e4",
				"slug": "coreos_stable",
				"name": "Container Linux - Stable",
				"distro": "coreos",
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
					"baremetal_0"
				],
				"preinstallable": false,
				"pricing": {},
				"licensed": false
			},
			"facility": {
				"id": "e1e9c52e-a0bc-4117-b996-0fc94843ea09",
				"name": "Parsippany, NJ",
				"code": "ewr1",
				"features": [
					"baremetal",
					"storage",
					"backend_transfer",
					"layer_2",
					"global_ipv4"
				],
				"address": null,
				"ip_ranges": [
					"2604:1380::/36",
					"147.75.192.0/21",
					"147.75.64.0/22",
					"147.75.72.0/22",
					"147.75.76.0/22",
					"147.75.104.0/22",
					"147.75.96.0/20"
				]
			},
			"project": {
				"href": "/projects/ca73364c-6023-4935-9137-2132e73c20b4"
			},
			"ssh_keys": [
				{
					"href": "/ssh-keys/test"
				}
			],
			"project_lite": {
				"href": "/projects/ca73364c-6023-4935-9137-2132e73c20b4"
			},
			"volumes": [],
			"ip_addresses": [
				{
					"id": "5d475f9c-bb56-455c-9347-5d7fcfc93e60",
					"address_family": 4,
					"netmask": "255.255.255.254",
					"created_at": "2019-07-04T19:16:17Z",
					"details": null,
					"public": true,
					"cidr": 31,
					"management": true,
					"manageable": true,
					"enabled": true,
					"global_ip": null,
					"customdata": {},
					"project": {},
					"project_lite": {},
					"facility": {
						"id": "e1e9c52e-a0bc-4117-b996-0fc94843ea09",
						"name": "Parsippany, NJ",
						"code": "ewr1",
						"features": [
							"baremetal",
							"storage",
							"backend_transfer",
							"layer_2",
							"global_ipv4"
						],
						"address": null,
						"ip_ranges": [
							"2604:1380::/36",
							"147.75.192.0/21",
							"147.75.64.0/22",
							"147.75.72.0/22",
							"147.75.76.0/22",
							"147.75.104.0/22",
							"147.75.96.0/20"
						]
					},
					"assigned_to": {
						"href": "/devices/519e59c4-1cf7-41a3-98d3-85d48eb9bdcd"
					},
					"interface": {
						"href": "/ports/57c5d427-49c3-4497-a5fa-ed646a3e1cd2"
					},
					"network": "147.75.38.178",
					"address": "147.75.38.179",
					"gateway": "147.75.38.178",
					"href": "/ips/5d475f9c-bb56-455c-9347-5d7fcfc93e60"
				},
				{
					"id": "9d382090-59cd-4da9-8f8a-82b9bdca5c70",
					"address_family": 6,
					"netmask": "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe",
					"created_at": "2019-07-04T19:16:16Z",
					"details": null,
					"public": true,
					"cidr": 127,
					"management": true,
					"manageable": true,
					"enabled": true,
					"global_ip": null,
					"customdata": {},
					"project": {},
					"project_lite": {},
					"facility": {
						"id": "e1e9c52e-a0bc-4117-b996-0fc94843ea09",
						"name": "Parsippany, NJ",
						"code": "ewr1",
						"features": [
							"baremetal",
							"storage",
							"backend_transfer",
							"layer_2",
							"global_ipv4"
						],
						"address": null,
						"ip_ranges": [
							"2604:1380::/36",
							"147.75.192.0/21",
							"147.75.64.0/22",
							"147.75.72.0/22",
							"147.75.76.0/22",
							"147.75.104.0/22",
							"147.75.96.0/20"
						]
					},
					"assigned_to": {
						"href": "/devices/519e59c4-1cf7-41a3-98d3-85d48eb9bdcd"
					},
					"interface": {
						"href": "/ports/57c5d427-49c3-4497-a5fa-ed646a3e1cd2"
					},
					"network": "2604:1380:0:da00::2",
					"address": "2604:1380:0:da00::3",
					"gateway": "2604:1380:0:da00::2",
					"href": "/ips/9d382090-59cd-4da9-8f8a-82b9bdca5c70"
				},
				{
					"id": "5932ba9b-cc9a-4ca3-8dac-217eb970c982",
					"address_family": 4,
					"netmask": "255.255.255.254",
					"created_at": "2019-07-04T19:16:16Z",
					"details": null,
					"public": false,
					"cidr": 31,
					"management": true,
					"manageable": true,
					"enabled": true,
					"global_ip": null,
					"customdata": {},
					"project": {},
					"project_lite": {},
					"facility": {
						"id": "e1e9c52e-a0bc-4117-b996-0fc94843ea09",
						"name": "Parsippany, NJ",
						"code": "ewr1",
						"features": [
							"baremetal",
							"storage",
							"backend_transfer",
							"layer_2",
							"global_ipv4"
						],
						"address": null,
						"ip_ranges": [
							"2604:1380::/36",
							"147.75.192.0/21",
							"147.75.64.0/22",
							"147.75.72.0/22",
							"147.75.76.0/22",
							"147.75.104.0/22",
							"147.75.96.0/20"
						]
					},
					"assigned_to": {
						"href": "/devices/519e59c4-1cf7-41a3-98d3-85d48eb9bdcd"
					},
					"interface": {
						"href": "/ports/57c5d427-49c3-4497-a5fa-ed646a3e1cd2"
					},
					"network": "10.99.141.2",
					"address": "10.99.141.3",
					"gateway": "10.99.141.2",
					"href": "/ips/5932ba9b-cc9a-4ca3-8dac-217eb970c982"
				}
			],
			"plan": {
				"id": "e69c0169-4726-46ea-98f1-939c9e8a3607",
				"slug": "t1.small.x86",
				"name": "t1.small.x86",
				"description": "Our Type 0 configuration is a general use \"cloud killer\" server, with a Intel Atom 2.4Ghz processor and 8GB of RAM.",
				"line": "baremetal",
				"specs": {
					"cpus": [
						{
							"count": 1,
							"type": "Intel Atom C2550 @ 2.4Ghz"
						}
					],
					"memory": {
						"total": "8GB"
					},
					"drives": [
						{
							"count": 1,
							"size": "80GB",
							"type": "SSD"
						}
					],
					"nics": [
						{
							"count": 2,
							"type": "1Gbps"
						}
					],
					"features": {
						"raid": false,
						"txt": true
					}
				},
				"legacy": false,
				"deployment_types": [
					"on_demand",
					"spot_market"
				],
				"available_in": [
					{
						"href": "/facilities/8e6470b3-b75e-47d1-bb93-45b225750975"
					},
					{
						"href": "/facilities/2b70eb8f-fa18-47c0-aba7-222a842362fd"
					},
					{
						"href": "/facilities/8ea03255-89f9-4e62-9d3f-8817db82ceed"
					},
					{
						"href": "/facilities/e1e9c52e-a0bc-4117-b996-0fc94843ea09"
					}
				],
				"class": "t1.small.x86",
				"pricing": {
					"hour": 0.07
				}
			},
			"userdata": "",
			"switch_uuid": "2b57eb1a",
			"network_ports": [
				{
					"id": "57c5d427-49c3-4497-a5fa-ed646a3e1cd2",
					"type": "NetworkBondPort",
					"name": "bond0",
					"data": {
						"bonded": true
					},
					"network_type": "layer3",
					"native_virtual_network": null,
					"hardware": {
						"href": "/hardware/71140b4d-b2c8-4862-b616-e0a2dfb8f355"
					},
					"virtual_networks": [],
					"connected_port": null,
					"href": "/ports/57c5d427-49c3-4497-a5fa-ed646a3e1cd2"
				},
				{
					"id": "5656cebe-c04b-4139-96cd-e6402f42a7cf",
					"type": "NetworkPort",
					"name": "eth0",
					"data": {
						"mac": "0c:c4:7a:e5:44:7a",
						"bonded": true
					},
					"bond": {
						"id": "57c5d427-49c3-4497-a5fa-ed646a3e1cd2",
						"name": "bond0"
					},
					"native_virtual_network": null,
					"hardware": {
						"href": "/hardware/71140b4d-b2c8-4862-b616-e0a2dfb8f355"
					},
					"virtual_networks": [],
					"connected_port": {
						"href": "/ports/c6fa3cb8-61c6-405e-af02-45a8cbe2e0fd"
					},
					"href": "/ports/5656cebe-c04b-4139-96cd-e6402f42a7cf"
				},
				{
					"id": "2562fed5-0ba8-4c41-af8d-79380cecb72f",
					"type": "NetworkPort",
					"name": "eth1",
					"data": {
						"mac": "0c:c4:7a:e5:44:7b",
						"bonded": true
					},
					"bond": {
						"id": "57c5d427-49c3-4497-a5fa-ed646a3e1cd2",
						"name": "bond0"
					},
					"native_virtual_network": null,
					"hardware": {
						"href": "/hardware/71140b4d-b2c8-4862-b616-e0a2dfb8f355"
					},
					"virtual_networks": [],
					"connected_port": {
						"href": "/ports/d4047432-cfe5-4c61-8809-5a2e7d89ad1d"
					},
					"href": "/ports/2562fed5-0ba8-4c41-af8d-79380cecb72f"
				}
			],
			"href": "/devices/519e59c4-1cf7-41a3-98d3-85d48eb9bdcd"
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
devices,err := a.GetDevices()
if err != nil {
t.Error(err)
}
if len(devices) == 0 {
t.Error("Unable to load devices from example payload")
}
}