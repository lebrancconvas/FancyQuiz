package models

import (
	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/forms"
)

type Report struct{}

func (r Report) GetAllReport() ([]forms.Report, error) {
	db := db.GetDB()

	var reports []forms.Report

	stmt, err := db.Prepare(`
		SELECT report_id, fk_user_id, report_content
		FROM reports
		WHERE used_flg = true
		ORDER BY created_at DESC
	`)
	if err != nil {
		return reports, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report forms.Report

		err := rows.Scan(&report.ReportID, &report.UserID, &report.Content)
		if err != nil {
			return reports, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}
