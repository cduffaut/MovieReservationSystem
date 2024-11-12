CREATE DATABASE IF NOT EXISTS `MovieReservationSys`;

CREATE TABLE IF NOT EXISTS `MovieReservationSys`.`MovieSession` (
    `MovieName` VARCHAR(255) NOT NULL,
	`ClientName` VARCHAR(255) NOT NULL,
	`ClientFirstName` VARCHAR(255) NOT NULL,
	`ClientMail` VARCHAR(255) NOT NULL
);
