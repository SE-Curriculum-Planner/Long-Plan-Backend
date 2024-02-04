CREATE TYPE GROUP_TYPE AS ENUM ('Core','Major Required','Major Elective','Learner Person','Co-Creator','Active Citizen','Free Elective');
CREATE TYPE MAJOR AS ENUM ('CPE','ISNE')

CREATE TABLE students (
    student_id INT PRIMARY KEY,
    major MAJOR NOT NULL,
    last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE course_enrolled (
    course_no INT NOT NULL,
    student_id INT NOT NULL,
    "year" INT NOT NULL,
    semester INT NOT NULL,
    grade VARCHAR(2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY(course_no,student_id,"year",semester),

    CONSTRAINT FK_enrolled_course FOREIGN KEY(course_no)
        REFERENCES courses(course_no),
    CONSTRAINT FK_enrolled_student FOREIGN KEY(student_id)
        REFERENCES students(student_id)
);

CREATE TABLE curriculums (
    id BIGSERIAL PRIMARY KEY,
    curriculum_program VARCHAR(255) NOT NULL,
    "year" INT NOT NULL,
    is_coop_plan BOOLEAN NOT NULL,
    require_credits INT NOT NULL,
    free_elective_credits INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE curriculum_details (
    curriculum_id INT PRIMARY KEY NOT NULL,
    group_name GROUP_TYPE NOT NULL,
    require_credits INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT FK_detail_curriculum FOREIGN KEY(curriculum_id)
        REFERENCES curriculums(id)
);

CREATE TABLE courses(
    course_no INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    recommend_semester INT NOT NULL,
    recommend_year INT NOT NULL,
    credit INT NOT NULL,
    group_name GROUP_TYPE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
);

CREATE TABLE course_co_requisites (
    course_no INT PRIMARY KEY,
    co_requisite_no INT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY(course_no,co_requisite_no),

    CONSTRAINT FK_co_requisite_course FOREIGN KEY(course_no)
        REFERENCES courses(course_no),
    CONSTRAINT FK_course_co_requisite FOREIGN KEY(co_requisite_no)
        REFERENCES courses(course_no)
)

CREATE TABLE course_pre_requisites (
    course_no INT PRIMARY KEY,
    pre_requisite_no INT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY(course_no,pre_requisite_no),

    CONSTRAINT FK_pre_requisite_course FOREIGN KEY(course_no)
        REFERENCES courses(course_no),
    CONSTRAINT FK_course_pre_requisite FOREIGN KEY(pre_requisite_no)
        REFERENCES courses(course_no)
)