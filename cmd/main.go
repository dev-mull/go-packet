package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	packet "packet-cli"
)

var (
	cfgFile                                                            string
	listOS, listFac, listPlan, listDevices, create, delete, typeDevice,typeProject,debug bool
	facility, plan, osName, hostname, id,action,eventType                              string
)

func init() {
	user, _ := user.Current()
	flag.StringVar(&cfgFile, "config", user.HomeDir+"/.packet_rc", "config yaml file")
	flag.BoolVar(&typeDevice, "device", false, "Flag to indicate actions are intended for a device")
	flag.BoolVar(&create, "create", false, "Create a new element")
	flag.BoolVar(&delete, "delete", false, "Delete existing element by id --id required")
	flag.BoolVar(&listOS, "oses", false, "List OSes")
	flag.BoolVar(&debug, "debug", false, "Debug requests to the api")
	flag.BoolVar(&listFac, "facilities", false, "List Facilities available to the project")
	flag.BoolVar(&listPlan, "plans", false, "List Plans available to the project")
	flag.BoolVar(&listDevices, "devices", false, "List Devices available to the project")
	flag.StringVar(&eventType, "events", "", "list events by type")
	flag.StringVar(&action, "action", "", "Action (power_on,power_off,etc)")
	flag.StringVar(&hostname, "hostname", "", "Hostname")
	flag.StringVar(&osName, "os", "", "Operating System Name")
	flag.StringVar(&id, "id", "", "Id of object to manipulate")
	flag.StringVar(&facility, "facility", "", "Facility Name")
	flag.StringVar(&plan, "plan", "", "Plan Name")
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		fmt.Printf("Failed to open the config file %v %v",cfgFile,err)
		os.Exit(1)
	}
	cfg := &packet.Config{}
	err = yaml.Unmarshal(b,cfg )
	if err != nil {
	fmt.Printf("Failed to decode config file %v %v",cfgFile,err)
	os.Exit(1)
	}
	a, err := packet.NewAPI(cfg)
	a.Debug = debug
	if err != nil {
		log.Fatal(err)
	}

	if listOS {
		oses, err := a.GetOSes()
		if err != nil {
			log.Fatal(err)
		}
		for _, os := range oses {
			fmt.Printf("%s\n", os.Name)
			fmt.Printf("\tID: %s\n", os.Id)
			fmt.Printf("\tVersion: %s\n", os.Version)
			fmt.Printf("\tDistro: %s\n", os.Distro)
			fmt.Printf("\tSlug: %s\n", os.Slug)
			fmt.Printf("\tProvisionable On:\n")
			if len(os.ProvisionableOn) > 0 {
				for _, p := range os.ProvisionableOn {
					fmt.Printf("\t\t%s\n", p)
				}
			}
			if os.Licensed {
				fmt.Printf("\tLicensed\n")
			} else {
			fmt.Printf("\tUnlicensed\n")
			}
			//, Version: %s, Distro: %s\n",os.Name,os.Id,os.Version,os.Distro)
		}
	}
	if eventType != "" {
		events,err := a.GetEvents(eventType,id)
		if err != nil {
			fmt.Printf("Failed: %v\n",err)
			return
		}
		for _,e := range events {
			fmt.Printf("At: %v Type: %s %v\n",e.CreatedAt,e.Type,e.Interpolated)
		}


	}
	if listFac {
		facilities, err := a.GetFacilites()
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range facilities {
			fmt.Printf("%s\n", f.Name)
			fmt.Printf("\tID:%s\n", f.Id)
			fmt.Printf("\tCode:%s\n", f.Code)
		}
	}

	if listPlan {
		plans, err := a.GetPlans()
		if err != nil {
			log.Fatal(err)
		}
		for _, plan := range plans {
			fmt.Printf("%s: %s\n", plan.Name, plan.Description)
			fmt.Printf("\tLine: %s\n", plan.Line)
			fmt.Printf("\tID: %s\n", plan.Id)
		}
	}

	if listDevices {
		devices, err := a.GetDevices()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(
			"%-40s%-20s%-15s%-15s%15s/%s%15s%15s\n",
			"ID",
			"Name",
			"State",
			"Facility",
			"Distro",
			"Version",
			"Plan",
			"Locked",
		)
		for _, device := range devices {
			fmt.Printf(
				"%-40s%-20s%-15s%-15s%15s/%s%15s%15v\n",
				device.Id,
				device.Hostname,
				device.State,
				device.Facility.Name,
				device.OS.Distro,
				device.OS.Version,
				device.Plan.Name,
				device.Locked,
			)
		}
	}
	if action != "" && typeDevice {
		if id == "" {
			fmt.Println("--id missing from arguments")
			return
		}
		err := a.DeviceAction(id,action)
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
		} else {
			fmt.Printf("Success: %s %s\n", id,action)
		}
		return

	}
	if typeDevice && create {
		device := &packet.NewDevice{
			Facility: facility,
			Plan:     plan,
			OS:       osName,
		}
		if hostname != "" {
			device.Hostname = &hostname
		}
		d, err := a.CreateDevice(device)
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
		}
		fmt.Printf("Created: %s %s %s\n", d.Id, d.Hostname, d.State)
		return
	}

	if typeDevice && delete {
		if id == "" {
			fmt.Println("--id missing from arguments")
			return
		}
		err := a.DeleteDevice(id)
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
		} else {
			fmt.Printf("Deleted: %s\n", id)
		}

		return
	}

}
