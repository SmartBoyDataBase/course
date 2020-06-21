package model

import (
	"course/infrastructure"
	"fmt"
)

type Course struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Credit    uint8  `json:"credit"`
	CollegeId uint64 `json:"college_id"`
}

func Get(id uint64) (Course, error) {
	result := Course{
		Id: id,
	}
	row := infrastructure.DB.QueryRow(`
	SELECT name, credit, college_id
	FROM course
	WHERE id=$1;
	`, id)
	err := row.Scan(&result.Name, &result.Credit, &result.CollegeId)
	return result, err
}

func Create(course Course) (Course, error) {
	fmt.Println(course)
	row := infrastructure.DB.QueryRow(`
	INSERT INTO course(name, credit,college_id)
	VALUES ($1, $2, $3)
	RETURNING id;`, course.Name, course.Credit, course.CollegeId)
	err := row.Scan(&course.Id)
	return course, err
}

func Delete(id uint64) error {
	_, err := infrastructure.DB.Exec(`
	DELETE FROM course
	WHERE id=$1;
	`, id)
	return err
}

func All() ([]Course, error) {
	rows, err := infrastructure.DB.Query(`
	SELECT id,name,credit,college_id
	FROM course;
	`)
	if err != nil {
		return nil, err
	}
	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.Id, &course.Name, &course.Credit, &course.CollegeId)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
