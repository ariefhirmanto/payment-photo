-- MySQL dump 10.16  Distrib 10.1.37-MariaDB, for Win32 (AMD64)
--
-- Host: localhost    Database: project
-- ------------------------------------------------------
-- Server version	5.7.37-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `amount` bigint(20) DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `payment_type` varchar(100) DEFAULT NULL,
  `qr_code_url` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`id`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,1000,'pending','0','','2022-04-14 22:41:49','2022-04-14 22:41:49'),(2,999,'pending','0','','2022-04-16 21:15:04','2022-04-16 21:15:04'),(3,995,'pending','0','','2022-04-16 21:17:07','2022-04-16 21:17:07'),(4,995,'pending','0','','2022-04-16 21:18:02','2022-04-16 21:18:02'),(5,998,'pending','0','','2022-04-16 21:18:50','2022-04-16 21:18:50'),(6,998,'pending','0','','2022-04-16 21:23:40','2022-04-16 21:23:40'),(7,998,'pending','0','https://api.sandbox.midtrans.com/v2/qris/894bfbbe-fca4-42ae-a560-0de251f679e6/qr-code','2022-04-17 12:03:20','2022-04-17 12:03:21'),(8,995,'pending','0','https://api.sandbox.midtrans.com/v2/qris/d9fbfa74-53fd-4257-943a-b6d7e7272736/qr-code','2022-04-17 12:29:12','2022-04-17 12:29:13'),(9,995,'pending','0','https://api.sandbox.midtrans.com/v2/qris/741a6c83-0869-466a-9aa3-be27768b3593/qr-code','2022-04-17 12:30:13','2022-04-17 12:30:14'),(10,995,'pending','0','','2022-04-17 12:31:26','2022-04-17 12:31:26'),(11,995,'pending','1','https://api.sandbox.midtrans.com/v2/gopay/154af8b6-c5cf-423e-a5d7-c6e158880d19/qr-code','2022-04-17 12:32:45','2022-04-17 12:32:46');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'project'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-17 14:55:12
