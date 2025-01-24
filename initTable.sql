-- --------------------------------------------------------
-- Host:                         fasms-db1.c50ay4oc0buf.us-east-1.rds.amazonaws.com
-- Server version:               10.11.9-MariaDB-log - managed by https://aws.amazon.com/rds/
-- Server OS:                    Linux
-- HeidiSQL Version:             12.8.0.6908
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for fasms
CREATE DATABASE IF NOT EXISTS `fasms` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci */;
USE `fasms`;

-- Dumping structure for table fasms.applicants
CREATE TABLE IF NOT EXISTS `applicants` (
  `id` varchar(36) NOT NULL DEFAULT '',
  `nric` char(9) DEFAULT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `employment_status` varchar(10) DEFAULT NULL,
  `sex` varchar(10) NOT NULL DEFAULT '',
  `date_of_birth` char(10) DEFAULT NULL,
  `household` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping data for table fasms.applicants: ~3 rows (approximately)
INSERT INTO `applicants` (`id`, `nric`, `name`, `employment_status`, `sex`, `date_of_birth`, `household`) VALUES
	('01913b7a-4493-74b2-93f8-e684c4ca935c', 'S9179333T', 'James', 'unemployed', 'male', '1990-07-01', NULL),
	('01913b80-2c04-7f9d-86a4-497ef68cb3a0', 'S9527231N', 'Mary', 'unemployed', 'female', '1984-10-06', '[\'01913b88-1d4d-7152-a7ce-75796a2e8ecf\',\'01913b88-65c6-7255-820f-9c4dd1e5ce79\']'),
	('6f17927c-8bd7-447b-8885-200801ffd716', 'S9827233T', 'Nicole', 'unemployed', 'female', '14-06-2011', '');

-- Dumping structure for table fasms.applications
CREATE TABLE IF NOT EXISTS `applications` (
  `id` char(36) NOT NULL,
  `date_of_application` char(10) NOT NULL,
  `scheme_id` char(36) NOT NULL,
  `applicant_id` char(36) NOT NULL,
  `status` varchar(50) NOT NULL DEFAULT '',
  `disbursed` varchar(50) DEFAULT NULL,
  `disbursed_date` char(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping data for table fasms.applications: ~0 rows (approximately)
INSERT INTO `applications` (`id`, `date_of_application`, `scheme_id`, `applicant_id`, `status`, `disbursed`, `disbursed_date`) VALUES
	('83213b80-2c04-7f9d-86a4-497ef68cb3a0', '02-01-2025', '01913b89-befc-7ae3-bb37-3079aa7f1be0', '01913b80-2c04-7f9d-86a4-497ef68cb3a0', 'approved', 'yes', '10-01-2025'),
	('95843b80-7b04-7f9d-86a4-497ef68cb3a0', '22-12-2024', '01913b89-9a43-7163-8757-01cc254783f3', 'c04b8ea1-511a-4d9b-ba18-4b6da80fda73', 'pending approval', 'no', ''),
	('01213b80-7b04-7f9d-86a4-497ef68cb3a0', '02-01-2025', '01913b89-9a43-7163-8757-01cc254783f3', '01913b80-2c04-7f9d-86a4-497ef68cb3a0', 'approved', 'yes', '10-01-2025');

-- Dumping structure for table fasms.benefits
CREATE TABLE IF NOT EXISTS `benefits` (
  `id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `amount` decimal(20,6) NOT NULL DEFAULT 0.000000,
  `scheme_id` char(36) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping data for table fasms.benefits: ~3 rows (approximately)
INSERT INTO `benefits` (`id`, `name`, `amount`, `scheme_id`) VALUES
	('1913b8b-9b12-7d2c-a1fa-ea613b802ebc', 'SkillsFuture Credits', 500.000000, '01913b89-9a43-7163-8757-01cc254783f3'),
	('73413b89-befc-7ae3-bb37-3079aa7f1be0', 'Monthly Allowance', 200.000000, '01913b89-befc-7ae3-bb37-3079aa7f1be0'),
	('3947f94a-1af0-45b4-bcb2-11611b159d91', 'One time cash payout', 300.000000, '5747f94a-1af0-45b4-bcb2-11611b159d91');

-- Dumping structure for table fasms.dependents
CREATE TABLE IF NOT EXISTS `dependents` (
  `id` char(36) NOT NULL,
  `name` char(100) NOT NULL,
  `employment_status` char(10) NOT NULL,
  `sex` char(10) NOT NULL,
  `relation` char(10) NOT NULL,
  `date_of_birth` char(10) NOT NULL,
  `parent_id` char(36) NOT NULL,
  PRIMARY KEY (`id` DESC),
  KEY `FK__applicants` (`parent_id`),
  CONSTRAINT `FK__applicants` FOREIGN KEY (`parent_id`) REFERENCES `applicants` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping data for table fasms.dependents: ~2 rows (approximately)
INSERT INTO `dependents` (`id`, `name`, `employment_status`, `sex`, `relation`, `date_of_birth`, `parent_id`) VALUES
	('01913b88-65c6-7255-820f-9c4dd1e5ce79', 'Jayden', 'unemployed', 'male', 'son', '03-12-2015', '01913b80-2c04-7f9d-86a4-497ef68cb3a0'),
	('01913b88-1d4d-7152-a7ce-75796a2e8ecf', 'Gwen', 'unemployed', 'female', 'daughter', '02-01-2016', '01913b80-2c04-7f9d-86a4-497ef68cb3a0');

-- Dumping structure for table fasms.schemes
CREATE TABLE IF NOT EXISTS `schemes` (
  `id` char(36) NOT NULL,
  `name` varchar(200) NOT NULL,
  `criteria` text DEFAULT NULL,
  `benefits` text DEFAULT NULL,
  PRIMARY KEY (`id` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping data for table fasms.schemes: ~3 rows (approximately)
INSERT INTO `schemes` (`id`, `name`, `criteria`, `benefits`) VALUES
	('5747f94a-1af0-45b4-bcb2-11611b159d91', 'Retrenchment Cash Assistance Scheme ', '{\r\n“age”: “>60”,“income”: “< 1500”\r\n}', NULL),
	('01913b89-befc-7ae3-bb37-3079aa7f1be0', 'Retrenchment Assistance Scheme (families)', '{\r\n“employment_status”: “unemployed”,\r\n“has_children”: {\r\n“school_level”: “== primary”\r\n}', NULL),
	('01913b89-9a43-7163-8757-01cc254783f3', 'Retrenchment Assistance Scheme', '{\r\n“employment_status”: “unemployed”\r\n}', '[\'01913b8b-9b12-7d2c-a1fa-ea613b802ebc\']');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
