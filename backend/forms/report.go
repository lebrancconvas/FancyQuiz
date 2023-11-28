package forms

type Report struct {
	ReportID uint64 `json:"report_id"`
	UserID uint64 `json:"user_id"`
	StatusID uint64 `json:"report_status_id"`
	Content string `json:"report_content"`
}
