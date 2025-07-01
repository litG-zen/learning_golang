package main

const (
	// PORT                = rand.IntN(65535) `This will not work as Golang returns a value error namely
	//                                         [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.'
	PORT                = 8000 // Adding server spinup on random for erytime, avoid attacks
	INVALID_API_MSG     = "Invalid API KEY"
	LOG_EXP_DELTA       = 86400 //24*60*60
	JWT_TOKEN_EXP_DELTA = 86400 //24*60*60
)
