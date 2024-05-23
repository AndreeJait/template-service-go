
-- +migrate Up
INSERT INTO users (email, username, password, full_name, date_of_birth, role)
    VALUES ('andree@mail.com',
            'andree11',
            '$2a$10$YEKbVIkCGKDGnCd5aIZceeycpstta7h067Oa1wX5SW4ztgzg20JvK',
            'Andree Panjaitan',
            '2002-01-01',
            'admin');
-- +migrate Down
