package sqlite

const create string = `
  CREATE TABLE IF NOT EXISTS bulbs (
  id INTEGER NOT NULL PRIMARY KEY,
  ip VARCHAR(30) NOT NULL,
  name VARCHAR(30) NOT NULL,
  type VARCHAR(30) NOT NULL,
  luminance INTEGER NOT NULL,
  is_on BOOLEAN NOT NULL,
  preset VARCHAR(30) NOT NULL,
  );`

const create string = `
	  CREATE TABLE IF NOT EXISTS presets (
	  