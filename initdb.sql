CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20),
    content VARCHAR(255),
    status SMALLINT
    );

INSERT INTO messages (phone_number, content, status) VALUES
    ('1234567890', 'Hello, this is a test message', 0),
    ('0987654321', 'This is another test message', 0),
    ('5555555555', 'Test message number three', 0),
    ('6666666666', 'Test message number four', 0);ges (phone_number, content, status) VALUES
                                                         ('1234567890', 'Hello, this is a test message', 0),
                                                         ('0987654321', 'This is another test message', 0),
                                                         ('5555555555', 'Test message number three', 0),
                                                         ('6666666666', 'Test message number four', 0);