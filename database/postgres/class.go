package postgresdb

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func (repo *PostgresRepository) GetClasses(ctx context.Context, filters *models.ClassFilter) ([]*models.Class, error) {
	// Using strings.Builder for best performance of string concatenations. For performance and query injection preventing,
	// this way is more complex to code, but it's more efficient and secure.
	var filterParams []any
	var queryString strings.Builder
	var totalFilters = 0

	queryString.WriteString("SELECT id, name, start_date, end_date, capacity FROM classes WHERE 1=1")

	if !filters.Date.IsZero() {
		/*
			AND f.date BETWEEN start_date AND end_date
		*/
		queryString.WriteString(" AND $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.Date)
		queryString.WriteString(" BETWEEN start_date AND end_date")
	}

	// Get all the classes found in specific date range. At least a single day of a class must match to it be returned
	if !filters.StartDate.IsZero() && !filters.EndDate.IsZero() {
		/*
			AND (
				start_date BETWEEN f.startDate AND f.endDate
				OR end_date BETWEEN f.startDate AND fhttp.StatusInternalServerError.endDate
				OR (
					start_date <= f.startDate
					AND f.endDate <= end_date
				)
			)
		*/
		queryString.WriteString(" AND (start_date BETWEEN $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.StartDate)
		queryString.WriteString(" AND $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.EndDate)
		queryString.WriteString(" OR end_date BETWEEN $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.StartDate)
		queryString.WriteString(" AND $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.EndDate)
		queryString.WriteString(" OR (start_date <= $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.StartDate)
		queryString.WriteString(" AND $")
		totalFilters++
		queryString.WriteString(strconv.Itoa(totalFilters))
		filterParams = append(filterParams, filters.EndDate)
		queryString.WriteString(" <= end_date))")
	}

	queryString.WriteString(" ORDER BY name ASC")

	rows, err := repo.db.QueryContext(ctx, queryString.String(), filterParams...)

	if err != nil {
		return nil, err
	}

	// Closing returned dataset
	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var classes []*models.Class

	for rows.Next() {
		var c = models.Class{}

		if err := rows.Scan(&c.Id, &c.Name, &c.StartDate, &c.EndDate, &c.Capacity); err != nil {
			return nil, err
		}

		classes = append(classes, &c)
	}

	return classes, nil
}

func (repo *PostgresRepository) InsertClass(ctx context.Context, c *models.Class) error {
	var filters = models.ClassFilter{
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
	}

	classes, err := repo.GetClasses(ctx, &filters)

	if err != nil {
		return err
	}

	if len(classes) > 0 {
		return errors.New("There can only be one class a day")
	}

	_, err = repo.db.ExecContext(ctx, "INSERT INTO classes (id, name, start_date, end_date, capacity) VALUES ($1, $2, $3, $4, $5)", c.Id, c.Name, c.StartDate, c.EndDate, c.Capacity)
	return err
}

func (repo *PostgresRepository) TruncateClass(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "TRUNCATE TABLE classes")
	return err
}
