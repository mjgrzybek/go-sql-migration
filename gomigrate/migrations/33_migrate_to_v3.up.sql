ALTER TABLE books ADD COLUMN description TEXT DEFAULT '<unknown>';
UPDATE books SET description = 'authored by ' || author || ' in ' || IFNULL(year, 0);
ALTER TABLE books DROP COLUMN author;
ALTER TABLE books DROP COLUMN year;