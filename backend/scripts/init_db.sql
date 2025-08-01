-- naming convention: snake case (https://www.postgresql.org/docs/7.0/syntax525.htm)

CREATE DATABASE avhbs;

\c avhbs;

CREATE TABLE category(
    category_id serial PRIMARY KEY,
    name   varchar(20) NOT NULL,
    enabled bool NOT NULL DEFAULT TRUE,
    icon varchar(20),
    type varchar(20)
);

CREATE TABLE account
(
    account_id serial PRIMARY KEY,
    first_name varchar(20),
    nickname   varchar(20),
    last_name  varchar(20),
    email      varchar(40),
    phone      varchar(20),
    balance    NUMERIC(5, 2) NOT NULL,
    max_debt   NUMERIC(5, 2) NOT NULL,
    category   integer       NOT NULL,
    enabled    bool                   DEFAULT TRUE,
    created_at timestamp     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (category) REFERENCES category (category_id)
        ON DELETE NO ACTION
);


CREATE TABLE account_option
(
    account integer     NOT NULL,
    key     varchar(40) NOT NULL,
    value   varchar(40) NOT NULL,
    PRIMARY KEY (account, key),
    FOREIGN KEY (account) REFERENCES account (account_id)
        ON DELETE CASCADE
);
