package captcha_resolver

type RandomString string
type SolvedCaptcha string

type CaptchaSolverGetIdResponse struct {
	Status  string
	Request string
}

type CaptchaSolverResponse struct {
	Status  string
	Text    string
	Request string
}
