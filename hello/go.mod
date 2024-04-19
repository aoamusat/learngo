module learngo.io/hello

go 1.21.1

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace learngo.io/greetings => ../greetings

require learngo.io/greetings v0.0.0-00010101000000-000000000000
