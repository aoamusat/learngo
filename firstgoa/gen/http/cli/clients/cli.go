// Code generated by goa v3.16.1, DO NOT EDIT.
//
// clients HTTP client CLI support package
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	clientc "learngo.io/firstgoa/gen/http/client/client"
	signinc "learngo.io/firstgoa/gen/http/signin/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `client (add|get|show)
signin authenticate
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` client add --body '{
      "ClientName": "Ducimus sapiente suscipit consequatur minima animi.",
      "ContactEmail": "Id enim deleniti qui est ab.",
      "ContactMobile": 3474453470839196953,
      "ContactName": "Magnam quod ratione provident consequatur ea esse."
   }' --client-id "Veniam dolor." --token "Repellat commodi omnis suscipit eaque."` + "\n" +
		os.Args[0] + ` signin authenticate --username "user" --password "password"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		clientFlags = flag.NewFlagSet("client", flag.ContinueOnError)

		clientAddFlags        = flag.NewFlagSet("add", flag.ExitOnError)
		clientAddBodyFlag     = clientAddFlags.String("body", "REQUIRED", "")
		clientAddClientIDFlag = clientAddFlags.String("client-id", "REQUIRED", "Client ID")
		clientAddTokenFlag    = clientAddFlags.String("token", "REQUIRED", "")

		clientGetFlags        = flag.NewFlagSet("get", flag.ExitOnError)
		clientGetClientIDFlag = clientGetFlags.String("client-id", "REQUIRED", "Client ID")
		clientGetTokenFlag    = clientGetFlags.String("token", "REQUIRED", "")

		clientShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		clientShowTokenFlag = clientShowFlags.String("token", "REQUIRED", "")

		signinFlags = flag.NewFlagSet("signin", flag.ContinueOnError)

		signinAuthenticateFlags        = flag.NewFlagSet("authenticate", flag.ExitOnError)
		signinAuthenticateUsernameFlag = signinAuthenticateFlags.String("username", "REQUIRED", "Username used to perform signin")
		signinAuthenticatePasswordFlag = signinAuthenticateFlags.String("password", "REQUIRED", "Password used to perform signin")
	)
	clientFlags.Usage = clientUsage
	clientAddFlags.Usage = clientAddUsage
	clientGetFlags.Usage = clientGetUsage
	clientShowFlags.Usage = clientShowUsage

	signinFlags.Usage = signinUsage
	signinAuthenticateFlags.Usage = signinAuthenticateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "client":
			svcf = clientFlags
		case "signin":
			svcf = signinFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "client":
			switch epn {
			case "add":
				epf = clientAddFlags

			case "get":
				epf = clientGetFlags

			case "show":
				epf = clientShowFlags

			}

		case "signin":
			switch epn {
			case "authenticate":
				epf = signinAuthenticateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "client":
			c := clientc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = clientc.BuildAddPayload(*clientAddBodyFlag, *clientAddClientIDFlag, *clientAddTokenFlag)
			case "get":
				endpoint = c.Get()
				data, err = clientc.BuildGetPayload(*clientGetClientIDFlag, *clientGetTokenFlag)
			case "show":
				endpoint = c.Show()
				data, err = clientc.BuildShowPayload(*clientShowTokenFlag)
			}
		case "signin":
			c := signinc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "authenticate":
				endpoint = c.Authenticate()
				data, err = signinc.BuildAuthenticatePayload(*signinAuthenticateUsernameFlag, *signinAuthenticatePasswordFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// clientUsage displays the usage of the client command and its subcommands.
func clientUsage() {
	fmt.Fprintf(os.Stderr, `The Client service allows access to client members
Usage:
    %[1]s [globalflags] client COMMAND [flags]

COMMAND:
    add: Add implements add.
    get: Get implements get.
    show: Show implements show.

Additional help:
    %[1]s client COMMAND --help
`, os.Args[0])
}
func clientAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] client add -body JSON -client-id STRING -token STRING

Add implements add.
    -body JSON: 
    -client-id STRING: Client ID
    -token STRING: 

Example:
    %[1]s client add --body '{
      "ClientName": "Ducimus sapiente suscipit consequatur minima animi.",
      "ContactEmail": "Id enim deleniti qui est ab.",
      "ContactMobile": 3474453470839196953,
      "ContactName": "Magnam quod ratione provident consequatur ea esse."
   }' --client-id "Veniam dolor." --token "Repellat commodi omnis suscipit eaque."
`, os.Args[0])
}

func clientGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] client get -client-id STRING -token STRING

Get implements get.
    -client-id STRING: Client ID
    -token STRING: 

Example:
    %[1]s client get --client-id "Accusantium eos commodi natus eaque a." --token "Reiciendis natus autem sapiente."
`, os.Args[0])
}

func clientShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] client show -token STRING

Show implements show.
    -token STRING: 

Example:
    %[1]s client show --token "Commodi qui qui."
`, os.Args[0])
}

// signinUsage displays the usage of the signin command and its subcommands.
func signinUsage() {
	fmt.Fprintf(os.Stderr, `The Signin service authenticates users and validate tokens
Usage:
    %[1]s [globalflags] signin COMMAND [flags]

COMMAND:
    authenticate: Creates a valid JWT

Additional help:
    %[1]s signin COMMAND --help
`, os.Args[0])
}
func signinAuthenticateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] signin authenticate -username STRING -password STRING

Creates a valid JWT
    -username STRING: Username used to perform signin
    -password STRING: Password used to perform signin

Example:
    %[1]s signin authenticate --username "user" --password "password"
`, os.Args[0])
}
