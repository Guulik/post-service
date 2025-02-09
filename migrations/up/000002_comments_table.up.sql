CREATE TABLE IF NOT EXISTS Comments
(
    id serial primary key,
    created_at timestamp default now(),
    content varchar(2000),
    post int not null,
    reply_to int,
    FOREIGN KEY (post) REFERENCES Posts(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    FOREIGN KEY (reply_to) REFERENCES Comments(id) ON DELETE SET NULL ON UPDATE CASCADE
);