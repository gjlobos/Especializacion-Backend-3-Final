CREATE DATABASE IF NOT EXISTS `finalEspBackend3DB` 
/*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `finalEspBackend3DB`;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dentists`
--

DROP TABLE IF EXISTS `dentists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `dentists` (
  `id` int DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `registration_number` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `patients`
--

DROP TABLE IF EXISTS `patients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `patients` (
  `id` int DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `address` varchar(50) DEFAULT NULL,
  `personal_id_number` int DEFAULT NULL,
  `creation_date` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `appointments`
--

DROP TABLE IF EXISTS `appointments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `appointments` (
  `id` int DEFAULT NULL,
  `dentist_registration_number` varchar(50) DEFAULT NULL,
  `patient_personal_id_number` int DEFAULT NULL,
  `appointment_date` varchar(50) DEFAULT NULL,
  `appointment_time` varchar(50) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dentists`
--

LOCK TABLES `dentists` WRITE;

/*!40000 ALTER TABLE `dentists` DISABLE KEYS */;

INSERT INTO `dentists` VALUES (1, 'Ezequiel', 'Bases', '421804'),(2, 'Natalia', 'Davolio', '420488'),(3, 'Natalia', 'Padovani', '580600');

/*!40000 ALTER TABLE `dentists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `patients`
--

LOCK TABLES `patients` WRITE;

/*!40000 ALTER TABLE `patients` DISABLE KEYS */;

INSERT INTO `patients` VALUES (1, 'Hernan', 'Colombo', 'Santa Maria de Oro 200',20345678,'19/01/2023'),(2, 'Veronica', 'Casella', 'San Martin 425',23480233,'20/02/2023'),(3, 'Pablo', 'Gonzalez', 'Lavalle 843',25765910,'10/03/2023');

/*!40000 ALTER TABLE `patients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `appointments`
--

LOCK TABLES `appointments` WRITE;

/*!40000 ALTER TABLE `appointments` DISABLE KEYS */;

INSERT INTO `appointments` VALUES (1, '421804', 20345678, '19/04/2023','18:00','Tratamiento de conducto'),(2, '580600', 23480233, '25/04/2023','11:00','Revision general'),(3, '420488', 25765910, '18/03/2023','15:00','Limpieza dientes superiores');

/*!40000 ALTER TABLE `appointments` ENABLE KEYS */;
UNLOCK TABLES;


/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;