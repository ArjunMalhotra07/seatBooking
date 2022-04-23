
CREATE TABLE bookSeat(
      id         INT AUTO_INCREMENT NOT NULL,
      seatID      VARCHAR(255) NOT NULL,
      bookedStatus  TINYINT NOT NULL,
      PRIMARY KEY (id)
    );
INSERT INTO bookSeat 
      (seatID, bookedStatus)
    VALUES
      ('A1', 0),
      ('A2', 0),
      ('A3', 0),
      ('A4', 0),
      ('A5', 0);