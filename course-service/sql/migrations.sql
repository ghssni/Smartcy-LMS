-- Drop and recreate the database
DROP DATABASE IF EXISTS smartcy_dev_1;
CREATE DATABASE smartcy_dev_1;

-- Connect to the database
\c smartcy_dev;

-- Enum types
CREATE TYPE lesson_type AS ENUM ('video', 'article');
CREATE TYPE payment_status_type AS ENUM ('paid', 'pending', 'failed');
CREATE TYPE transaction_status_type AS ENUM ('paid', 'unpaid', 'refund');
CREATE TYPE payment_method_type AS ENUM ('credit_card', 'bank_transfer', 'e_wallet');
CREATE TYPE course_category AS ENUM ('frontend', 'backend', 'fullstack', 'devops');

-- Courses table
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

-- Enrollments table
CREATE TABLE enrollments
(
    id             SERIAL PRIMARY KEY,
    course_id      INTEGER             NOT NULL REFERENCES courses (id),
    student_id     VARCHAR             NOT NULL, -- References MongoDB user ID
    payment_status payment_status_type NOT NULL,
    enrolled_at    TIMESTAMP DEFAULT NOW()
);

-- Payments table
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

-- Lessons table
CREATE TABLE lessons
(
    id          SERIAL PRIMARY KEY,
    course_id   INTEGER        NOT NULL REFERENCES courses (id),
    title       VARCHAR        NOT NULL,
    content_url VARCHAR        NOT NULL,
    lesson_type lesson_type,
    sequence    INTEGER        NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    CONSTRAINT lessons_course_sequence_unique UNIQUE (course_id, sequence)
);


-- Reviews table
CREATE TABLE reviews
(
    id          SERIAL PRIMARY KEY,
    course_id   INTEGER NOT NULL REFERENCES courses (id),
    student_id  VARCHAR NOT NULL, -- References MongoDB user ID
    rating      INTEGER NOT NULL,
    review_text TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  TIMESTAMP
    Constraint reviews_course_student_unique UNIQUE (course_id, student_id)
);

-- Add CONSTRAINT reviews_course_student_unique UNIQUE (course_id, student_id) to reviews table
-- ALTER TABLE reviews ADD CONSTRAINT reviews_course_student_unique UNIQUE (course_id, student_id );

-- User activity log table
CREATE TABLE user_activity_log
(
    id                 SERIAL PRIMARY KEY,
    user_id            VARCHAR NOT NULL, -- References MongoDB user ID
    course_id          INTEGER NOT NULL REFERENCES courses (id),
    activity_type      VARCHAR NOT NULL,
    activity_timestamp TIMESTAMP
);

-- Learning progress table
CREATE TABLE learning_progress
(
    id            SERIAL PRIMARY KEY,
    enrollment_id INTEGER NOT NULL REFERENCES enrollments (id),
    lesson_id     INTEGER NOT NULL REFERENCES lessons (id),
    status        BOOLEAN,
    last_accessed TIMESTAMP,
    completed_at  TIMESTAMP
);

-- Assessments table
CREATE TABLE assessments
(
    id              SERIAL PRIMARY KEY,
    enrollment_id   INTEGER NOT NULL REFERENCES enrollments (id),
    score           INTEGER NOT NULL,
    assessment_type VARCHAR NOT NULL,
    taken_at        TIMESTAMP
);

-- Certificates table
CREATE TABLE certificates
(
    id              SERIAL PRIMARY KEY,
    enrollment_id   INTEGER NOT NULL REFERENCES enrollments (id),
    issued_at       TIMESTAMP DEFAULT NOW(),
    certificate_url VARCHAR NOT NULL
);

-- Seeding data for courses
INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category)
VALUES ('Introduction to Frontend Development', 'Learn the basics of HTML, CSS, and JavaScript.', 199.99,
        'https://example.com/frontend-thumbnail.jpg', 'user_id_1', 'frontend'),
       ('Advanced Backend Development', 'Deep dive into Node.js, Express, and MongoDB.', 299.99,
        'https://example.com/backend-thumbnail.jpg', 'user_id_2', 'backend'),
       ('Fullstack Developer Bootcamp', 'Become a fullstack developer with React and Node.js.', 399.99,
        'https://example.com/fullstack-thumbnail.jpg', 'user_id_3', 'fullstack'),
       ('DevOps Essentials', 'Learn the principles of DevOps and cloud infrastructure.', 249.99,
        'https://example.com/devops-thumbnail.jpg', 'user_id_4', 'devops'),
       ('Introduction to UI/UX Design', 'Discover the fundamentals of UI/UX design principles.', 149.99,
        'https://example.com/uiux-thumbnail.jpg', 'user_id_5', 'frontend');

