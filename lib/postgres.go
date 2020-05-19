package lib

import (
	"CodingExercise/shared/log"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	Pool *pgxpool.Pool
)

type student string

//ConnectDB ...
func ConnectDB(ctx context.Context) {
	DSN := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		ServiceConfig.Postgres.User, ServiceConfig.Postgres.Password, ServiceConfig.Postgres.Host, ServiceConfig.Postgres.Port, ServiceConfig.Postgres.DBName, ServiceConfig.Postgres.SSLMode, ServiceConfig.Postgres.PoolMaxConns)

	var err error
	Pool, err = pgxpool.Connect(ctx, DSN)
	if err != nil {
		fmt.Println(DSN, "Unable to Connect to Postgres: ", err)
		log.Error(ctx, DSN, "Unable to Connect to Postgres: ", err)
		return
	}

	if _, err := Pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS students(student_id SERIAL PRIMARY KEY, student_name character varying NOT NULL,enrollment_at timestamp(6) without time zone NOT NULL);"); err != nil {
		log.Error(ctx, DSN, "Unable to CREATE student table: ", err)
	}

	if _, err = Pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS assignments(assignment_id SERIAL PRIMARY KEY, assignment_name character varying NOT NULL,maximum_score int,recorded_at timestamp(6) without time zone NOT NULL,updated_at timestamp(6) without time zone NOT NULL);"); err != nil {
		log.Error(ctx, DSN, "Unable to CREATE assignments table: ", err)
	}

	if _, err = Pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS student_assignments_score(id SERIAL PRIMARY KEY, student_id int,assignment_id int,assignment_name character varying NOT NULL,score int,maximum_score int,recorded_at timestamp(6) without time zone NOT NULL,updated_at timestamp(6) without time zone NOT NULL, UNIQUE(student_id,assignment_id));"); err != nil {
		log.Error(ctx, DSN, "Unable to CREATE assignments table: ", err)
	}
}

//Insert ...
func (s *Students) Insert(ctx context.Context) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, s.Name, "Insert unable to Acquire pg conn: ", err)
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf("INSERT INTO %s.%s (student_name,enrollment_at) values ($1,$2) ON CONFLICT DO NOTHING;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentsTable)
	if _, err := conn.Exec(ctx, query, s.Name, time.Now()); err != nil {
		fmt.Println("Insert Error DB Insert : ", err)
		log.Error(ctx, "Insert Error DB Insert :", err)
		return err
	}
	return nil
}

func (str student) deleteStudent(ctx context.Context) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, str, "deleteStudent unable to Acquire pg conn: ", err)
		return err
	}
	defer conn.Release()
	query := fmt.Sprintf("DELETE  from %s.%s where student_id=$1;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentsTable)
	if _, err := conn.Exec(ctx, query, str); err != nil {
		fmt.Println(query, str, "deleteStudent from StudentsTable unable to exec query : ", err)
		log.Error(ctx, query, str, "deleteStudent from StudentsTable unable to exec query : ", err)
		return err
	}
	query = fmt.Sprintf("DELETE  from %s.%s where student_id=$1;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentAssignmentsScoreTable)
	if _, err := conn.Exec(ctx, query, str); err != nil {
		fmt.Println(query, str, "deleteStudent from StudentAssignmentsScoreTable unable to exec query : ", err)
		log.Error(ctx, query, str, "deleteStudent from StudentAssignmentsScoreTable unable to exec query : ", err)
		return err
	}

	return nil
}

func (str student) getStudent(ctx context.Context) (*Students, error) {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, str, "getStudent unable to Acquire pg conn: ", err)
		return nil, err
	}
	defer conn.Release()

	var std Students

	query := fmt.Sprintf("select * from %s.%s where student_id=$1;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentsTable)
	if err := conn.QueryRow(ctx, query, str).Scan(&std.ID, &std.Name, &std.EnrollmentAT); err != nil {
		fmt.Println(query, str, "getStudent unable to query : ", err)
		log.Error(ctx, query, str, "getStudent unable to query : ", err)
		return nil, err
	}

	return &std, nil
}

