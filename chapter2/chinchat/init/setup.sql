create database chinchat;
\c chinchat

-- drop table posts;
-- drop table threads;
-- drop table sessions;
-- drop table users;


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null  
);


insert into users values(1, 'user1', 'tarotarotaro', 'taro@taro', 'passpass', '2022-05-04');
insert into threads values(1, 'id1', 'topictopictopic', 1, '2022-05-04');
insert into threads values(2, 'id2', 'hogehogehoge', 1, '2022-05-01');