-- Seeding data for enrollments
INSERT INTO enrollments (course_id, student_id, payment_status)
VALUES (1, 'student_id_1', 'paid'),
       (1, 'student_id_2', 'pending'),
       (2, 'student_id_3', 'paid'),
       (3, 'student_id_4', 'failed'),
       (4, 'student_id_5', 'paid');

-- Seeding data for payments
-- Seeding data for payments
INSERT INTO payments (enrollment_id, amount, transaction_status, transaction_date, invoice_id, payment_method,
                      payment_provider, description)
VALUES (1, 199.99, 'paid', NOW(), 'invoice_1', 'credit_card', 'PaymentGatewayX',
        'Payment for Frontend Development course.'),
       (2, 199.99, 'unpaid', NOW(), 'invoice_2', 'bank_transfer', 'PaymentGatewayY',
        'Unpaid payment for Frontend Development course.'), -- 'pending' diganti 'unpaid'
       (3, 299.99, 'paid', NOW(), 'invoice_3', 'e_wallet', 'PaymentGatewayZ',
        'Payment for Backend Development course.'),
       (4, 399.99, 'refund', NOW(), 'invoice_4', 'credit_card', 'PaymentGatewayX',
        'Refund issued for Fullstack Bootcamp.'),
       (5, 249.99, 'paid', NOW(), 'invoice_5', 'bank_transfer', 'PaymentGatewayY',
        'Payment for DevOps Essentials course.');

-- Seeding data for lessons
-- Seeding data for lessons
INSERT INTO lessons (course_id, title, content_url, lesson_type, sequence)
VALUES
    (1, 'Introduction to Frontend Development', 'https://example.com/lesson1', 'video', 1),
    (1, 'HTML Basics', 'https://example.com/lesson2', 'article', 2),
    (2, 'Backend Development Overview', 'https://example.com/lesson3', 'video', 1),
    (2, 'Database Design', 'https://example.com/lesson4', 'article', 2),
    (3, 'Fullstack Bootcamp Introduction', 'https://example.com/lesson5', 'video', 1),
    (3, 'Working with APIs', 'https://example.com/lesson6', 'article', 2),
    (4, 'DevOps Fundamentals', 'https://example.com/lesson7', 'video', 1),
    (4, 'CI/CD Pipelines', 'https://example.com/lesson8', 'article', 2),
    (5, 'Introduction to Cloud Computing', 'https://example.com/lesson9', 'video', 1),
    (5, 'Kubernetes Basics', 'https://example.com/lesson10', 'article', 2);

-- Seeding data for reviews
INSERT INTO reviews (course_id, student_id, rating, review_text)
VALUES (1, 'student_id_1', 5, 'Great course! Learned a lot about frontend development.'),
       (1, 'student_id_2', 4, 'Good content but needs more examples.'),
       (2, 'student_id_3', 5, 'Excellent explanation of backend concepts.'),
       (3, 'student_id_4', 3, 'Course was okay, but the pace was a bit slow.'),
       (4, 'student_id_5', 4, 'Informative course, highly recommend for beginners.');

-- Seeding data for user activity log
INSERT INTO user_activity_log (user_id, course_id, activity_type, activity_timestamp)
VALUES ('student_id_1', 1, 'enrolled', NOW()),
       ('student_id_2', 1, 'completed_lesson', NOW()),
       ('student_id_3', 2, 'enrolled', NOW()),
       ('student_id_4', 3, 'viewed', NOW()),
       ('student_id_5', 4, 'enrolled', NOW());

-- Seeding data for learning progress
INSERT INTO learning_progress (enrollment_id, lesson_id, status, last_accessed, completed_at)
VALUES (1, 1, TRUE, NOW(), NOW()),
       (1, 2, FALSE, NOW(), NULL),
       (3, 1, TRUE, NOW(), NOW()),
       (5, 1, TRUE, NOW(), NOW()),
       (5, 2, FALSE, NOW(), NULL),
       (5, 4, TRUE, NOW(), NOW()),
       (5, 2, FALSE, NOW(), NULL);

-- Seeding data for assessments
INSERT INTO assessments (enrollment_id, score, assessment_type, taken_at)
VALUES (1, 85, 'quiz', NOW()),
       (2, 90, 'final_exam', NOW()),
       (3, 70, 'quiz', NOW()),
       (4, 95, 'final_exam', NOW()),
       (5, 80, 'quiz', NOW());

-- Seeding data for certificates
INSERT INTO certificates (enrollment_id, issued_at, certificate_url)
VALUES (1, NOW(), 'https://example.com/certificate1'),
       (2, NOW(), 'https://example.com/certificate2'),
       (3, NOW(), 'https://example.com/certificate3'),
       (4, NOW(), 'https://example.com/certificate4'),
       (5, NOW(), 'https://example.com/certificate5');
