create table todos (
    id serial not null primary key,
    assignee varchar(24) not null,
    title varchar(28) not null,
    summary varchar(28) not null,
    deadline timestamp not null,
    status varchar(24) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

