CREATE TABLE players (
	ID_players INT AUTO_INCREMENT PRIMARY KEY,
    full_Name VARCHAR(100) NOT NULL,
    birthdate DATE NOT NULL,
    dominant_foot ENUM('Derecho', 'Izquierdo', 'Ambos'),
    position VARCHAR(100) NOT NULL,
    `Number_T-shirt` INT NOT NULL,
    ID_Equipment INT,
    FOREIGN KEY (ID_Equipment) REFERENCES equipment(ID_Equipment) ON DELETE SET NULL
);