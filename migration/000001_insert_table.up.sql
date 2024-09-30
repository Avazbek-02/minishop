create type category as enum('ovqat','texnika','telefon','avtamabil','uy');

create table if not exists shop(
    id uuid primary key not null,
    name varchar(250) not null,
    img_url varchar(300) not null,
    categorys category not null,
    user_name varchar(100) not null, 
    user_phone varchar(20) not null,
    created_at timestamp default now() not null, 
    updated_at timestamp default now() not null, 
    deleted_at bigint default 0 not null
);


INSERT INTO shop (name, img_url, categorys, user_name, user_phone)
VALUES ('My Shop', 'https://example.com/image.jpg', 'texnika', 'John Doe', '+1234567890');
