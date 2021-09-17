module github.ibm.com/jmuro/tronci

go 1.16

replace github.com/bndr/gojenkins v1.1.0 => github.com/muroj/gojenkins v1.1.1

#replace github.com/muroj/gojenkins v1.1.1 => ../gojenkins

require (
	github.com/antihax/optional v1.0.0
	github.com/muroj/gojenkins v1.1.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
)
