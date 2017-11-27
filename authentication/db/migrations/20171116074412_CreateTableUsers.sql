
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table Users (
  id            SERIAL     primary key,
  email char(40),
  password char(60),
  UNIQUE(email)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table Users;
