create table if not exists users
(
	id bigserial not null
		constraint users_pk
			primary key,
	name varchar not null,
	email varchar not null,
	password varchar not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	deleted_at timestamp
);

create table if not exists players
(
	id bigserial not null
		constraint players_pk
			primary key,
	user_id bigint
		constraint players_users_id_fk
			references users
				on update cascade on delete set null,
	name varchar not null,
	last_received_at timestamp,
	created_at timestamp not null,
	updated_at timestamp not null
);

create table if not exists roles
(
	id bigserial not null
		constraint roles_pk
			primary key,
	name varchar not null,
	description text,
	created_at timestamp not null,
	updated_at timestamp not null,
	deleted_at timestamp
);

create table if not exists permissions
(
	id varchar not null
		constraint permissions_pk
			primary key,
	label varchar not null,
	description text not null,
	created_at timestamp not null,
	updated_at timestamp not null
);

