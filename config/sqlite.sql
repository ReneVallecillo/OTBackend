BEGIN TRANSACTION;
CREATE TABLE `user` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`first_name`	TEXT,
	`last_name`	TEXT,
	`email`	TEXT UNIQUE,
	`password`	TEXT,
	`status_id`	INTEGER,
	`created_at`	TEXT,
	`updated_at`	TEXT
);
CREATE TABLE `attendance` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`checkin_time`	TEXT,
	`checkout_time`	TEXT
);
COMMIT;
