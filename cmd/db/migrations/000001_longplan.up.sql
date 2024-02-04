CREATE TYPE GROUP_TYPE AS ENUM ('Core','Major Required','Major Elective','Learner Person','Co-Creator','Active Citizen');
CREATE TYPE MAJOR AS ENUM ('CPE','ISNE')

CREATE TABLE Student (
    studentId VARCHAR(255) PRIMARY KEY NOT NULL,
    major MAJOR NOT NULL, 
    update_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE CourseEnrolled(
    courseNo VARCHAR(255) PRIMARY KEY NOT NULL,
    studentId VARCHAR(255) PRIMARY KEY NOT NULL,
    "year" INT NOT NULL,
    semester INT NOT NULL,
    credit INT NOT NULL,
    grade VARCHAR(255) NOT NULL
);

CREATE TABLE Curriculum(
    id BIGSERIAL PRIMARY KEY,
    curriculumProgram VARCHAR(255) NOT NULL,
    "year" VARCHAR(255) NOT NULL,
    isCOOPPlan BOOLEAN NOT NULL,
    requireCredits INT NOT NULL,
    freeElectiveCredits INT NOT NULL
);

CREATE TABLE CurriculumDetail(
    curriculumId VARCHAR(255) PRIMARY KEY NOT NULL,
    groupName GROUP_TYPE NOT NULL,
    requireCredits INT NOT NULL
);

CREATE TABLE Course(
    courseNo VARCHAR(255) PRIMARY KEY,
    courseTitleEng VARCHAR(255) NOT NULL,
    recommendSemester INT NOT NULL,
    recommendYear INT NOT NULL,
    credit INT NOT NULL,
    groupName GROUP_TYPE NOT NULL
);

CREATE TABLE CourseCORequisite (
    courseNo VARCHAR(255) PRIMARY KEY,
    COcourseNo VARCHAR(255) PRIMARY KEY,
)

CREATE TABLE CoursePreRequisite (
    courseNo VARCHAR(255) PRIMARY KEY,
    preCourseNo VARCHAR(255) PRIMARY KEY,
)