CREATE TABLE affix_category (
  id            TEXT PRIMARY KEY,
  'name'        TEXT NOT NULL,
  'description' TEXT
);

CREATE TABLE affix (
    id INTEGER PRIMARY KEY,
    'name' text NOT NULL,
    kind text NOT NULL,
    'description' text,
    category_id INTEGER NOT NULL
);

-- CREATE TABLE affix_abbreviation (
--     affix_id INTEGER PRIMARY KEY,
--     order INTEGER PRIMARY KEY,
--     abbreviation TEXT NOT NULL
-- );
