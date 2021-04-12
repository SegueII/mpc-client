package main

import (
	"log"
	"net/http"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/modules/service"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/types/store"
)

var (
	nodeURI  = "tcp://localhost:26657"
	grpcAddr = "localhost:9090"
	chainID  = "testing"
	client   sdk.IRISHUBClient
)

const (
	// addr     = "iaa1jw2cg8cuthlazycv3h44qrvkqccw2w36uup8tt"
	keyName  = "key"
	password = "12345678"
	mnemonic = "fiscal fix ecology become link exile web critic weird boring canvas engine stomach delay click song amount maze skill october web rude choose female"
)

func main() {
	client = initClient()
	setKey(client, keyName, password, mnemonic)

	startServer()
}

func initClient() sdk.IRISHUBClient {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	return sdk.NewIRISHUBClient(cfg)
}

func setKey(client sdk.IRISHUBClient, name string, password string, mnemonic string) {
	_, _ = client.Key.Recover(name, password, mnemonic)
}

func startServer() {
	http.HandleFunc("/", QueryRateHandlerFn)
	log.Println("Starting mock server ...")
	log.Fatal(http.ListenAndServe(":1210", nil))
}

func QueryRateHandlerFn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("content-type", "application/json")

	baseTx := types.BaseTx{
		From:     keyName,
		Gas:      200000,
		Memo:     "",
		Mode:     types.Sync,
		Password: password,
	}

	invokeServiceRequest := service.InvokeServiceRequest{
		ServiceName:       "test",
		Providers:         []string{"iaa179jc96sqxfyegr4vtwgu3gr32q42xn8mkefver"},
		Input:             `{"header":{},"body":{}}`,
		ServiceFeeCap:     types.NewDecCoins(types.NewDecCoin("iris", types.NewInt(500))),
		Timeout:           100,
		Repeated:          false,
		RepeatedFrequency: 0,
		RepeatedTotal:     0,
		Callback:          nil,
	}

	id, result, err := client.Service.InvokeService(invokeServiceRequest, baseTx)
	if err != nil {
		println(err.Error())
		_, _ = w.Write([]byte(`{"mssage":"ok"}`))
		return
	}
	println("===================")
	println(id)
	println(result.Events.String())

	_, _ = w.Write([]byte(`{"mssage":"ok"}`))
}
