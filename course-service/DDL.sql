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
