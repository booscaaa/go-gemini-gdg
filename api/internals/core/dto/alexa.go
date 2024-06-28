package dto

type AlexaResponse struct {
	Version  string               `json:"version"`
	Response AlexaResponseContent `json:"response"`
}

type AlexaResponseContent struct {
	OutputSpeech     AlexaOutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool              `json:"shouldEndSession"`
}

type AlexaOutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
