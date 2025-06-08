SELECT id FROM name_prefix WHERE id = ?;

SELECT id FROM name_prefix WHERE id = :id AND name IN ('a','b') AND id BETWEEN 'a' AND 'z' OR id > 1 OR id > '${id}' JOIN name_prefix__abbr ON name_prefix.id = name_prefix__abbr.name_prefix__id JOIN table USING (column) ORDER BY id ASC LIMIT 10 OFFSET 100;
