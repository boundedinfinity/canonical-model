CREATE TABLE affix_category (
  id   INTEGER PRIMARY KEY,
  'name' text    NOT NULL,
  'description'  text
);

CREATE TABLE affix (
    id INTEGER PRIMARY KEY,
    'name' text NOT NULL,
    kind text NOT NULL,
    'description' text,
    category_id INTEGER NOT NULL
)

CREATE TABEL affix_abbreviation (
    affix_id INTEGER PRIMARY KEY,
    order INTEGER PRIMARY KEY,
    abbreviation TEXT NOT NULL
)
