package postgres

import (
	"context"
	"database/sql"

	"github.com/golang-module/carbon/v2"
	"github.com/lvdigitalpro/back/src/data/contracts"
	"golang.org/x/crypto/bcrypt"
)

type ProjectsRepository struct {
	DB *sql.DB
	contracts.IProjectsRepository
}

func NewProjectsRepository(db *sql.DB) contracts.IProjectsRepository {
	return &ProjectsRepository{DB: db}
}

func (r *ProjectsRepository) NewProject(ctx context.Context,
	input contracts.InputNewProjectContract,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "insert into platform.projects (cod_project, type, status, name, description, created_at, user_id) values(nextval('cod_project'), $1, $2, $3, $4, $5, $6);")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(input.Type, contracts.QUEUE, input.Name, input.Description, carbon.Now().ToIso8601String(), input.UserID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	response := "Project created successfully"

	return &response, nil
}

func (r *ProjectsRepository) UpdateProject(ctx context.Context,
	input contracts.InputUpdateProjectContract,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "update platform.projects set type = $1, name = $2, description = $3, updated_at = $4 where cod_project = $5;")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(input.Type, input.Name, input.Description, carbon.Now().ToIso8601String(), input.CodProject)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	response := "Project updated successfully"

	return &response, nil
}

func (r *ProjectsRepository) DeleteProject(ctx context.Context,
	input contracts.InputDeleteProjectContract,
) (*string, error) {

	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	get := tx.QueryRowContext(ctx, "select cod_project,user_id from platform.projects where cod_project = $1;", input.CodProject)

	var projectGet contracts.ProjectContract

	err = get.Scan(&projectGet.CodProject, &projectGet.UserID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var userGet contracts.UserContract

	err = tx.QueryRowContext(ctx, "select u.user_id, u.password from platform.users as u where u.user_id = $1;", projectGet.UserID).Scan(&userGet.UserID, &userGet.Password)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userGet.Password), []byte(input.Password))

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "delete from platform.projects where cod_project = $1;")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(input.CodProject)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	response := "Project deleted successfully"

	return &response, nil
}

func (r *ProjectsRepository) GetProjects(ctx context.Context) ([]*contracts.ProjectContract, error) {

	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, "select * from platform.projects;")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var projects []*contracts.ProjectContract

	for rows.Next() {
		var project contracts.ProjectContract

		err = rows.Scan(&project.CodProject, &project.Type, &project.Status, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt, &project.StartDate, &project.EndDate, &project.UserID)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		projects = append(projects, &project)

	}

	tx.Commit()

	return projects, nil
}

func (r *ProjectsRepository) GetProject(ctx context.Context, cod_project int) (*contracts.ProjectContract, error) {

	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from platform.projects where cod_project = $1;")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var project contracts.ProjectContract

	err = tx.StmtContext(ctx, prepare).QueryRowContext(ctx, cod_project).Scan(&project.CodProject, &project.Type, &project.Status, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt, &project.StartDate, &project.EndDate, &project.UserID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &project, nil
}
