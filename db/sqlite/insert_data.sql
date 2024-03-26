-- Insert bulbs
INSERT INTO bulbs (bulb_id, bulb_name, type, luminance, red, green, blue)
VALUES ('0x000000001be46fb3', 'salon_main_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (bulb_id, bulb_name, type, luminance, red, green, blue)
VALUES ('0x000000001e369abf', 'bed_right_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (bulb_id, bulb_name, type, luminance, red, green, blue)
VALUES ('0x000000001b37dc76', 'bedroom_main_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (bulb_id, bulb_name, type, luminance, red, green, blue)
VALUES ('0x000000001be3a66e', 'table_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (bulb_id, bulb_name, type, luminance, red, green, blue)
VALUES ('0x000000001be3b3cf', 'sofa_lamp', 'rgb', 10, 255, 223, 142);

-- Insert presets
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cinema', 'table_lamp', 5);
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cozy', 'table_lamp', 10);
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cozy', 'sofa_lamp', 10);
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cleaning', 'table_lamp', 100);
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cleaning', 'sofa_lamp', 100);
INSERT INTO presets (preset_name, bulb_id, luminance)
VALUES ('cleaning', 'salon_main_lamp', 100);
