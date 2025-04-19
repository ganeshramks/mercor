CREATE DATABASE IF NOT EXISTS mercor;

USE mercor;

DROP TABLE if exists job;
CREATE TABLE IF NOT EXISTS job (
    id INT NOT NULL AUTO_INCREMENT,
    version INT NOT NULL,
    jobId VARCHAR(30) NOT NULL,
    status ENUM('active', 'inactive', 'extended') NOT NULL,
    rate FLOAT DEFAULT 0,
    title VARCHAR(50),
    companyId INT NOT NULL,
    contractorId INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


drop table if exists timeLog;
CREATE TABLE IF NOT EXISTS timeLog (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    timeLogId INT NOT NULL,
    duration INT NOT NULL,
    timeStart TIMESTAMP NOT NULL,
    timeEnd TIMESTAMP NOT NULL,
    type ENUM('captured', 'adjusted') DEFAULT 'captured' NOT NULL,
    version INT NOT NULL,
    jobUid INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS paymentLineItems;

CREATE TABLE paymentLineItems (
    id INT PRIMARY KEY AUTO_INCREMENT,
    paymentLineItemId INT NOT NULL,
    jobUid INT NOT NULL,
    timeLogUid INT NOT NULL,
    amount DOUBLE NOT NULL,
    status ENUM('paid', 'not-paid') NOT NULL,
    version INT NOT NULL,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_payment_line_item (paymentLineItemId, jobUid, timeLogUid),
    INDEX idx_status (status)
);

