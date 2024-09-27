DROP DATABASE IF EXISTS smartcy_dev;
CREATE DATABASE smartcy_dev;


CREATE TYPE lesson_type AS ENUM ('video', 'article');
CREATE TYPE payment_status_type AS ENUM ('paid', 'pending', 'failed');
CREATE TYPE transaction_status_type AS ENUM ('paid', 'unpaid', 'refund');
CREATE TYPE payment_method_type AS ENUM ('credit_card', 'bank_transfer', 'e_wallet');
CREATE TYPE course_category AS ENUM ('frontend', 'backend', 'fullstack', 'devops');

-- Removed users table as it will be in MongoDB

CREATE TABLE courses
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR         NOT NULL,
    description   TEXT            NOT NULL,
    price         DECIMAL(10, 2)  NOT NULL,
    thumbnail_url VARCHAR         NOT NULL,
    instructor_id VARCHAR         NOT NULL, -- References MongoDB user ID
    category      course_category NOT NULL,
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW(),
    deleted_at    TIMESTAMP
);

CREATE TABLE enrollments
(
    id             SERIAL PRIMARY KEY,
    course_id      INTEGER             NOT NULL REFERENCES courses (id),
    student_id     VARCHAR             NOT NULL, -- References MongoDB user ID
    payment_status payment_status_type NOT NULL,
    enrolled_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payments
(
    id                 SERIAL PRIMARY KEY,
    enrollment_id      INTEGER                 NOT NULL REFERENCES enrollments (id),
    amount             DECIMAL(10, 2)          NOT NULL,
    transaction_status transaction_status_type NOT NULL,
    transaction_date   TIMESTAMP               NOT NULL,
    invoice_id         VARCHAR                 NOT NULL,
    payment_method     payment_method_type     NOT NULL,
    payment_provider   VARCHAR                 NOT NULL,
    description        TEXT,
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW(),
    deleted_at         TIMESTAMP
);

CREATE TABLE lessons
(
    id          SERIAL PRIMARY KEY,
    course_id   INTEGER NOT NULL REFERENCES courses (id),
    title       VARCHAR NOT NULL,
    content_url VARCHAR NOT NULL,
    lesson_type lesson_type,
    sequence    INTEGER NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE reviews
(
    id          SERIAL PRIMARY KEY,
    course_id   INTEGER NOT NULL REFERENCES courses (id),
    student_id  VARCHAR NOT NULL, -- References MongoDB user ID
    rating      INTEGER NOT NULL,
    review_text TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_activity_log
(
    id                 SERIAL PRIMARY KEY,
    user_id            VARCHAR NOT NULL, -- References MongoDB user ID
    course_id          INTEGER NOT NULL REFERENCES courses (id),
    activity_type      VARCHAR NOT NULL,
    activity_timestamp TIMESTAMP DEFAULT NOW()
);

CREATE TABLE learning_progress
(
    id            SERIAL PRIMARY KEY,
    enrollment_id INTEGER NOT NULL REFERENCES enrollments (id),
    lesson_id     INTEGER NOT NULL REFERENCES lessons (id),
    status        BOOLEAN,
    completed_at  TIMESTAMP
);

CREATE TABLE assessments
(
    id              SERIAL PRIMARY KEY,
    enrollment_id   INTEGER NOT NULL REFERENCES enrollments (id),
    score           INTEGER NOT NULL,
    assessment_type VARCHAR NOT NULL,
    taken_at        TIMESTAMP
);

CREATE TABLE certificates
(
    id              SERIAL PRIMARY KEY,
    enrollment_id   INTEGER NOT NULL REFERENCES enrollments (id),
    issued_at       TIMESTAMP DEFAULT NOW(),
    certificate_url VARCHAR NOT NULL
);

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Frontend Web Development', 'Learn how to build responsive websites using HTML, CSS, and JavaScript.', 49.99, 'http://example.com/thumb1.jpg', 'mongodb_instructor_1', 'frontend', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Backend Development with Node.js', 'Master backend development using Node.js and Express.', 59.99, 'http://example.com/thumb2.jpg', 'mongodb_instructor_2', 'backend', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Fullstack Web Development', 'Become a fullstack web developer by learning both frontend and backend skills.', 79.99, 'http://example.com/thumb3.jpg', 'mongodb_instructor_3', 'fullstack', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('DevOps Fundamentals', 'Learn the basics of DevOps and how to automate your infrastructure.', 69.99, 'http://example.com/thumb4.jpg', 'mongodb_instructor_4', 'devops', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Advanced Frontend Techniques', 'Master advanced frontend techniques such as React and modern CSS.', 89.99, 'http://example.com/thumb5.jpg', 'mongodb_instructor_5', 'frontend', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Backend API Development with Go', 'Build robust backend APIs using Go.', 99.99, 'http://example.com/thumb6.jpg', 'mongodb_instructor_6', 'backend', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Fullstack JavaScript Developer', 'Become a fullstack JavaScript developer using MERN stack.', 79.99, 'http://example.com/thumb7.jpg', 'mongodb_instructor_7', 'fullstack', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('DevOps with Docker and Kubernetes', 'Learn how to use Docker and Kubernetes for DevOps workflows.', 89.99, 'http://example.com/thumb8.jpg', 'mongodb_instructor_8', 'devops', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('Introduction to Frontend Development', 'Get started with frontend web development using HTML, CSS, and JavaScript.', 39.99, 'http://example.com/thumb9.jpg', 'mongodb_instructor_9', 'frontend', NOW(), NOW());

INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at)
VALUES ('DevOps Engineering with AWS', 'Learn DevOps practices using AWS cloud platform.', 109.99, 'http://example.com/thumb10.jpg', 'mongodb_instructor_10', 'devops', NOW(), NOW());

