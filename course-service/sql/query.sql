SELECT * FROM lessons WHERE course_id = 1 ORDER BY sequence;

SELECT * FROM lessons WHERE id = 1;

-- soft delete lesson id 3
UPDATE lessons SET deleted_at = NOW() WHERE id = 3;

UPDATE courses SET deleted_at = NULL WHERE id = 1;

-- undelete lesson id 3
UPDATE lessons SET deleted_at = NULL WHERE id = 1;

SELECT * FROM lessons WHERE sequence = 1 AND course_id = 1 ORDER BY sequence;

SELECT id, title, content_url, lesson_type, sequence, course_id, created_at, updated_at FROM lessons WHERE course_id = 1 AND title ILIKE '%' || 'Basic' || '%' AND deleted_at IS NULL ORDER BY sequence;

SELECT * FROM enrollments e
         JOIN learning_progress lp ON e.id = lp.enrollment_id
         WHERE e.id = 1;

SELECT * FROM enrollments;

-- update student id where student id is student_id
UPDATE enrollments SET student_id = '66f5752ef5bb513d8c8de1cb' WHERE student_id = 'student_id_1';

SELECT COUNT(*) AS total_completed FROM learning_progress WHERE enrollment_id = 1 AND completed_at IS NOT NULL ;

SELECT enrollment_id, COUNT(*) FROM learning_progress WHERE status = true GROUP BY enrollment_id ;


SELECT lp.enrollment_id, COUNT(*) AS total_completed
FROM learning_progress lp
         JOIN enrollments e ON lp.enrollment_id = e.id
WHERE e.student_id = '66f5752ef5bb513d8c8de1cb' AND lp.status = true
GROUP BY lp.enrollment_id;

SELECT TRUE FROM enrollments WHERE id = 1 AND student_id = '66f5752ef5bb513d8c8de1cb';

SELECT * FROM learning_progress WHERE enrollment_id = 1;

INSERT INTO learning_progress (enrollment_id, lesson_id, status, last_accessed, completed_at) VALUES (1,1,true,NOW(),NOW()) WHERE (SELECT TRUE FROM enrollments WHERE id = $1 AND student_id = $6)

SELECT id, enrollment_id, lesson_id, status, last_accessed, completed_at FROM learning_progress WHERE enrollment_id = 1;

-- Get review by course_id and student_id
SELECT * FROM reviews WHERE course_id = 1 AND student_id = '66f5752ef5bb513d8c8de1cb';