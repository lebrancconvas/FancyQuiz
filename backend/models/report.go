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
		SELECT report_id, fk_user_id, fk_report_status_id, report_content
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

		err := rows.Scan(&report.ReportID, &report.UserID, &report.StatusID, &report.Content)
		if err != nil {
			return reports, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func (r Report) CreateReport(userID uint64, report string) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		INSERT INTO reports (fk_user_id, fk_report_status_id, report_content)
		VALUES ($1, 1, $2)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, report)
	if err != nil {
		return err
	}

	return nil
}

func (r Report) DeleteReport(reportID uint64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE reports
		SET used_flg = false
		WHERE report_id = $1 
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(reportID)
	if err != nil {
		return err
	}

	return nil
}

func (r Report) UpdateReportToBeAccepted(reportID uint64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE reports
		SET fk_report_status_id = 2
		WHERE report_id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(reportID)
	if err != nil {
		return err
	}

	return nil
}

func (r Report) UpdateReportToBeCompleted(reportID uint64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		UPDATE reports
		SET fk_report_status_id = 3
		WHERE report_id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(reportID)
	if err != nil {
		return err
	}

	return nil
}
