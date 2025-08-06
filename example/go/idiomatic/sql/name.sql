-- -----------------------------------------------------------------------------------------
-- Prefix
-- -----------------------------------------------------------------------------------------  
CREATE TABLE "person_name_prefix" (
	"id"	TEXT,
	"text"	TEXT NOT NULL,
	"description"	TEXT,
	PRIMARY KEY("id")
);

CREATE TABLE "person_name_prefix_abbr" (
	"id"	TEXT,
	"prefix_id"	TEXT NOT NULL,
	"abbr"	TEXT NOT NULL,
	"index"	INTEGER NOT NULL,
    PRIMARY KEY("id"),
	FOREIGN KEY("prefix_id") REFERENCES "person_name_prefix"("id") ON DELETE CASCADE	
);

-- -----------------------------------------------------------------------------------------
-- Suffix
-- -----------------------------------------------------------------------------------------  

CREATE TABLE "person_name_suffix" (
	"id"	TEXT,
	"text"	TEXT NOT NULL,
	"description"	TEXT,
	PRIMARY KEY("id")
);

CREATE TABLE "person_name_suffix_abbr" (
	"id"	TEXT,
	"suffix_id"	TEXT NOT NULL,
	"abbr"	TEXT NOT NULL,
	"index"	INTEGER NOT NULL,
    PRIMARY KEY("id"),
	FOREIGN KEY("suffix_id") REFERENCES "person_name_suffix"("id") ON DELETE CASCADE
);

-- -----------------------------------------------------------------------------------------
-- Name
-- -----------------------------------------------------------------------------------------  

CREATE TABLE "person_name_first" (
	"id"	TEXT,
	"name_id"	TEXT NOT NULL,
	"name"	INTEGER NOT NULL,
    "index"	INTEGER NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("name_id") REFERENCES "person_name"("id") ON DELETE CASCADE
);

CREATE TABLE "person_name_middle" (
	"id"	TEXT,
	"name_id"	TEXT NOT NULL,
	"name"	INTEGER NOT NULL,
    "index"	INTEGER NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("name_id") REFERENCES "person_name"("id") ON DELETE CASCADE
);

CREATE TABLE "person_name_last" (
	"id"	TEXT,
	"name_id"	TEXT NOT NULL,
	"name"	INTEGER NOT NULL,
    "index"	INTEGER NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("name_id") REFERENCES "person_name"("id") ON DELETE CASCADE
);

CREATE TABLE "person_name_person_name_prefix" (
	"id"	TEXT,
	"name_id"	TEXT NOT NULL,
	"prefix_id"	INTEGER NOT NULL,
    "index"	INTEGER NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("name_id") REFERENCES "person_name"("id") ON DELETE CASCADE,
    FOREIGN KEY("prefix_id") REFERENCES "person_name_prefix"("id") ON DELETE CASCADE
);

CREATE TABLE "person_name_person_name_suffix" (
	"id"	TEXT,
	"name_id"	TEXT NOT NULL,
	"suffix_id"	INTEGER NOT NULL,
    "index"	INTEGER NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("name_id") REFERENCES "person_name"("id") ON DELETE CASCADE,
    FOREIGN KEY("suffix_id") REFERENCES "person_name_suffix"("id") ON DELETE CASCADE
);

CREATE TABLE "person_name" (
    "id"	TEXT,
    PRIMARY KEY("id"),
);
