-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: ebindalwasmin
-- ------------------------------------------------------
-- Server version	8.0.22-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `activity_log`
--

DROP TABLE IF EXISTS `activity_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `activity_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `log_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `subject_id` bigint unsigned DEFAULT NULL,
  `subject_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `causer_id` bigint unsigned DEFAULT NULL,
  `causer_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `properties` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `activity_log_log_name_index` (`log_name`),
  KEY `subject` (`subject_id`,`subject_type`),
  KEY `causer` (`causer_id`,`causer_type`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_log`
--

LOCK TABLES `activity_log` WRITE;
/*!40000 ALTER TABLE `activity_log` DISABLE KEYS */;
INSERT INTO `activity_log` VALUES (1,'default','created',1,'App\\Domains\\Announcement\\Models\\Announcement',NULL,NULL,'{\"attributes\": {\"area\": null, \"type\": \"info\", \"enabled\": true, \"ends_at\": null, \"message\": \"This is a <strong>Global</strong> announcement that will show on both the frontend and backend. <em>See <strong>AnnouncementSeeder</strong> for more usage examples.</em>\", \"starts_at\": null}}','2020-09-05 09:43:01','2020-09-05 09:43:01'),(2,'role','Super Admin created role entri data with permissions: None',2,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"entri data\", \"type\": \"user\"}, \"permissions\": \"None\"}','2020-09-05 14:16:42','2020-09-05 14:16:42'),(3,'user','Super Admin updated user Test User with roles: entri data and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"entri data\", \"permissions\": \"\"}','2020-09-05 14:16:54','2020-09-05 14:16:54'),(4,'user','Super Admin updated user Test User with roles: entri data and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"entri data\", \"permissions\": \"\"}','2020-09-05 14:22:49','2020-09-05 14:22:49'),(5,'role','Super Admin updated role entri data with permissions: Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"entri data\", \"type\": \"user\"}, \"permissions\": \"Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri\"}','2020-09-05 14:27:06','2020-09-05 14:27:06'),(6,'user','Super Admin updated user Test User with roles: entri data and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"entri data\", \"permissions\": \"\"}','2020-09-05 14:27:36','2020-09-05 14:27:36'),(7,'user','Super Admin updated user Test User with roles: entri data and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"entri data\", \"permissions\": \"\"}','2020-09-07 10:16:31','2020-09-07 10:16:31'),(8,'user','Super Admin updated user Test User with roles: None and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"\"}','2020-09-07 11:00:50','2020-09-07 11:00:50'),(9,'user','Super Admin updated user Test User with roles: Kantor Pangkal Pinang and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-07 11:01:11','2020-09-07 11:01:11'),(10,'role','Super Admin updated role Kantor Babel with permissions: Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Babel\", \"type\": \"admin\"}, \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 11:25:14','2020-09-08 11:25:14'),(11,'role','Super Admin deleted role User Entri Bangka Belitung',6,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','[]','2020-09-08 11:27:00','2020-09-08 11:27:00'),(12,'role','Super Admin deleted role User Entri Pangkal Pinang',5,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','[]','2020-09-08 11:27:06','2020-09-08 11:27:06'),(13,'role','Super Admin deleted role entri data',2,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','[]','2020-09-08 11:27:14','2020-09-08 11:27:14'),(14,'role','Super Admin updated role Kantor Pangkal Pinang with permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Pangkal Pinang\", \"type\": \"admin\"}, \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 11:27:40','2020-09-08 11:27:40'),(15,'user','Super Admin updated user Test User with roles: Kantor Pangkal Pinang and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 11:28:50','2020-09-08 11:28:50'),(16,'user','Super Admin created user tes with roles: Kantor Babel and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\", \"active\": true, \"email_verified_at\": \"2020-09-08T11:52:59.000000Z\"}, \"roles\": \"Kantor Babel\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 11:52:59','2020-09-08 11:52:59'),(17,'role','Super Admin updated role Kantor Babel with permissions: All User Permissions, Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Babel\", \"type\": \"admin\"}, \"permissions\": \"All User Permissions, Petugas Dapat mengentri\"}','2020-09-08 11:55:47','2020-09-08 11:55:47'),(18,'role','Super Admin updated role Kantor Babel with permissions: None',4,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Babel\", \"type\": \"admin\"}, \"permissions\": \"None\"}','2020-09-08 11:59:57','2020-09-08 11:59:57'),(19,'role','Super Admin updated role Kantor Pangkal Pinang with permissions: None',3,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Pangkal Pinang\", \"type\": \"admin\"}, \"permissions\": \"None\"}','2020-09-08 12:00:06','2020-09-08 12:00:06'),(20,'role','Super Admin updated role Kantor Babel with permissions: Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Babel\", \"type\": \"user\"}, \"permissions\": \"Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri\"}','2020-09-08 12:00:24','2020-09-08 12:00:24'),(21,'role','Super Admin updated role Kantor Pangkal Pinang with permissions: Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','{\"role\": {\"name\": \"Kantor Pangkal Pinang\", \"type\": \"user\"}, \"permissions\": \"Semua Akses Petugas, Melihat list Data, Petugas Dapat mengentri\"}','2020-09-08 12:00:35','2020-09-08 12:00:35'),(22,'user','Super Admin updated user tes with roles: Kantor Babel and permissions: ',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"user\", \"email\": \"tes@tes.com\"}, \"roles\": \"Kantor Babel\", \"permissions\": \"\"}','2020-09-08 12:01:13','2020-09-08 12:01:13'),(23,'user','Super Admin updated user Test User with roles: Kantor Pangkal Pinang and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"\"}','2020-09-08 12:01:49','2020-09-08 12:01:49'),(24,'user','Super Admin updated user Test User with roles: Kantor Pangkal Pinang and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"\"}','2020-09-08 12:02:10','2020-09-08 12:02:10'),(25,'user','Super Admin updated user Test User with roles: Kantor Pangkal Pinang and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"\"}','2020-09-08 12:02:26','2020-09-08 12:02:26'),(26,'user','Super Admin updated user tes with roles: None and permissions: ',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"user\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"\"}','2020-09-08 12:04:16','2020-09-08 12:04:16'),(27,'user','Super Admin updated user Test User with roles: None and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"user\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"\"}','2020-09-08 12:05:10','2020-09-08 12:05:10'),(28,'user','Super Admin updated user tes with roles: None and permissions: ',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"\"}','2020-09-08 12:05:30','2020-09-08 12:05:30'),(29,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 12:08:39','2020-09-08 12:08:39'),(30,'user','Super Admin updated user tes with roles: Kantor Pangkal Pinang and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"Kantor Pangkal Pinang\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 12:12:15','2020-09-08 12:12:15'),(31,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 12:42:45','2020-09-08 12:42:45'),(32,'user','Super Admin updated user Test User with roles: None and permissions: ',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"\"}','2020-09-08 12:42:53','2020-09-08 12:42:53'),(33,'role','Super Admin deleted role Kantor Pangkal Pinang',3,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','[]','2020-09-08 12:43:05','2020-09-08 12:43:05'),(34,'role','Super Admin deleted role Kantor Babel',4,'App\\Domains\\Auth\\Models\\Role',1,'App\\Domains\\Auth\\Models\\User','[]','2020-09-08 12:43:09','2020-09-08 12:43:09'),(35,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 12:43:39','2020-09-08 12:43:39'),(36,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-08 12:43:48','2020-09-08 12:43:48'),(37,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:10:50','2020-09-10 00:10:50'),(38,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:11:08','2020-09-10 00:11:08'),(39,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:16:00','2020-09-10 00:16:00'),(40,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:17:49','2020-09-10 00:17:49'),(41,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:20:25','2020-09-10 00:20:25'),(42,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:24:24','2020-09-10 00:24:24'),(43,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": 1}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 00:34:47','2020-09-10 00:34:47'),(44,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:37:38','2020-09-10 00:37:38'),(45,'user','Super Admin updated user Test User with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test User\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:38:56','2020-09-10 00:38:56'),(46,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:44:28','2020-09-10 00:44:28'),(47,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:47:24','2020-09-10 00:47:24'),(48,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:50:15','2020-09-10 00:50:15'),(49,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:54:29','2020-09-10 00:54:29'),(50,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 00:54:58','2020-09-10 00:54:58'),(51,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:06:33','2020-09-10 01:06:33'),(52,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:06:49','2020-09-10 01:06:49'),(53,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:09:03','2020-09-10 01:09:03'),(54,'user','Super Admin updated user Test xx with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xx\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:16:23','2020-09-10 01:16:23'),(55,'user','Super Admin updated user Test xxi with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxi\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:18:31','2020-09-10 01:18:31'),(56,'user','Super Admin updated user Test xxi with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxi\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:19:30','2020-09-10 01:19:30'),(57,'user','Super Admin updated user Test xxid with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxid\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:20:24','2020-09-10 01:20:24'),(58,'user','Super Admin updated user Test xxid with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxid\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:21:57','2020-09-10 01:21:57'),(59,'user','Super Admin updated user Test xxid with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxid\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:22:27','2020-09-10 01:22:27'),(60,'user','Super Admin updated user Test xxid with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxid\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:23:21','2020-09-10 01:23:21'),(61,'user','Super Admin updated user Test xxidg with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxidg\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:25:38','2020-09-10 01:25:38'),(62,'user','Super Admin updated user Test xxidg with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxidg\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:28:06','2020-09-10 01:28:06'),(63,'user','Super Admin updated user Test xxidg with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxidg\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:30:43','2020-09-10 01:30:43'),(64,'user','Super Admin updated user Test xxidg with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxidg\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:30:58','2020-09-10 01:30:58'),(65,'user','Super Admin updated user Test xxidg with roles: None and permissions: Petugas Dapat mengentri',2,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Test xxidg\", \"type\": \"admin\", \"email\": \"user@user.com\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:31:24','2020-09-10 01:31:24'),(66,'user','Super Admin created user tes with roles: None and permissions: Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.123\", \"active\": true, \"id_kantor\": null, \"email_verified_at\": \"2020-09-10T08:32:43.000000Z\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:32:43','2020-09-10 01:32:43'),(67,'user','Super Admin updated user tescccc with roles: None and permissions: Petugas Dapat mengentri',3,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tescccc\", \"type\": \"admin\", \"email\": \"tes@tes.com\", \"id_kantor\": 2}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:34:32','2020-09-10 01:34:32'),(68,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.123\", \"id_kantor\": 1}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:35:44','2020-09-10 01:35:44'),(69,'user','Super Admin updated user tes with roles: None and permissions: Petugas Dapat mengentri',4,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"tes\", \"type\": \"admin\", \"email\": \"tes@tes.123\", \"id_kantor\": \"2\"}, \"roles\": \"None\", \"permissions\": \"Petugas Dapat mengentri\"}','2020-09-10 01:46:58','2020-09-10 01:46:58'),(70,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": \"99\"}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 07:37:00','2020-09-10 07:37:00'),(71,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": \"1\"}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 17:24:18','2020-09-10 17:24:18'),(72,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": \"99\"}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 18:05:17','2020-09-10 18:05:17'),(73,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": \"2\"}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 18:05:46','2020-09-10 18:05:46'),(74,'user','Super Admin updated user Super Admin with roles: Administrator and permissions: ',1,'App\\Domains\\Auth\\Models\\User',1,'App\\Domains\\Auth\\Models\\User','{\"user\": {\"name\": \"Super Admin\", \"type\": \"admin\", \"email\": \"admin@admin.com\", \"id_kantor\": \"99\"}, \"roles\": \"Administrator\", \"permissions\": \"\"}','2020-09-10 18:19:25','2020-09-10 18:19:25');
/*!40000 ALTER TABLE `activity_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `announcements`
--

DROP TABLE IF EXISTS `announcements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `announcements` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `area` enum('frontend','backend') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` enum('info','danger','warning','success') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'info',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT '1',
  `starts_at` timestamp NULL DEFAULT NULL,
  `ends_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `announcements`
--

LOCK TABLES `announcements` WRITE;
/*!40000 ALTER TABLE `announcements` DISABLE KEYS */;
INSERT INTO `announcements` VALUES (1,NULL,'info','This is a <strong>Global</strong> announcement that will show on both the frontend and backend. <em>See <strong>AnnouncementSeeder</strong> for more usage examples.</em>',0,NULL,NULL,'2020-09-05 09:43:01','2020-09-05 09:43:01');
/*!40000 ALTER TABLE `announcements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_izinkeimigrasian`
--

DROP TABLE IF EXISTS `data_izinkeimigrasian`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_izinkeimigrasian` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_jenis` int NOT NULL,
  `id_user` int DEFAULT NULL,
  `id_kantor` int DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `laki` int unsigned DEFAULT '0',
  `perempuan` int unsigned DEFAULT '0',
  `total` bigint unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=154 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_izinkeimigrasian`
--

LOCK TABLES `data_izinkeimigrasian` WRITE;
/*!40000 ALTER TABLE `data_izinkeimigrasian` DISABLE KEYS */;
INSERT INTO `data_izinkeimigrasian` VALUES (1,13,3,2,'2020-09-08',111,111,111000000),(2,14,3,2,'2020-09-08',111,111,111000000),(3,15,3,2,'2020-09-08',111,111,155400000),(4,17,3,2,'2020-09-08',NULL,NULL,NULL),(5,18,3,2,'2020-09-08',111,324,435000000),(6,19,3,2,'2020-09-08',NULL,NULL,NULL),(7,20,3,2,'2020-09-08',NULL,NULL,NULL),(8,21,3,2,'2020-09-08',NULL,NULL,NULL),(9,22,3,2,'2020-09-08',234,234,140400000),(10,51,3,2,'2020-09-08',NULL,NULL,NULL),(11,24,3,2,'2020-09-08',34,34,340000000),(12,25,3,2,'2020-09-08',34,34,340000000),(13,26,3,2,'2020-09-08',NULL,NULL,NULL),(14,28,3,2,'2020-09-08',NULL,NULL,NULL),(15,29,3,2,'2020-09-08',34,34,68000000),(16,30,3,2,'2020-09-08',NULL,NULL,NULL),(17,31,3,2,'2020-09-08',NULL,NULL,NULL),(18,13,1,1,'2020-09-09',NULL,NULL,NULL),(19,14,1,1,'2020-09-09',NULL,NULL,NULL),(20,15,1,1,'2020-09-09',NULL,NULL,NULL),(21,17,1,1,'2020-09-09',NULL,NULL,NULL),(22,18,1,1,'2020-09-09',NULL,NULL,NULL),(23,19,1,1,'2020-09-09',NULL,NULL,NULL),(24,20,1,1,'2020-09-09',NULL,NULL,NULL),(25,21,1,1,'2020-09-09',NULL,NULL,NULL),(26,22,1,1,'2020-09-09',NULL,NULL,NULL),(27,51,1,1,'2020-09-09',NULL,NULL,NULL),(28,24,1,1,'2020-09-09',NULL,NULL,NULL),(29,25,1,1,'2020-09-09',NULL,NULL,NULL),(30,26,1,1,'2020-09-09',NULL,NULL,NULL),(31,28,1,1,'2020-09-09',NULL,NULL,NULL),(32,29,1,1,'2020-09-09',NULL,NULL,NULL),(33,30,1,1,'2020-09-09',NULL,NULL,NULL),(34,31,1,1,'2020-09-09',NULL,NULL,NULL),(35,13,1,1,'2020-09-10',NULL,NULL,NULL),(36,14,1,1,'2020-09-10',NULL,NULL,NULL),(37,15,1,1,'2020-09-10',NULL,NULL,NULL),(38,17,1,1,'2020-09-10',NULL,NULL,NULL),(39,18,1,1,'2020-09-10',NULL,NULL,NULL),(40,19,1,1,'2020-09-10',NULL,NULL,NULL),(41,20,1,1,'2020-09-10',NULL,NULL,NULL),(42,21,1,1,'2020-09-10',NULL,NULL,NULL),(43,22,1,1,'2020-09-10',NULL,NULL,NULL),(44,51,1,1,'2020-09-10',NULL,NULL,NULL),(45,24,1,1,'2020-09-10',NULL,NULL,NULL),(46,25,1,1,'2020-09-10',NULL,NULL,NULL),(47,26,1,1,'2020-09-10',NULL,NULL,NULL),(48,28,1,1,'2020-09-10',NULL,NULL,NULL),(49,29,1,1,'2020-09-10',NULL,NULL,NULL),(50,30,1,1,'2020-09-10',NULL,NULL,NULL),(51,31,1,1,'2020-09-10',NULL,NULL,NULL),(52,13,1,1,'2020-09-10',NULL,NULL,NULL),(53,14,1,1,'2020-09-10',NULL,NULL,NULL),(54,15,1,1,'2020-09-10',NULL,NULL,NULL),(55,17,1,1,'2020-09-10',NULL,NULL,NULL),(56,18,1,1,'2020-09-10',NULL,NULL,NULL),(57,19,1,1,'2020-09-10',NULL,NULL,NULL),(58,20,1,1,'2020-09-10',NULL,NULL,NULL),(59,21,1,1,'2020-09-10',NULL,NULL,NULL),(60,22,1,1,'2020-09-10',NULL,NULL,NULL),(61,51,1,1,'2020-09-10',NULL,NULL,NULL),(62,24,1,1,'2020-09-10',NULL,NULL,NULL),(63,25,1,1,'2020-09-10',NULL,NULL,NULL),(64,26,1,1,'2020-09-10',NULL,NULL,NULL),(65,28,1,1,'2020-09-10',NULL,NULL,NULL),(66,29,1,1,'2020-09-10',NULL,NULL,NULL),(67,30,1,1,'2020-09-10',NULL,NULL,NULL),(68,31,1,1,'2020-09-10',NULL,NULL,NULL),(69,13,4,2,'2020-09-10',7,77,42000000),(70,14,4,2,'2020-09-10',7,76,41500000),(71,15,4,2,'2020-09-10',7,77,58800000),(72,17,4,2,'2020-09-10',55,NULL,41250000),(73,18,4,2,'2020-09-10',5,NULL,5000000),(74,19,4,2,'2020-09-10',6,NULL,9000000),(75,20,4,2,'2020-09-10',7,NULL,NULL),(76,21,4,2,'2020-09-10',66,NULL,66000000),(77,22,4,2,'2020-09-10',4,NULL,1200000),(78,51,4,2,'2020-09-10',44,NULL,220000000),(79,24,4,2,'2020-09-10',NULL,NULL,NULL),(80,25,4,2,'2020-09-10',NULL,NULL,NULL),(81,26,4,2,'2020-09-10',NULL,NULL,NULL),(82,28,4,2,'2020-09-10',NULL,NULL,NULL),(83,29,4,2,'2020-09-10',NULL,NULL,NULL),(84,30,4,2,'2020-09-10',NULL,NULL,NULL),(85,31,4,2,'2020-09-10',NULL,NULL,NULL),(86,13,1,99,'2020-09-10',NULL,NULL,NULL),(87,14,1,99,'2020-09-10',NULL,NULL,NULL),(88,15,1,99,'2020-09-10',NULL,NULL,NULL),(89,17,1,99,'2020-09-10',NULL,NULL,NULL),(90,18,1,99,'2020-09-10',NULL,NULL,NULL),(91,19,1,99,'2020-09-10',NULL,NULL,NULL),(92,20,1,99,'2020-09-10',NULL,NULL,NULL),(93,21,1,99,'2020-09-10',NULL,NULL,NULL),(94,22,1,99,'2020-09-10',NULL,NULL,NULL),(95,51,1,99,'2020-09-10',NULL,NULL,NULL),(96,24,1,99,'2020-09-10',NULL,NULL,NULL),(97,25,1,99,'2020-09-10',NULL,NULL,NULL),(98,26,1,99,'2020-09-10',NULL,NULL,NULL),(99,28,1,99,'2020-09-10',NULL,NULL,NULL),(100,29,1,99,'2020-09-10',NULL,NULL,NULL),(101,30,1,99,'2020-09-10',NULL,NULL,NULL),(102,31,1,99,'2020-09-10',NULL,NULL,NULL),(103,13,1,1,'2020-09-11',6,6,6000000),(104,14,1,1,'2020-09-11',66,6,36000000),(105,15,1,1,'2020-09-11',4,4,5600000),(106,17,1,1,'2020-09-11',44,NULL,37500000),(107,18,1,1,'2020-09-11',6,66,72000000),(108,19,1,1,'2020-09-11',9,9,27000000),(109,20,1,1,'2020-09-11',9,88,194000000),(110,21,1,1,'2020-09-11',9,9,18000000),(111,22,1,1,'2020-09-11',8,NULL,2400000),(112,51,1,1,'2020-09-11',8,NULL,40000000),(113,24,1,1,'2020-09-11',9,NULL,45000000),(114,25,1,1,'2020-09-11',99,NULL,495000000),(115,26,1,1,'2020-09-11',NULL,99,1009800000),(116,28,1,1,'2020-09-11',8,19,16200000),(117,29,1,1,'2020-09-11',22,2,24000000),(118,30,1,1,'2020-09-11',2,55,99750000),(119,31,1,1,'2020-09-11',5,9,45500000),(120,13,1,2,'2020-09-11',NULL,1,500000),(121,14,1,2,'2020-09-11',NULL,23,11500000),(122,15,1,2,'2020-09-11',NULL,2,1400000),(123,17,1,2,'2020-09-11',NULL,NULL,2250000),(124,18,1,2,'2020-09-11',NULL,3,3000000),(125,19,1,2,'2020-09-11',11,NULL,16500000),(126,20,1,2,'2020-09-11',34,NULL,68000000),(127,21,1,2,'2020-09-11',4,NULL,4000000),(128,22,1,2,'2020-09-11',34,NULL,10200000),(129,51,1,2,'2020-09-11',43,NULL,215000000),(130,24,1,2,'2020-09-11',3,NULL,15000000),(131,25,1,2,'2020-09-11',NULL,3,15000000),(132,26,1,2,'2020-09-11',33,NULL,336600000),(133,28,1,2,'2020-09-11',33,NULL,19800000),(134,29,1,2,'2020-09-11',3,NULL,3000000),(135,30,1,2,'2020-09-11',3,NULL,5250000),(136,31,1,2,'2020-09-11',34,NULL,110500000),(137,13,1,99,'2020-09-12',2,2,2000000),(138,14,1,99,'2020-09-12',3,4,3500000),(139,15,1,99,'2020-09-12',2,4,4200000),(140,17,1,99,'2020-09-12',4,NULL,6000000),(141,18,1,99,'2020-09-12',5,5,10000000),(142,19,1,99,'2020-09-12',2,5,10500000),(143,20,1,99,'2020-09-12',2,5,14000000),(144,21,1,99,'2020-09-12',2,5,7000000),(145,22,1,99,'2020-09-12',25,5,9000000),(146,51,1,99,'2020-09-12',2,NULL,10000000),(147,24,1,99,'2020-09-12',22,NULL,10000000),(148,25,1,99,'2020-09-12',2,NULL,10000000),(149,26,1,99,'2020-09-12',2,NULL,20400000),(150,28,1,99,'2020-09-12',3,NULL,1800000),(151,29,1,99,'2020-09-12',3,NULL,3000000),(152,30,1,99,'2020-09-12',2,NULL,3500000),(153,31,1,99,'2020-09-12',3,NULL,9750000);
/*!40000 ALTER TABLE `data_izinkeimigrasian` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_paspor`
--

DROP TABLE IF EXISTS `data_paspor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_paspor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_jenis` int NOT NULL,
  `id_user` int DEFAULT NULL,
  `id_kantor` int DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `laki` int DEFAULT '0',
  `perempuan` int DEFAULT '0',
  `total` bigint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_paspor`
--

LOCK TABLES `data_paspor` WRITE;
/*!40000 ALTER TABLE `data_paspor` DISABLE KEYS */;
INSERT INTO `data_paspor` VALUES (1,5,3,2,'2019-02-01',111,111,77700000),(2,6,3,2,'2020-08-08',111,111,144300000),(3,7,3,2,'2020-07-09',111,111,33300000),(4,8,3,2,'2020-09-08',111,111,222000000),(5,5,1,1,'2020-09-09',21,3,8400000),(6,6,1,1,'2020-09-09',32,33,42250000),(7,7,1,1,'2020-09-09',2,33,5250000),(8,8,1,1,'2020-09-09',33,9,42000000),(9,5,1,1,'2020-09-10',2,5,2450000),(10,6,1,1,'2020-09-10',7,77,54600000),(11,7,1,1,'2020-09-10',7,7,2100000),(12,8,1,1,'2020-09-10',8,7,15000000),(13,5,1,1,'2020-09-10',NULL,NULL,NULL),(14,6,1,1,'2020-09-10',NULL,NULL,NULL),(15,7,1,1,'2020-09-10',NULL,NULL,NULL),(16,8,1,1,'2020-09-10',NULL,NULL,NULL),(17,5,4,2,'2020-09-10',22,3,8750000),(18,6,4,2,'2020-09-10',1,3,2600000),(19,7,4,2,'2020-09-10',3,3,900000),(20,8,4,2,'2020-09-10',3,3,6000000),(21,5,1,2,'2020-09-10',NULL,NULL,NULL),(22,6,1,2,'2020-09-10',NULL,NULL,NULL),(23,7,1,2,'2020-09-10',NULL,NULL,NULL),(24,8,1,2,'2020-09-10',NULL,NULL,NULL),(25,5,1,1,'2020-09-11',3,5,2800000),(26,6,1,1,'2020-09-11',6,3,5850000),(27,7,1,1,'2020-09-11',3,2,750000),(28,8,1,1,'2020-09-11',2,2,4000000),(29,5,1,2,'2020-09-11',2,NULL,700000),(30,6,1,2,'2020-09-11',NULL,3,1950000),(31,7,1,2,'2020-09-11',NULL,3,450000),(32,8,1,2,'2020-09-11',1,NULL,1000000),(33,5,1,99,'2020-09-12',3,4,2450000),(34,6,1,99,'2020-09-12',2,1,1950000),(35,7,1,99,'2020-09-12',11,3,2100000),(36,8,1,99,'2020-09-12',3,3,6000000);
/*!40000 ALTER TABLE `data_paspor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_pnbplainnya`
--

DROP TABLE IF EXISTS `data_pnbplainnya`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_pnbplainnya` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_jenis` int NOT NULL,
  `id_user` int DEFAULT NULL,
  `id_kantor` int DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `laki` int DEFAULT '0',
  `perempuan` int DEFAULT '0',
  `total` bigint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=109 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_pnbplainnya`
--

LOCK TABLES `data_pnbplainnya` WRITE;
/*!40000 ALTER TABLE `data_pnbplainnya` DISABLE KEYS */;
INSERT INTO `data_pnbplainnya` VALUES (1,33,3,2,'2020-08-31',34,34,68000000),(2,34,3,2,'2020-09-08',NULL,NULL,NULL),(3,35,3,2,'2020-09-08',NULL,NULL,NULL),(4,36,3,2,'2020-09-08',34,34,68000000),(5,37,3,2,'2020-09-08',NULL,NULL,NULL),(6,38,3,2,'2020-09-08',34,34,68000000),(7,39,3,2,'2020-09-08',34,34,68000000),(8,40,3,2,'2020-09-08',4,34,57000000),(9,42,3,2,'2020-09-08',4,3,17500000),(10,43,3,2,'2020-09-08',4,3,17500000),(11,44,3,2,'2020-09-08',4,234,95200000),(12,45,3,2,'2020-09-08',34,34,204000000),(13,33,1,1,'2020-09-09',NULL,NULL,NULL),(14,34,1,1,'2020-09-09',NULL,NULL,NULL),(15,35,1,1,'2020-09-09',NULL,NULL,NULL),(16,36,1,1,'2020-09-09',NULL,NULL,NULL),(17,37,1,1,'2020-09-09',NULL,NULL,NULL),(18,38,1,1,'2020-09-09',NULL,NULL,NULL),(19,39,1,1,'2020-09-09',NULL,NULL,NULL),(20,40,1,1,'2020-09-09',NULL,NULL,NULL),(21,42,1,1,'2020-09-09',NULL,NULL,NULL),(22,43,1,1,'2020-09-09',NULL,NULL,NULL),(23,44,1,1,'2020-09-09',NULL,NULL,NULL),(24,45,1,1,'2020-09-09',NULL,NULL,NULL),(25,33,1,1,'2020-09-10',NULL,NULL,NULL),(26,34,1,1,'2020-09-10',NULL,NULL,NULL),(27,35,1,1,'2020-09-10',NULL,NULL,NULL),(28,36,1,1,'2020-09-10',NULL,NULL,NULL),(29,37,1,1,'2020-09-10',NULL,NULL,NULL),(30,38,1,1,'2020-09-10',NULL,NULL,NULL),(31,39,1,1,'2020-09-10',NULL,NULL,NULL),(32,40,1,1,'2020-09-10',NULL,NULL,NULL),(33,42,1,1,'2020-09-10',NULL,NULL,NULL),(34,43,1,1,'2020-09-10',NULL,NULL,NULL),(35,44,1,1,'2020-09-10',NULL,NULL,NULL),(36,45,1,1,'2020-09-10',NULL,NULL,NULL),(37,33,1,1,'2020-09-10',NULL,NULL,NULL),(38,34,1,1,'2020-09-10',NULL,NULL,NULL),(39,35,1,1,'2020-09-10',NULL,NULL,NULL),(40,36,1,1,'2020-09-10',NULL,NULL,NULL),(41,37,1,1,'2020-09-10',NULL,NULL,NULL),(42,38,1,1,'2020-09-10',NULL,NULL,NULL),(43,39,1,1,'2020-09-10',NULL,NULL,NULL),(44,40,1,1,'2020-09-10',NULL,NULL,NULL),(45,42,1,1,'2020-09-10',NULL,NULL,NULL),(46,43,1,1,'2020-09-10',NULL,NULL,NULL),(47,44,1,1,'2020-09-10',NULL,NULL,NULL),(48,45,1,1,'2020-09-10',NULL,NULL,NULL),(49,33,4,2,'2020-09-10',NULL,NULL,NULL),(50,34,4,2,'2020-09-10',NULL,NULL,NULL),(51,35,4,2,'2020-09-10',NULL,NULL,NULL),(52,36,4,2,'2020-09-10',NULL,NULL,NULL),(53,37,4,2,'2020-09-10',NULL,NULL,NULL),(54,38,4,2,'2020-09-10',NULL,NULL,NULL),(55,39,4,2,'2020-09-10',NULL,NULL,NULL),(56,40,4,2,'2020-09-10',NULL,NULL,NULL),(57,42,4,2,'2020-09-10',NULL,NULL,NULL),(58,43,4,2,'2020-09-10',NULL,NULL,NULL),(59,44,4,2,'2020-09-10',NULL,NULL,NULL),(60,45,4,2,'2020-09-10',NULL,NULL,NULL),(61,33,1,99,'2020-09-10',NULL,NULL,NULL),(62,34,1,99,'2020-09-10',NULL,NULL,NULL),(63,35,1,99,'2020-09-10',NULL,NULL,NULL),(64,36,1,99,'2020-09-10',NULL,NULL,NULL),(65,37,1,99,'2020-09-10',NULL,NULL,NULL),(66,38,1,99,'2020-09-10',NULL,NULL,NULL),(67,39,1,99,'2020-09-10',NULL,NULL,NULL),(68,40,1,99,'2020-09-10',NULL,NULL,NULL),(69,42,1,99,'2020-09-10',NULL,NULL,NULL),(70,43,1,99,'2020-09-10',NULL,NULL,NULL),(71,44,1,99,'2020-09-10',NULL,NULL,NULL),(72,45,1,99,'2020-09-10',NULL,NULL,NULL),(73,33,1,1,'2020-09-11',99,99,198000000),(74,34,1,1,'2020-09-11',NULL,7,350000000),(75,35,1,1,'2020-09-11',55,5,3000000000),(76,36,1,1,'2020-09-11',5,5,10000000),(77,37,1,1,'2020-09-11',5,5,5000000),(78,38,1,1,'2020-09-11',5,5,10000000),(79,39,1,1,'2020-09-11',NULL,NULL,NULL),(80,40,1,1,'2020-09-11',NULL,NULL,NULL),(81,42,1,1,'2020-09-11',NULL,NULL,NULL),(82,43,1,1,'2020-09-11',NULL,NULL,NULL),(83,44,1,1,'2020-09-11',5,5,4000000),(84,45,1,1,'2020-09-11',5,5,30000000),(85,33,1,2,'2020-09-11',4,NULL,4000000),(86,34,1,2,'2020-09-11',4,NULL,200000000),(87,35,1,2,'2020-09-11',4,NULL,200000000),(88,36,1,2,'2020-09-11',NULL,3,3000000),(89,37,1,2,'2020-09-11',NULL,34,17000000),(90,38,1,2,'2020-09-11',NULL,4,4000000),(91,39,1,2,'2020-09-11',NULL,NULL,NULL),(92,40,1,2,'2020-09-11',NULL,NULL,NULL),(93,42,1,2,'2020-09-11',NULL,NULL,NULL),(94,43,1,2,'2020-09-11',NULL,NULL,NULL),(95,44,1,2,'2020-09-11',NULL,4,1600000),(96,45,1,2,'2020-09-11',NULL,44,132000000),(97,33,1,99,'2020-09-12',NULL,3,3000000),(98,34,1,99,'2020-09-12',NULL,33,1650000000),(99,35,1,99,'2020-09-12',NULL,3,150000000),(100,36,1,99,'2020-09-12',NULL,3,3000000),(101,37,1,99,'2020-09-12',NULL,3,1500000),(102,38,1,99,'2020-09-12',NULL,3,3000000),(103,39,1,99,'2020-09-12',NULL,NULL,NULL),(104,40,1,99,'2020-09-12',NULL,NULL,NULL),(105,42,1,99,'2020-09-12',NULL,NULL,NULL),(106,43,1,99,'2020-09-12',NULL,NULL,NULL),(107,44,1,99,'2020-09-12',NULL,3,1200000),(108,45,1,99,'2020-09-12',NULL,3,9000000);
/*!40000 ALTER TABLE `data_pnbplainnya` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_visa`
--

DROP TABLE IF EXISTS `data_visa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_visa` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_jenis` int NOT NULL,
  `id_user` int DEFAULT NULL,
  `id_kantor` int DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `laki` int DEFAULT '0',
  `perempuan` int DEFAULT '0',
  `total` bigint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_visa`
--

LOCK TABLES `data_visa` WRITE;
/*!40000 ALTER TABLE `data_visa` DISABLE KEYS */;
INSERT INTO `data_visa` VALUES (1,47,3,2,'2020-07-07',111,111,33300),(2,48,3,2,'2020-07-31',111,111,155400000),(3,49,3,2,'2020-09-08',111,111,44400000),(4,9,3,2,'2020-09-08',111,111,11100),(5,10,3,2,'2020-09-08',111,111,24420),(6,11,3,2,'2020-09-08',111,111,111000000),(7,47,1,1,'2020-09-09',9,9,2700),(8,48,1,1,'2020-09-09',9,9,12600000),(9,49,1,1,'2020-09-09',9,9,3600000),(10,9,1,1,'2020-09-09',NULL,NULL,0),(11,10,1,1,'2020-09-09',NULL,NULL,NULL),(12,11,1,1,'2020-09-09',NULL,NULL,NULL),(13,47,1,1,'2020-09-10',8,5,1950),(14,48,1,1,'2020-09-10',34,4,26600000),(15,49,1,1,'2020-09-10',4,44,9600000),(16,9,1,1,'2020-09-10',4,44,2400),(17,10,1,1,'2020-09-10',4,44,5280),(18,11,1,1,'2020-09-10',NULL,NULL,NULL),(19,47,1,1,'2020-09-10',NULL,NULL,NULL),(20,48,1,1,'2020-09-10',NULL,NULL,NULL),(21,49,1,1,'2020-09-10',NULL,NULL,NULL),(22,9,1,1,'2020-09-10',NULL,NULL,NULL),(23,10,1,1,'2020-09-10',NULL,NULL,NULL),(24,11,1,1,'2020-09-10',NULL,NULL,NULL),(25,47,4,2,'2020-09-10',3,3,900),(26,48,4,2,'2020-09-10',3,3,4200000),(27,49,4,2,'2020-09-10',3,33,7200000),(28,9,4,2,'2020-09-10',33,3,1800),(29,10,4,2,'2020-09-10',3,66,7590),(30,11,4,2,'2020-09-10',7,8,7500000),(31,47,1,99,'2020-09-10',NULL,NULL,NULL),(32,48,1,99,'2020-09-10',NULL,NULL,NULL),(33,49,1,99,'2020-09-10',NULL,NULL,NULL),(34,9,1,99,'2020-09-10',NULL,NULL,NULL),(35,10,1,99,'2020-09-10',NULL,NULL,NULL),(36,11,1,99,'2020-09-10',NULL,NULL,NULL),(37,47,1,1,'2020-09-11',1,22,3450),(38,48,1,1,'2020-09-11',2,4,4200000),(39,49,1,1,'2020-09-11',4,4,1600000),(40,9,1,1,'2020-09-11',6,6,600),(41,10,1,1,'2020-09-11',6,4,1100),(42,11,1,1,'2020-09-11',4,44,24000000),(43,47,1,2,'2020-09-11',1,NULL,150),(44,48,1,2,'2020-09-11',NULL,11,7700000),(45,49,1,2,'2020-09-11',12,NULL,2400000),(46,9,1,2,'2020-09-11',NULL,23,1150),(47,10,1,2,'2020-09-11',NULL,23,2530),(48,11,1,2,'2020-09-11',NULL,1,500000),(49,47,1,99,'2020-09-12',3,3,900),(50,48,1,99,'2020-09-12',4,4,5600000),(51,49,1,99,'2020-09-12',4,4,1600000),(52,9,1,99,'2020-09-12',4,4,400),(53,10,1,99,'2020-09-12',4,2,660),(54,11,1,99,'2020-09-12',2,2,2000000);
/*!40000 ALTER TABLE `data_visa` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `failed_jobs`
--

DROP TABLE IF EXISTS `failed_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `failed_jobs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `connection` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `failed_jobs`
--

LOCK TABLES `failed_jobs` WRITE;
/*!40000 ALTER TABLE `failed_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `failed_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `groupbymonthyear`
--

DROP TABLE IF EXISTS `groupbymonthyear`;
/*!50001 DROP VIEW IF EXISTS `groupbymonthyear`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `groupbymonthyear` AS SELECT 
 1 AS `periode`,
 1 AS `visa`,
 1 AS `paspor`,
 1 AS `izintinggal`,
 1 AS `pnbplainnya`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `inputdatatas`
--

DROP TABLE IF EXISTS `inputdatatas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inputdatatas` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tanggal` date DEFAULT NULL,
  `id_kategori` int NOT NULL,
  `laki_laki` int DEFAULT NULL,
  `perempuan` int DEFAULT NULL,
  `jumlah` int DEFAULT NULL,
  `user_id` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inputdatatas`
--

LOCK TABLES `inputdatatas` WRITE;
/*!40000 ALTER TABLE `inputdatatas` DISABLE KEYS */;
/*!40000 ALTER TABLE `inputdatatas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `izintinggal`
--

DROP TABLE IF EXISTS `izintinggal`;
/*!50001 DROP VIEW IF EXISTS `izintinggal`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `izintinggal` AS SELECT 
 1 AS `izintinggal`,
 1 AS `laki`,
 1 AS `perempuan`,
 1 AS `tanggal`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `kantor`
--

DROP TABLE IF EXISTS `kantor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kantor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_kantor` int DEFAULT NULL,
  `nama_kantor` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kantor`
--

LOCK TABLES `kantor` WRITE;
/*!40000 ALTER TABLE `kantor` DISABLE KEYS */;
INSERT INTO `kantor` VALUES (1,1,'pangkal pinang'),(2,2,'bangka belitung'),(3,99,'ALL');
/*!40000 ALTER TABLE `kantor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kategoripnbps`
--

DROP TABLE IF EXISTS `kategoripnbps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kategoripnbps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama_layanan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `satuan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parent` int NOT NULL DEFAULT '0',
  `tarifpnbp` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `status` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kategoripnbps`
--

LOCK TABLES `kategoripnbps` WRITE;
/*!40000 ALTER TABLE `kategoripnbps` DISABLE KEYS */;
INSERT INTO `kategoripnbps` VALUES (1,'DOKUMEN PERJALANAN REPUBLIK INDONESIA',NULL,0,NULL,NULL,NULL,0),(2,'VISA',NULL,0,NULL,NULL,NULL,0),(3,'IZIN KEIMIGRASIAN',NULL,0,NULL,NULL,NULL,0),(4,'PNBP KEIMIGRASIAN LAINNYA',NULL,0,NULL,NULL,NULL,0),(5,'Paspor Biasa 48 Halarnan','per permohonan',1,350000,NULL,NULL,0),(6,'Paspor Biasa 48 Halaman Elektronik','per permohonan',1,650000,NULL,NULL,0),(7,'Surat Perjalanan Laksana Paspor untuk Orang Asing','per permohonan',1,150000,NULL,NULL,0),(8,'Layanan Percepatan Paspor Selesai Pada Hari yang Sama','per permohonanper permohonan',1,1000000,NULL,NULL,0),(9,'Visa Kunjungan sekali Perjalanan','per permohonan',50,50,NULL,NULL,0),(10,'Visa Kunjungan Beberapa Kali Perjalanan Dihitung per Tahun',NULL,50,110,NULL,NULL,0),(11,'Visa Kunjungan Saat Kedatangan',NULL,50,500000,NULL,NULL,0),(12,'Izin Kunjungan',NULL,3,NULL,NULL,NULL,0),(13,'Pemberian Izin Kunjungan Masa Berlaku 30 hari',NULL,12,500000,NULL,NULL,0),(14,'Perpanjangan Izin Kunjungan Masa Berlaku 30 hari',NULL,12,500000,NULL,NULL,0),(15,'Perpanjangan Izin Kunjungan Masa Berlaku 60 hari',NULL,12,700000,NULL,NULL,0),(16,'Izin Tinggal Terbatas',NULL,3,NULL,NULL,NULL,0),(17,'ITAS Saat Kedatangan',NULL,16,750000,NULL,NULL,0),(18,'ITAS Masa Berlaku Paling Lama 6 Bulan',NULL,16,1000000,NULL,NULL,0),(19,'ITAS Masa Berlaku Paling Lama 1 Tahun',NULL,16,1500000,NULL,NULL,0),(20,'ITAS Masa Berlaku Paling Lama 2 Tahun',NULL,16,2000000,NULL,NULL,0),(21,'Persetujuan ITAS Untuk Pekerja di Perairan Indonesia',NULL,16,1000000,NULL,NULL,0),(22,'Teraan ITAS Untuk Pekerja di Perairan Indonesia',NULL,16,300000,NULL,NULL,0),(23,'Izin Tinggal Tetap',NULL,3,NULL,NULL,NULL,0),(24,'Pemberian ITAP Masa Berlaku 5 (Lima) Tahun',NULL,23,5000000,NULL,NULL,0),(25,'Perpanjangan ITAP Masa Berlaku 5 (Lima) Tahun',NULL,23,5000000,NULL,NULL,0),(26,'Perpanjangan ITAP untuk Jangka Waktu yang Tidak Terbatas',NULL,23,10200000,NULL,NULL,0),(27,'Izin Masuk Kembali (Re-Entry Permit)',NULL,3,NULL,NULL,NULL,0),(28,'Izin Masuk Kembali Masa Berlaku Paling Lama 6 (Enam) Bulan',NULL,27,600000,NULL,NULL,0),(29,'Izin Masuk Kembali Masa Berlaku Paling Lama 1 (Satu) Tahun',NULL,27,1000000,NULL,NULL,0),(30,'Izin Masuk Kembali Masa Berlaku Paling Lama 2 (Dua) Tahun',NULL,27,1750000,NULL,NULL,0),(31,'Izin Masuk Kembali Masa Berlaku 5 (Lima) Tahun Khusus pada Kawasan Ekonomi Khusus (KEK)',NULL,27,3250000,NULL,NULL,0),(32,'Biaya Beban',NULL,4,NULL,NULL,NULL,0),(33,'Orang Asing yang Berada di Wilayah Indonesia Melampaui Waktu Tidak lebih dari 60 Hari','per hari',32,1000000,NULL,NULL,0),(34,'Penanggung Jawab Alat Angkut yang Tidak Memenuhi Pasal 18',NULL,32,50000000,NULL,NULL,0),(35,'Penanggung Jawab Alat Angkut yang Mengangkut Penumpang yang Tidak Memiliki Dokumen Keimigrasian yang Sah dan Berlaku',NULL,32,50000000,NULL,NULL,0),(36,'Biaya Beban Paspor Hilang',NULL,32,1000000,NULL,NULL,0),(37,'Biaya Beban Paspor Rusak',NULL,32,500000,NULL,NULL,0),(38,'Biaya Beban Kartu Izin Tinggal Tetap Hilang',NULL,32,1000000,NULL,NULL,0),(39,'Biaya Beban KPP APEC Hilang atau Rusak',NULL,32,1000000,NULL,NULL,1),(40,'Smart Card','per permohonan',4,1500000,NULL,NULL,1),(41,'Kartu(KPP APEC) atau APEC Business Travel Card (ABTC)',NULL,4,NULL,NULL,NULL,0),(42,'Permohonan Baru KPP APEC','per permohonan',41,2500000,NULL,NULL,1),(43,'Penggantian KPP APEC','per permohonan',41,2500000,NULL,NULL,1),(44,'Fasilitas Keimigrasian (Afidavit)','per permohonan',4,400000,NULL,NULL,0),(45,'Surat Keterangan Keimigrasian','per permohonan',4,3000000,NULL,NULL,0),(46,'Visa Tinggal Terbatas',NULL,2,NULL,NULL,NULL,0),(47,'Visa Tinggal Terbatas','per permohonan',46,150,NULL,NULL,0),(48,'Visa Tinggal Terbatas Saat Kedatangan','per permohonan',46,700000,NULL,NULL,0),(49,'Persetujuan Visa Direktur Jenderal Imigrasi',NULL,2,200000,NULL,NULL,0),(50,'Visa Kunjungan',NULL,2,NULL,NULL,NULL,0),(51,'Izin Tinggal Terbatas Khusus Masa Berlaku Paling Lama 5 (lima) Tahun Khusus pada Kawasan Ekonomi Khusus (KEK)',NULL,16,5000000,NULL,NULL,0);
/*!40000 ALTER TABLE `kategoripnbps` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'2014_10_12_000000_create_users_table',1),(2,'2014_10_12_100000_create_password_resets_table',1),(3,'2019_08_19_000000_create_failed_jobs_table',1),(4,'2020_02_25_034148_create_permission_tables',1),(5,'2020_04_02_000000_create_two_factor_authentications_table',1),(6,'2020_05_25_021239_create_announcements_table',1),(7,'2020_05_29_020244_create_password_histories_table',1),(8,'2020_07_06_215139_create_activity_log_table',1),(9,'2020_09_05_144040_create_kategoripnbps_table',2);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `model_has_permissions`
--

DROP TABLE IF EXISTS `model_has_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `model_has_permissions` (
  `permission_id` bigint unsigned NOT NULL,
  `model_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`model_id`,`model_type`),
  KEY `model_has_permissions_model_id_model_type_index` (`model_id`,`model_type`),
  CONSTRAINT `model_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `model_has_permissions`
--

LOCK TABLES `model_has_permissions` WRITE;
/*!40000 ALTER TABLE `model_has_permissions` DISABLE KEYS */;
INSERT INTO `model_has_permissions` VALUES (11,'App\\Domains\\Auth\\Models\\User',2),(11,'App\\Domains\\Auth\\Models\\User',3),(11,'App\\Domains\\Auth\\Models\\User',4);
/*!40000 ALTER TABLE `model_has_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `model_has_roles`
--

DROP TABLE IF EXISTS `model_has_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `model_has_roles` (
  `role_id` bigint unsigned NOT NULL,
  `model_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`model_id`,`model_type`),
  KEY `model_has_roles_model_id_model_type_index` (`model_id`,`model_type`),
  CONSTRAINT `model_has_roles_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `model_has_roles`
--

LOCK TABLES `model_has_roles` WRITE;
/*!40000 ALTER TABLE `model_has_roles` DISABLE KEYS */;
INSERT INTO `model_has_roles` VALUES (1,'App\\Domains\\Auth\\Models\\User',1);
/*!40000 ALTER TABLE `model_has_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `paspor`
--

DROP TABLE IF EXISTS `paspor`;
/*!50001 DROP VIEW IF EXISTS `paspor`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `paspor` AS SELECT 
 1 AS `paspor`,
 1 AS `laki`,
 1 AS `perempuan`,
 1 AS `tanggal`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `password_histories`
--

DROP TABLE IF EXISTS `password_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `model_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint unsigned NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_histories`
--

LOCK TABLES `password_histories` WRITE;
/*!40000 ALTER TABLE `password_histories` DISABLE KEYS */;
INSERT INTO `password_histories` VALUES (1,'App\\Domains\\Auth\\Models\\User',1,'$2y$10$lT7CDscW4o2NBS99IQGSpO83lgjXmyLuQWAxQIJwJbLzglRbcexBW','2020-09-05 09:43:00','2020-09-05 09:43:00'),(2,'App\\Domains\\Auth\\Models\\User',2,'$2y$10$cXr/xNNNjkcjPxUW25eTGuG8xyHC7MPyvuYBtKzUvo2tqbfyRmilC','2020-09-05 09:43:00','2020-09-05 09:43:00'),(3,'App\\Domains\\Auth\\Models\\User',3,'$2y$10$UPC065z8zUJyk/tiXlclGuyNtVc913AtiWtLqnUWrlTwffNook.le','2020-09-08 11:52:59','2020-09-08 11:52:59'),(4,'App\\Domains\\Auth\\Models\\User',4,'$2y$10$ubMHd5PSpRr5.ZZGPfJpYeU.LMdVlVM8iqoh05OExqpGEYk8VQxTu','2020-09-10 01:32:43','2020-09-10 01:32:43');
/*!40000 ALTER TABLE `password_histories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `password_resets`
--

DROP TABLE IF EXISTS `password_resets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_resets` (
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  KEY `password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_resets`
--

LOCK TABLES `password_resets` WRITE;
/*!40000 ALTER TABLE `password_resets` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_resets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` enum('admin','user') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `guard_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL,
  `sort` tinyint NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `permissions_parent_id_foreign` (`parent_id`),
  CONSTRAINT `permissions_parent_id_foreign` FOREIGN KEY (`parent_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,'admin','web','admin.access.user','All User Permissions',NULL,1,'2020-09-05 09:43:00','2020-09-05 09:43:00'),(2,'admin','web','admin.access.user.list','View Users',1,1,'2020-09-05 09:43:00','2020-09-05 09:43:00'),(3,'admin','web','admin.access.user.deactivate','Deactivate Users',1,2,'2020-09-05 09:43:01','2020-09-05 09:43:01'),(4,'admin','web','admin.access.user.reactivate','Reactivate Users',1,3,'2020-09-05 09:43:01','2020-09-05 09:43:01'),(5,'admin','web','admin.access.user.clear-session','Clear User Sessions',1,4,'2020-09-05 09:43:01','2020-09-05 09:43:01'),(6,'admin','web','admin.access.user.impersonate','Impersonate Users',1,5,'2020-09-05 09:43:01','2020-09-05 09:43:01'),(7,'admin','web','admin.access.user.change-password','Change User Passwords',1,6,'2020-09-05 09:43:01','2020-09-05 09:43:01'),(8,'user','web','user.access.user','Semua Akses Petugas',NULL,7,NULL,NULL),(9,'user','web','user.access.list','Melihat list Data',8,8,NULL,NULL),(10,'user','web','user.access.input','Petugas Dapat mengentri',8,1,NULL,NULL),(11,'admin','web','admin.access.input','Petugas Dapat mengentri',NULL,1,NULL,NULL);
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `pnbp`
--

DROP TABLE IF EXISTS `pnbp`;
/*!50001 DROP VIEW IF EXISTS `pnbp`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `pnbp` AS SELECT 
 1 AS `id_jenis`,
 1 AS `laki`,
 1 AS `perempuan`,
 1 AS `total`,
 1 AS `tanggal`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `pnbplainnya`
--

DROP TABLE IF EXISTS `pnbplainnya`;
/*!50001 DROP VIEW IF EXISTS `pnbplainnya`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `pnbplainnya` AS SELECT 
 1 AS `pnbplainnya`,
 1 AS `laki`,
 1 AS `perempuan`,
 1 AS `tanggal`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `role_has_permissions`
--

DROP TABLE IF EXISTS `role_has_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_has_permissions` (
  `permission_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`role_id`),
  KEY `role_has_permissions_role_id_foreign` (`role_id`),
  CONSTRAINT `role_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_has_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_has_permissions`
--

LOCK TABLES `role_has_permissions` WRITE;
/*!40000 ALTER TABLE `role_has_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_has_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` enum('admin','user') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `guard_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin','Administrator','web','2020-09-05 09:43:00','2020-09-05 09:43:00');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `two_factor_authentications`
--

DROP TABLE IF EXISTS `two_factor_authentications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `two_factor_authentications` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `authenticatable_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `authenticatable_id` bigint unsigned NOT NULL,
  `shared_secret` blob NOT NULL,
  `enabled_at` timestamp NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `digits` tinyint unsigned NOT NULL DEFAULT '6',
  `seconds` tinyint unsigned NOT NULL DEFAULT '30',
  `window` tinyint unsigned NOT NULL DEFAULT '0',
  `algorithm` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'sha1',
  `recovery_codes` json DEFAULT NULL,
  `recovery_codes_generated_at` timestamp NULL DEFAULT NULL,
  `safe_devices` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `2fa_auth_type_auth_id_index` (`authenticatable_type`,`authenticatable_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `two_factor_authentications`
--

LOCK TABLES `two_factor_authentications` WRITE;
/*!40000 ALTER TABLE `two_factor_authentications` DISABLE KEYS */;
INSERT INTO `two_factor_authentications` VALUES (1,'App\\Domains\\Auth\\Models\\User',1,_binary '¶bpè§Ω«Å®pÑì™\‡\˜Ä⁄Åà',NULL,'admin@admin.com',6,30,1,'sha1',NULL,NULL,NULL,'2020-09-10 20:37:49','2020-09-10 20:37:49');
/*!40000 ALTER TABLE `two_factor_authentications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` enum('admin','user') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'user',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password_changed_at` timestamp NULL DEFAULT NULL,
  `active` tinyint unsigned NOT NULL DEFAULT '1',
  `timezone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_login_at` timestamp NULL DEFAULT NULL,
  `last_login_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `to_be_logged_out` tinyint(1) NOT NULL DEFAULT '0',
  `provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `provider_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_kantor` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','Super Admin','admin@admin.com','2020-09-05 09:43:00','$2y$10$lT7CDscW4o2NBS99IQGSpO83lgjXmyLuQWAxQIJwJbLzglRbcexBW',NULL,1,'America/New_York','2020-09-11 22:46:51','127.0.0.1',0,NULL,NULL,NULL,'2020-09-05 09:43:00','2020-09-11 22:46:51',NULL,99),(2,'admin','Test xxidg','user@user.com','2020-09-05 09:43:00','$2y$10$cXr/xNNNjkcjPxUW25eTGuG8xyHC7MPyvuYBtKzUvo2tqbfyRmilC',NULL,1,NULL,NULL,NULL,1,NULL,NULL,NULL,'2020-09-05 09:43:00','2020-09-10 01:25:38',NULL,1),(3,'admin','tescccc','tes@tes.com','2020-09-08 11:52:59','$2y$10$UPC065z8zUJyk/tiXlclGuyNtVc913AtiWtLqnUWrlTwffNook.le',NULL,1,'America/New_York','2020-09-08 12:47:25','127.0.0.1',0,NULL,NULL,NULL,'2020-09-08 11:52:59','2020-09-10 01:34:32',NULL,2),(4,'admin','tes','tes@tes.123','2020-09-10 01:32:43','$2y$10$ubMHd5PSpRr5.ZZGPfJpYeU.LMdVlVM8iqoh05OExqpGEYk8VQxTu',NULL,1,'America/New_York','2020-09-10 04:02:10','127.0.0.1',0,NULL,NULL,NULL,'2020-09-10 01:32:43','2020-09-10 04:02:10',NULL,2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `visa`
--

DROP TABLE IF EXISTS `visa`;
/*!50001 DROP VIEW IF EXISTS `visa`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `visa` AS SELECT 
 1 AS `visa`,
 1 AS `laki`,
 1 AS `perempuan`,
 1 AS `tanggal`*/;
SET character_set_client = @saved_cs_client;

--
-- Dumping routines for database 'ebindalwasmin'
--

--
-- Final view structure for view `groupbymonthyear`
--

/*!50001 DROP VIEW IF EXISTS `groupbymonthyear`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `groupbymonthyear` AS select concat(monthname(`pnbp`.`tanggal`),' ',year(`pnbp`.`tanggal`)) AS `periode`,(select sum(`visa`.`visa`) from `visa` where ((month(`visa`.`tanggal`) = month(`pnbp`.`tanggal`)) and (year(`visa`.`tanggal`) = year(`pnbp`.`tanggal`)))) AS `visa`,(select sum(`paspor`.`paspor`) from `paspor` where ((month(`paspor`.`tanggal`) = month(`pnbp`.`tanggal`)) and (year(`paspor`.`tanggal`) = year(`pnbp`.`tanggal`)))) AS `paspor`,(select sum(`izintinggal`.`izintinggal`) from `izintinggal` where ((month(`izintinggal`.`tanggal`) = month(`pnbp`.`tanggal`)) and (year(`izintinggal`.`tanggal`) = year(`pnbp`.`tanggal`)))) AS `izintinggal`,(select sum(`pnbplainnya`.`pnbplainnya`) from `pnbplainnya` where ((month(`pnbplainnya`.`tanggal`) = month(`pnbp`.`tanggal`)) and (year(`pnbplainnya`.`tanggal`) = year(`pnbp`.`tanggal`)))) AS `pnbplainnya` from `pnbp` group by year(`pnbp`.`tanggal`),month(`pnbp`.`tanggal`) order by `pnbp`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `izintinggal`
--

/*!50001 DROP VIEW IF EXISTS `izintinggal`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `izintinggal` AS select sum(`data_izinkeimigrasian`.`total`) AS `izintinggal`,sum(`data_izinkeimigrasian`.`laki`) AS `laki`,sum(`data_izinkeimigrasian`.`perempuan`) AS `perempuan`,`data_izinkeimigrasian`.`tanggal` AS `tanggal` from `data_izinkeimigrasian` group by `data_izinkeimigrasian`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `paspor`
--

/*!50001 DROP VIEW IF EXISTS `paspor`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `paspor` AS select sum(`data_paspor`.`total`) AS `paspor`,sum(`data_paspor`.`laki`) AS `laki`,sum(`data_paspor`.`perempuan`) AS `perempuan`,`data_paspor`.`tanggal` AS `tanggal` from `data_paspor` group by `data_paspor`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `pnbp`
--

/*!50001 DROP VIEW IF EXISTS `pnbp`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `pnbp` AS select `subquery`.`id_jenis` AS `id_jenis`,`subquery`.`laki` AS `laki`,`subquery`.`perempuan` AS `perempuan`,`subquery`.`total` AS `total`,`subquery`.`tanggal` AS `tanggal` from (select `data_paspor`.`id_jenis` AS `id_jenis`,`data_paspor`.`laki` AS `laki`,`data_paspor`.`perempuan` AS `perempuan`,`data_paspor`.`total` AS `total`,`data_paspor`.`tanggal` AS `tanggal` from `data_paspor` union all select `data_visa`.`id_jenis` AS `id_jenis`,`data_visa`.`laki` AS `laki`,`data_visa`.`perempuan` AS `perempuan`,`data_visa`.`total` AS `total`,`data_visa`.`tanggal` AS `tanggal` from `data_visa` union all select `data_izinkeimigrasian`.`id_jenis` AS `id_jenis`,`data_izinkeimigrasian`.`laki` AS `laki`,`data_izinkeimigrasian`.`perempuan` AS `perempuan`,`data_izinkeimigrasian`.`total` AS `total`,`data_izinkeimigrasian`.`tanggal` AS `tanggal` from `data_izinkeimigrasian` union all select `data_pnbplainnya`.`id_jenis` AS `id_jenis`,`data_pnbplainnya`.`laki` AS `laki`,`data_pnbplainnya`.`perempuan` AS `perempuan`,`data_pnbplainnya`.`total` AS `total`,`data_pnbplainnya`.`tanggal` AS `tanggal` from `data_pnbplainnya`) `subquery` order by `subquery`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `pnbplainnya`
--

/*!50001 DROP VIEW IF EXISTS `pnbplainnya`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `pnbplainnya` AS select sum(`data_pnbplainnya`.`total`) AS `pnbplainnya`,sum(`data_pnbplainnya`.`laki`) AS `laki`,sum(`data_pnbplainnya`.`perempuan`) AS `perempuan`,`data_pnbplainnya`.`tanggal` AS `tanggal` from `data_pnbplainnya` group by `data_pnbplainnya`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `visa`
--

/*!50001 DROP VIEW IF EXISTS `visa`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `visa` AS select sum(`data_visa`.`total`) AS `visa`,sum(`data_visa`.`laki`) AS `laki`,sum(`data_visa`.`perempuan`) AS `perempuan`,`data_visa`.`tanggal` AS `tanggal` from `data_visa` group by `data_visa`.`tanggal` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-07 23:51:38
