CREATE DATABASE IF NOT EXISTS mercor;

USE mercor;

DROP TABLE if exists job;
CREATE TABLE IF NOT EXISTS job(
	id INT NOT NULL AUTO_INCREMENT,
    version int not null,
    jobId VARCHAR(30),
    status enum('active', 'inactive', 'extended'),
    rate float,
    title VARCHAR(50),
    companyId int not null,
    contractorId int not null,
    PRIMARY KEY (id)
);

-- INSERT INTO job 
-- (jobId,version,status,rate,title,companyId,contractorId)
-- VALUES
-- (1,1,"extended",20,"Software Engineer",1,1),
-- (1,2,"active",20,"Software Engineer",1,1),
-- (1,3,"active",15.5,"Software Engineer",1,1);

drop table if exists timeLog;
create table if not exists timeLog(
	id int not null AUTO_INCREMENT primary key,
    timeLogId int not null,
    duration int not null,
    timeStart TIMESTAMP not null,
    timeEnd TIMESTAMP not null,
    type enum ('captured', 'adjusted') default 'captured' not null,
    version int not null,
    jobUid int not null
);

-- insert into timeLog (timeLogId, duration, timeStart, timeEnd, type, version, jobUid)
-- values
-- (1,23,12564,125465,'captured', 1, 1),
-- (1,21,12564,125465,'adjusted', 2, 1),
-- (2,25,12564,125465,'captured', 1, 2),
-- (2,24,12564,125465,'adjusted', 2, 2),
-- (3,25,12564,125465,'captured', 1, 1),
-- (3,24,12564,125465,'adjusted', 2, 1),
-- (4,34,12564,125465,'captured', 1, 1),
-- (4,21,12564,125465,'adjusted', 2, 1),
-- (4,11,12564,125465,'adjusted', 3, 1)
-- ;
