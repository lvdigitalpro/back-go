create schema back;

create table back.users (
                            id text primary key unique,
                            role text not null,
                            name text not null,
                            lastname text not null,
                            ir text not null unique,
                            email text not null unique,
                            password text not null,
                            enterpriseName text,
                            nrle text unique,
                            projects text[],
                            created_at text not null,
                            updated_at text not null
);

create unique index users_id_uindex on back.users (id);

create unique index users_email_uindex on back.users (email);

create unique index users_ir_uindex on back.users (ir);

create unique index users_nrle_uindex on back.users (nrle);

create type Statuss as enum ('QUEUE', 'IN_REVIEW','IN_PROGRESS','DONE');

create type Typee as enum ('WEB', 'MOBILE','DESKTOP','SYSTEM','UI_UX','OTHER');

create type TaskType as enum ('SEARCH', 'DESIGN','DEVELOPMENT','TEST','OTHER');

create type Task as (
                        id text,
                        type TaskType,
                        name text,
                        description text,
                        status Statuss,
                        owner text,
                        created_at text,
                        updated_at text
                    );

create table back.projects (
                               id text primary key unique,
                               type Typee not null,
                               owner text not null,
                               name text not null,
                               description text not null,
                               tasks Task[],
                               status Statuss not null,
                               created_at text not null,
                               updated_at text not null,
                               startDate text not null,
                               endDate text not null
);

create unique index projects_id_uindex on back.projects (id);

alter table back.projects add constraint projects_users_fk foreign key (back.projects.owner) references back.users (id);

alter table back.users add constraint users_projects_fk foreign key (EACH ELEMENT OF back.users.projects) references back.projects (id);