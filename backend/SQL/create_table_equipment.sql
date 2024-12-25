CREATE TABLE equipment (
    ID_Equipment INT AUTO_INCREMENT PRIMARY KEY,
    team_name VARCHAR(100) NOT NULL,     
    country VARCHAR(50) NOT NULL,             
    founding_date DATE,                
    stadium VARCHAR(100),                
    UNIQUE (team_name)                  
);