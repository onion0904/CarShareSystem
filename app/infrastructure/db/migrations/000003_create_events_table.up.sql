CREATE TABLE events (
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	user_id VARCHAR(255) NOT NULL,
	together BOOLEAN NOT NULL,
	description VARCHAR(255) NOT NULL,
	year INT NOT NULL,	
	month INT NOT NULL,
	day INT NOT NULL,
    date DATETIME NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	start_date DATETIME NOT NULL,
	end_date DATETIME NOT NULL,
    important BOOLEAN NOT NULL,
	CONSTRAINT foreign_key_products_user_id FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;