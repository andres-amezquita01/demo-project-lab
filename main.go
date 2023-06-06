package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkbootmux "github.com/rookie-ninja/rk-boot/mux"
	rkmuxinter "github.com/rookie-ninja/rk-mux/interceptor"
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

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests received",
		},
	)
	requestsFail = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_fails",
			Help: "Total number of requests fails",
		},
	)
	requestsAdd = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_add",
			Help: "Total number of requests to add",
		},
	)
	requestsBin = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_bin",
			Help: "Total number of requests to convert from decimal to binary",
		},
	)
	requestsSub = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_sub",
			Help: "Total number of requests to subtract",
		},
	)
	requestsMul = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_mul",
			Help: "Total number of requests to multiply",
		},
	)
	requestsDiv = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_div",
			Help: "Total number of requests to division",
		},
	)
	requestsHealth = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_health",
			Help: "Total number of requests to verify the health of the app",
		},
	)
	requestsMain = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_main",
			Help: "Total number of requests to the main page",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestsFail)
	prometheus.MustRegister(requestsAdd)
	prometheus.MustRegister(requestsBin)
	prometheus.MustRegister(requestsSub)
	prometheus.MustRegister(requestsMul)
	prometheus.MustRegister(requestsDiv)
	prometheus.MustRegister(requestsHealth)
	prometheus.MustRegister(requestsMain)
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
	requestsAdd.Inc()
	requestsTotal.Inc()
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
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	subtraction := num1 + num2
	requestsTotal.Inc()
	requestsSub.Inc()

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
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	result := DecimalToBinary(num1)

	requestsTotal.Inc()
	requestsBin.Inc()

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
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	subtraction := num1 * num2
	requestsTotal.Inc()
	requestsMul.Inc()
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
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(request.FormValue("num2"))
	if err != nil {
		requestsFail.Inc()
		http.Error(writer, "Invalid value for num2", http.StatusBadRequest)
		return
	}

	division := num1 / num2
	requestsTotal.Inc()
	requestsDiv.Inc()

	rkmuxinter.WriteJson(writer, http.StatusOK, &DivResponse{
		Result: division,
	})
}

// @Summary Health service
// @Id 7
// @version 1.0
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /health [get]
func Health(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Health ok!"),
	})
	requestsTotal.Inc()
	requestsHealth.Inc()
}

func Test(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("testing!"),
	})
	requestsTotal.Inc()
}

func MainPage(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello, please go to /docs to see the documentation!!!"),
	})
	requestsTotal.Inc()
	requestsMain.Inc()
}

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkbootmux.GetMuxEntry("greeter")
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/v1/greeter").HandlerFunc(Greeter)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/health").HandlerFunc(Health)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/test").HandlerFunc(Test)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/").HandlerFunc(MainPage)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/add").HandlerFunc(AddNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/bin").HandlerFunc(ConvertIntToBinary)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/sub").HandlerFunc(SubtractionNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/mul").HandlerFunc(MultiplyNumbers)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/div").HandlerFunc(DivisionNumbers)
	entry.Router.NewRoute().Path("/metrics").Handler(promhttp.Handler())

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}
