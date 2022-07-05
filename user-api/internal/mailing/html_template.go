package mailing

type IHtmlTemplate interface {
	GetTrackingTemplate(number string) string
}
type HtmlTemplate struct{}

func NewHtmlTemplate() *HtmlTemplate {
	return &HtmlTemplate{}
}

func (h *HtmlTemplate) GetTrackingTemplate(number string) string {
	return ""
}