func fetchAllStudent(ctx context.Context) ([]StudentAssignmentsScore, error) {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, "fetchAllStudent unable to Acquire pg conn: ", err)
		return nil, err
	}
	defer conn.Release()

	var assings []StudentAssignmentsScore
	var assing StudentAssignmentsScore
	query := fmt.Sprintf("select id,student_id,assignment_id,assignment_name,score,maximum_score,recorded_at,updated_at from %s.%s ;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentAssignmentsScoreTable)
	fmt.Println("fetchAllStudent query : ", query)

	rows, err := conn.Query(ctx, query)
	fmt.Println("fetchAllStudent conn.Query : ", query)
	if err != nil {
		return assings, err
	}

	for rows.Next() {
		if err := rows.Scan(&assing.ID, &assing.StudentID, &assing.AssignmentID, &assing.AssignmentName, &assing.Score, &assing.MaximumScore, &assing.RecordedAT, &assing.UpdatedAT); err != nil {
			fmt.Println("error in row Scan: ", err)
			continue
		}
		assings = append(assings, assing)
	}

	return assings, nil
}

//Insert ...
func (assing Assignments) Insert(ctx context.Context) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, assing.AssignmentName, "Assignments Insert unable to Acquire pg conn: ", err)
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf("INSERT INTO %s.%s (assignment_name,maximum_score,recorded_at,updated_at) values ($1,$2,$3,$4) ON CONFLICT DO NOTHING;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.AssignmentsTable)
	if _, err := conn.Exec(ctx, query, assing.AssignmentName, assing.MaximumScore, time.Now(), time.Now()); err != nil {
		fmt.Println("Insert Error DB Insert : ", err)
		log.Error(ctx, "Insert Error DB Insert :", err)
		return err
	}
	return nil
}

//Insert ...
func (score StudentAssignmentsScore) Insert(ctx context.Context) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, score.AssignmentName, "StudentAssignmentsScore Insert unable to Acquire pg conn: ", err)
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf("select assignment_name,maximum_score from %s.%s where assignment_id=$1;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.AssignmentsTable)
	if err := conn.QueryRow(ctx, query, score.AssignmentID).Scan(&score.AssignmentName, &score.MaximumScore); err != nil {
		log.Error(ctx, "StudentAssignmentsScore Insert, QueryRow Error :", err)
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s.%s (student_id,assignment_id,assignment_name,score,maximum_score,recorded_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT DO NOTHING;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentAssignmentsScoreTable)
	if _, err := conn.Exec(ctx, query, score.StudentID, score.AssignmentID, score.AssignmentName, score.Score, score.MaximumScore, time.Now(), time.Now()); err != nil {
		fmt.Println("StudentAssignmentsScore Insert, Error DB Insert : ", err)
		log.Error(ctx, "StudentAssignmentsScore Insert, Error DB Insert :", err)
		return err
	}
	return nil
}

func fetchAssignmentsList(ctx context.Context) ([]Assignments, error) {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, "fetchAssignmentsList unable to Acquire pg conn: ", err)
		return nil, err
	}
	defer conn.Release()

	var assings []Assignments
	var assing Assignments
	query := fmt.Sprintf("select assignment_id,assignment_name,maximum_score,recorded_at,updated_at from %s.%s ;", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.AssignmentsTable)
	fmt.Println("fetchAssignmentsList query : ", query)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return assings, err
	}

	for rows.Next() {
		if err := rows.Scan(&assing.AssignmentsID, &assing.AssignmentName, &assing.MaximumScore, &assing.RecordedAT, &assing.UpdatedAT); err != nil {
			fmt.Println("fetchAssignmentsList error in row Scan: ", err)
			continue
		}
		assings = append(assings, assing)
	}

	return assings, nil
}

// Update ...
func (assignScore StudentAssignmentsScore) Update(ctx context.Context) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		log.Error(ctx, "assignScore StudentAssignmentsScore Update unable to Acquire pg conn: ", err)
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf("update %s.%s set updated_at=now()", ServiceConfig.Postgres.Schema, ServiceConfig.Postgres.StudentAssignmentsScoreTable)
	if assignScore.Score != 0 {
		query = query + fmt.Sprintf(",score=%v", assignScore.Score)
	}
	if assignScore.MaximumScore != 0 {
		query = query + fmt.Sprintf(",maximum_score=%v", assignScore.MaximumScore)
	}

	query = query + fmt.Sprintf(" where student_id=%v and assignment_id=%v;", assignScore.StudentID, assignScore.AssignmentID)
	fmt.Println("assignScore StudentAssignmentsScore) Update Query: ", query)

	if _, err := conn.Exec(ctx, query); err != nil {
		fmt.Println("assignScore StudentAssignmentsScore Update, Error DB Insert : ", err)
		log.Error(ctx, "assignScore StudentAssignmentsScore Update, Error DB Insert :", err)
		return err
	}
	return nil
}
