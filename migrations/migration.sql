CREATE TABLE if not exists USERS(
    id serial primary key,
    name varchar(255),
    password varchar(255),
    role int
);

CREATE TABLE CATEGORIES(
    id serial primary key,
    name varchar(255),
    description varchar(255),
    parent_id int null
)