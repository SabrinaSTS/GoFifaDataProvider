CREATE TYPE EVENT_TYPE AS enum('male soccer', 'female soccer');

CREATE TABLE events(
    id serial primary key,
    name varchar,
    "type" EVENT_TYPE not null,
    year varchar(4) not null 
);

CREATE TABLE national_teams(
    id serial primary key,
    name varchar,
    country varchar not null
);

CREATE TABLE participations(
    id serial primary key,
    event_id integer not null,
    team_id integer not null,
    CONSTRAINT fk_events_id FOREIGN KEY(event_id) REFERENCES events(id),
    CONSTRAINT fk_national_temas_id FOREIGN KEY(team_id) REFERENCES national_teams(id)
);

INSERT INTO events(name, "type", year) VALUES
('Copa Qatar', 'male soccer', '2022'),
('AU-NZ', 'female soccer', '2023');

INSERT INTO national_teams(name, country) VALUES
('CBF', 'Brasil'),
('AFA', 'Argentina');

