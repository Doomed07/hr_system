package postgres

import (
	"HR_system/http_server"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepo struct {
	Pool *pgxpool.Pool
}

func NewEmployeeRepo(pool *pgxpool.Pool) *EmployeeRepo {
	return &EmployeeRepo{
		Pool: pool,
	}
}

func (r *EmployeeRepo) CreateTableEmployees(ctx context.Context) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS employees (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR(200) NOT NULL,
	position VARCHAR(200) NOT NULL
	)`

	_, err := r.Pool.Exec(ctx, sqlQuery)

	return err
}

func (r *EmployeeRepo) AddEmp(ctx context.Context, emp http_server.Employee) (int, error) {
	sqlQuery := `
	INSERT INTO employees(full_name, position)
	VALUES ($1, $2)
	RETURNING id;
	`
	var id int
	err := r.Pool.QueryRow(
		ctx,
		sqlQuery,
		emp.FullName,
		emp.Position).
		Scan(&id)

	return id, err
}

func (r *EmployeeRepo) AllEmp(ctx context.Context) ([]http_server.Employee, error) {
	sqlQuery := `
	SELECT id, full_name, position
	FROM employees
	ORDER by id ASC
	`

	rows, err := r.Pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := make([]http_server.Employee, 0)

	for rows.Next() {
		var employee http_server.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Position); err != nil {
			return nil, err
		}

		employees = append(employees, employee)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepo) DelEmp(ctx context.Context, id int) error {
	sqlQuery := `
	DELETE FROM employees
	WHERE id = $1;
	`

	tag, err := r.Pool.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return http_server.ErrNotFoundEmployee
	}

	return nil
}
