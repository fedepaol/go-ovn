package main

import (
	"fmt"
	"time"

	apis "github.com/ebayinc/libovndb"
)

var ovndbapi apis.OVNDBApi

func init() {
	ovndbapi = apis.GetInstance("/var/run/openvswitch/ovnnb_db.sock", apis.UNIX, "", 0)
}

func LSWAdd() {
	ocmd := ovndbapi.LSWAdd("sentineltest")
	ovndbapi.Execute(ocmd)
}

func LSPAdd() {
	ocmd := ovndbapi.LSPAdd("sentineltest", "test")
	ovndbapi.Execute(ocmd)
}

func LSPSetAddress() {
	ocmd := ovndbapi.LSPSetAddress("test", "mac ip")
	ovndbapi.Execute(ocmd)
}

func LSPDel() {
	ocmd := ovndbapi.LSPDel("test")
	ovndbapi.Execute(ocmd)
}

func LSWDel() {
	ocmd := ovndbapi.LSWDel("sentineltest")
	ovndbapi.Execute(ocmd)
}

func ACLAdd() {
	ocmd := ovndbapi.ACLAdd("sentineltest", "to-lport", "outport == \"96d44061-1823-428b-a7ce-f473d10eb3d0\" && ip && ip.dst == 10.97.183.61", "drop", 1001, nil, false)
	ovndbapi.Execute(ocmd)
}

func ACLDel() {
	ocmd := ovndbapi.ACLDel("sentineltest", "to-lport", "outport == \"96d44061-1823-428b-a7ce-f473d10eb3d0\" && ip && ip.dst == 10.97.183.61", 1001)
	ovndbapi.Execute(ocmd)
}


func LISTLS() {
	ocmd := ovndbapi.LSWList()
	ovndbapi.Execute(ocmd)
	fmt.Printf("return: %v", ocmd.Results)
}

func ADAS() {
	ocmd := ovndbapi.ASAdd("test", []string{" "})
	ovndbapi.Execute(ocmd)
	fmt.Printf("return: %v", ocmd.Results)
}

func main() {

	LSWAdd()
	LSPAdd()
	LSPSetAddress()
	ACLAdd()
	LISTLS()
	ADAS()
	time.Sleep(10 * time.Second)
	ACLDel()
	time.Sleep(10 * time.Second)
	LSPDel()
	LSWDel()

}
