CREATE TABLE IF NOT EXISTS users(
    id varchar(250) primary key,
    first_name varchar(250),
    last_name varchar(250),
    mobile varchar(50),
    password varchar(50),
    session_id varchar(250)
);
