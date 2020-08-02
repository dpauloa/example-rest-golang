CREATE TABLE phone_book (
    id serial NOT NULL,
    first_name varchar(30) NOT NULL,
    last_name varchar(30) NOT NULL,
    phone_number varchar(11) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);