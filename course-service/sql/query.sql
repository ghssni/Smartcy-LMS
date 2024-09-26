SELECT * FROM lessons WHERE course_id = 1 ORDER BY sequence;

SELECT * FROM lessons WHERE id = 1;

-- soft delete lesson id 3
UPDATE lessons SET deleted_at = NOW() WHERE id = 3;

UPDATE courses SET deleted_at = NULL WHERE id = 1;

-- undelete lesson id 3
UPDATE lessons SET deleted_at = NULL WHERE id = 1;

SELECT * FROM lessons WHERE sequence = 1 AND course_id = 1 ORDER BY sequence;

SELECT id, title, content_url, lesson_type, sequence, course_id, created_at, updated_at FROM lessons WHERE course_id = 1 AND title ILIKE '%' || 'Basic' || '%' AND deleted_at IS NULL ORDER BY sequence;