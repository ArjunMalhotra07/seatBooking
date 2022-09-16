
CREATE TABLE bookSeat(
      id         INT AUTO_INCREMENT NOT NULL,
      seatID      VARCHAR(255) NOT NULL,
      bookedStatus  TINYINT NOT NULL,
      PRIMARY KEY (id)
    );
INSERT INTO bookSeat 
      (seatID, bookedStatus)
    VALUES
      ('D1', 0),
      ('D2', 0),
      ('D3', 0),
      ('D4', 0),
      ('D5', 0);