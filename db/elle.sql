DROP TABLE elle;
CREATE TABLE elle (
  id int PRIMARY KEY,
  content_id int ,
  type varchar (50) NOT NULL,
  date_unix int,
  page int,
  url VARCHAR(300),
  scan_date TIMESTAMP
);
TRUNCATE elle;