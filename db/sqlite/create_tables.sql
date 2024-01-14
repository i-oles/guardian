-- Create bulbs table
CREATE TABLE IF NOT EXISTS bulbs (
    bulb_id VARCHAR NOT NULL PRIMARY KEY,
    ip_addr VARCHAR(30) NOT NULL,
    type VARCHAR(30) NOT NULL,
    luminance INTEGER NOT NULL,
    red INTEGER,
    green INTEGER,
    blue INTEGER
);

-- Create presets table
CREATE TABLE IF NOT EXISTS presets (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    preset_name VARCHAR NOT NULL,
    bulb_id VARCHAR NOT NULL,
    luminance INTEGER NOT NULL,
    FOREIGN KEY (bulb_id) REFERENCES bulbs(bulb_id)
);
