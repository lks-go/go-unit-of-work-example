CREATE TABLE verifications
(
    user_id INTEGER NOT NULL,
    code CHAR(6) NOT NULL,
    verified BOOL
);