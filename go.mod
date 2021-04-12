module github.com/segueII/mpc-client

go 1.16

require github.com/irisnet/irishub-sdk-go v0.0.0-20210412080344-9444eb9fe310

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
