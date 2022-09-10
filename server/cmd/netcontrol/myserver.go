package netcontrol

import (
	"fmt"
	//"math"

	"net/http"
	"strconv"

	"github.com/698un/owmcalcgo/cnfg"
	"github.com/698un/owmcalcgo/entity"
	//"github.com/698un/owmcalcgo/mymath"
)

func StartServer() {

	fmt.Println("Server Init...")

	//ONE_POINT
	http.HandleFunc("/physicsval", func(
		w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, reqJsonInfo(r))
	}) //HandleFunc

	//FUll_Parallel
	http.HandleFunc("/fullparallel", func(
		w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, reqJsonParallel(r))
	}) //HandleFunc

	//PART_Parallel
	http.HandleFunc("/partparallel", func(
		w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, reqJsonPartParallel(r))
	}) //HandleFunc

	//STAR_SEVER
	err := http.ListenAndServe(cnfg.PortNumber, nil)

	if err != nil {
		fmt.Println("Error_ServerNotStart")
	}

	//	fmt.Println("Server Init complette")

}

func reqJsonInfo(r *http.Request) string {
	var angX float64
	var angY float64
	var actualTime int64

	var str1 string

	//Default error
	var res string = "{\"Error\":true}"

	//Parse ANGX
	str1 = r.URL.Query().Get("angx")
	if str1 == "" {
		return res
	}
	angX, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		return res
	}

	//Parse ANGY
	str1 = r.URL.Query().Get("angy")
	if str1 == "" {
		return res
	}
	angY, err1 := strconv.ParseFloat(str1, 64)
	if err1 != nil {
		return res
	}

	//Parse ACTUAL
	str1 = r.URL.Query().Get("actual")
	if str1 == "" {
		actualTime = 10000 //very long time if error
	}
	actualTime, err2 := strconv.ParseInt(str1, 0, 64)
	if err2 != nil {
		actualTime = 10000 //very long time if error
	}

	//print parameters
	fmt.Println("angXY/actual=", angX, ",", angY, "actual=", actualTime)

	//get jsonStr from DAO
	res = entity.GetJsonStr(angX, angY, actualTime)
	//fmt.Println(res)

	if res == "X" {
		res = "{\"Error\":true}"

	}

	return res
}
