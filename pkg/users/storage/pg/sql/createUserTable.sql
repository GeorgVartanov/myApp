CREATE TABLE If NOT EXISTS app_user(
    id serial NOT NULL,
    email varchar(1024) UNIQUE NOT NULL,
    password varchar(1024) NOT NULL,
    display_name varchar(255),
    CONSTRAINT "app_user_pk" PRIMARY KEY ("id")
)with (
OIDS=FALSE
);