CREATE TABLE `images` 
(
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "path" varchar(1024) NOT NULL,
  "content_type" varchar(50) NOT NULL,
  "size" bigint NOT NULL,
  "created_at" datetime DEFAULT CURRENT_TIMESTAMP
);
