package main

import (
	"context"
	"fmt"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-boot/mux"
	"github.com/rookie-ninja/rk-mux/interceptor"
	"net/http"
	"strconv"
)

type ResultResponse struct {
	Message string
}
type AddRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

// Response.
type AddResponse struct {
	Result int `json:"result"`
}

// Response.
type SubResponse struct {
	Result int `json:"result"`
}

// Response.
type MulResponse struct {
	Result int `json:"result"`
}

// Response.
type DivResponse struct {
	Result int `json:"result"`
}

// Response.
type GreeterResponse struct {
	Message string
}

// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Param name query string true "Input name"
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", request.URL.Query().Get("name")),
	})
}

// @Summary Add numbers service
// @Id 2
// @version 1.0
// @produce application/json
// @Param num1 formData int true "first number"
// @Param num2 formData int true "second number"
// @Success 200 {object} AddResponse
// @Router /add [post]
func AddNumbers(writer http.ResponseWriter, request *http.Request) {
	num1, err := strconv.Atoi(request.FormValue("num1"))
	if err != nil {
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	sum := num1 + num2

	rkmuxinter.WriteJson(writer, http.StatusOK, &AddResponse{
		Result: sum,
	})
}

// @Summary Subtract service
// @Id 3
// @version 1.0
// @produce application/json
// @Param num1 formData int true "first number"
// @Param num2 formData int true "second number"
// @Success 200 {object} SubResponse
// @Router /sub [post]
func SubtractionNumbers(writer http.ResponseWriter, request *http.Request) {
	num1, err := strconv.Atoi(request.FormValue("num1"))
	if err != nil {
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	subtraction := num1 + num2

	rkmuxinter.WriteJson(writer, http.StatusOK, &SubResponse{
		Result: subtraction,
	})
}

// @Summary Decimal to binary service
// @Id 4
// @version 1.0
// @produce application/json
// @Param num1 formData int true "number"
// @Success 200 {object} ResultResponse
// @Router /bin [post]
func ConvertIntToBinary(writer http.ResponseWriter, request *http.Request) {
	num1, err := strconv.Atoi(request.FormValue("num1"))
	if err != nil {
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	result := DecimalToBinary(num1)
	rkmuxinter.WriteJson(writer, http.StatusOK, &ResultResponse{
		Message: result,
	})

}

// @Summary Multiply service
// @Id 5
// @version 1.0
// @produce application/json
// @Param num1 formData int true "first number"
// @Param num2 formData int true "second number"
// @Success 200 {object} MulResponse
// @Router /mul [post]
func MultiplyNumbers(writer http.ResponseWriter, request *http.Request) {
	num1, err := strconv.Atoi(request.FormValue("num1"))
	if err != nil {
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	subtraction := num1 * num2

	rkmuxinter.WriteJson(writer, http.StatusOK, &MulResponse{
		Result: subtraction,
	})
}

// @Summary Division service
// @Id 6
// @version 1.0
// @produce application/json
// @Param num1 formData int true "first number"
// @Param num2 formData int true "second number"
// @Success 200 {object} DivResponse
// @Router /div [post]
func DivisionNumbers(writer http.ResponseWriter, request *http.Request) {
	num1, err := strconv.Atoi(request.FormValue("num1"))
	if err != nil {
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	division := num1 / num2

	rkmuxinter.WriteJson(writer, http.StatusOK, &DivResponse{
		Result: division,
	})
}

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkbootmux.GetMuxEntry("greeter")
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/v1/greeter").HandlerFunc(Greeter)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/add").HandlerFunc(AddNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/bin").HandlerFunc(ConvertIntToBinary)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/sub").HandlerFunc(SubtractionNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/mul").HandlerFunc(MultiplyNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/div").HandlerFunc(DivisionNumbers)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}
