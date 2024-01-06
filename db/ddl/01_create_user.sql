CREATE USER 'tasktree_user'@'%' IDENTIFIED BY 'tasktree_pass';
GRANT SELECT,INSERT,UPDATE,DELETE,EXECUTE,SHOW VIEW ON diary.* TO 'tasktree_user'@'%';
