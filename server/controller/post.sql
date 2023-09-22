select * from users where id = '64208a79b4cbdc001e064a1a00000000';

select * from follow;

select id from posts where title = '蟲蟲剋星';



select name
from posts 
inner join users
on posts.posted_by = users.id
where title = '蟲蟲剋星' 
;


ALTER TABLE posts
ADD COLUMN user_name varchar(255) NOT NULL;

ALTER TABLE posts
DROP COLUMN name;


# insert name to posts

select users.name, users.id from users
inner join posts
on users.id = posts.posted_by;

select name from users;

SELECT E'\\xDEADBEEF' from posts;


SELECT *
FROM posts
WHERE id = UNHEX('60ccb0568e68901844d0929300000000');



BEGIN;

UPDATE posts
INNER JOIN users ON posts.posted_by = users.id
SET posts.user_name = users.name;

select * from posts;

COMMIT;


select * from posts;