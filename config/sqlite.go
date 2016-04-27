package config

//InitSqls has the init sql for each table
var InitSqls = []string{
	`
    CREATE TABLE IF NOT EXISTS 'user' (
	'id'	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	'first_name'	TEXT,
	'last_name'	TEXT,
	'email'	TEXT UNIQUE,
	'password'	TEXT,
	'status_id'	INTEGER,
	'created_at'	TEXT,
	'updated_at'	TEXT
);`,
	`CREATE TABLE IF NOT EXISTS 'attendance' (
	'id'	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	'checkin_time'	TEXT,
	'checkout_time'	TEXT
);`,
}
