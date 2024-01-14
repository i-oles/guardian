-- Insert bulbs
INSERT INTO bulbs (ip_addr, bulb_id, type, luminance, red, green, blue) VALUES ('192.168.0.15', 'salon_main_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (ip_addr, bulb_id, type, luminance, red, green, blue) VALUES ('192.168.0.20', 'sofa_lamp', 'white', 30, NULL, NULL, NULL);
INSERT INTO bulbs (ip_addr, bulb_id, type, luminance, red, green, blue) VALUES ('192.168.0.23', 'table_lamp', 'white', 30, NULL, NULL, NULL);

-- Insert presets
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cinema', 'table_lamp', 5);
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cozy', 'table_lamp', 10);
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cozy', 'sofa_lamp', 10);
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cleaning', 'table_lamp', 100);
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cleaning', 'sofa_lamp', 100);
INSERT INTO presets (preset_name, bulb_id, luminance) VALUES ('cleaning', 'salon_main_lamp', 100);
