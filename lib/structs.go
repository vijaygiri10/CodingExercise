package lib

import "time"

//ServiceConfig  Configuration Structs
var ServiceConfig Configuration

//Configuration ...
type Configuration struct {
	Service
	Postgres
}

//Service ...
type Service struct {
	Port       string `yaml:"port"`
	ProjectID  string `yaml:"project_id"`
	LogName    string `yaml:"log_name"`
	PixelImage string `yaml:"pixel_image"`
}

//ServiceConfiguration ...
type ServiceConfiguration struct {
	Env map[string]Service `yaml:"Environment"`
}

//Postgres ...
type Postgres struct {
	User                         string `yaml:"user"`
	Password                     string `yaml:"password"`
	Host                         string `yaml:"host"`
	Port                         int64  `yaml:"port"`
	SSLMode                      string `yaml:"sslmode"`
	DBName                       string `yaml:"dbname"`
	Schema                       string `yaml:"schema"`
	PoolMaxConns                 int64  `yaml:"pool_max_conns"`
	StudentsTable                string `yaml:"students_table"`
	AssignmentsTable             string `yaml:"assignments_table"`
	StudentAssignmentsScoreTable string `yaml:"student_assignments_score"`
}

//PostgresConfiguration ...
type PostgresConfiguration struct {
	Env map[string]Postgres `yaml:"Environment"`
}

//Students ...
type Students struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	EnrollmentAT time.Time `json:"enrollment_at"`
}

//Assignments ...
type Assignments struct {
	AssignmentsID  int64     `json:"id"`
	AssignmentName string    `json:"name"`
	MaximumScore   int64     `json:"maximum_score"`
	RecordedAT     time.Time `json:"recorded_at"`
	UpdatedAT      time.Time `json:"updated_at"`
}

//StudentAssignmentsScore ...
type StudentAssignmentsScore struct {
	ID             int64     `json:"id"`
	StudentID      int64     `json:"student_id"`
	AssignmentID   int64     `json:"assignment_id"`
	AssignmentName string    `json:"assignment_name"`
	Score          int64     `json:"score"`
	MaximumScore   int64     `json:"maximum_score"`
	RecordedAT     time.Time `json:"recorded_at"`
	UpdatedAT      time.Time `json:"updated_at"`
}
