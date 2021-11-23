create database simpleLogin;

create table users (
	uuid varchar(36) not null,
	email varchar(100) not null,
	password varchar(72) not null,
	created_at timestamp not null,
	primary key uuid (uuid)
);
