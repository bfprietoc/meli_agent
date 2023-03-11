
CREATE TABLE IF NOT EXISTS server_data (
	id serial PRIMARY KEY,	
	cpu_info VARCHAR (100),
	process_info JSON,
	users_info JSON,
	os_info VARCHAR(100)
);
