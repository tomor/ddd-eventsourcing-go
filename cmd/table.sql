

create table if not exists eventstore
(
	id serial not null,
	event_name varchar(255) not null,
	stream_id varchar(255) not null,
	stream_version integer default 0 not null,
	payload jsonb default '{}'::jsonb not null,
	occurred_at timestamp with time zone not null
);

--alter table eventstore owner to esworkshop;
create unique index if not exists eventstore_id_uindex
	on eventstore (id);
create unique index if not exists stream_unique
	on eventstore (stream_id, stream_version);
create index if not exists eventstore_event_name_idx
	on eventstore (event_name);
create index if not exists eventstore_occurred_at_idx
	on eventstore (occurred_at);
