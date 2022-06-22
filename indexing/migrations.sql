create database fakerdata;

create table users (
    id int primary key auto_increment,
    name varchar(100),
    username varchar(50),
    email varchar(50),
    phone varchar(30),
    uuid varchar(255)
);

create table usersdata (
    id int primary key auto_increment,
    user_id int,
    user_data text,
    jwt text,
    sentence text,
    constraint fk_users foreign key (user_id) references users(id)
);