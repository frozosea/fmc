package captcha_resolver

type RandomString string
type SolvedCaptcha string

type CaptchaSolverGetIdResponse struct {
	Status  int
	Request string
}

type CaptchaSolverResponse struct {
	Status  int
	Text    string
	Request string
}
