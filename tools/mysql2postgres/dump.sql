-- root@ip-172-31-13-114:/home/ubuntu# mysqldump website_review --no-data
-- MySQL dump 10.19  Distrib 10.3.32-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: website_review
-- ------------------------------------------------------
-- Server version	10.3.32-MariaDB-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ca_cloud`
--

DROP TABLE IF EXISTS `ca_cloud`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_cloud` (
  `wid` int(10) unsigned NOT NULL,
  `words` mediumtext NOT NULL,
  `matrix` mediumtext NOT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_content`
--

DROP TABLE IF EXISTS `ca_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_content` (
  `wid` int(10) unsigned NOT NULL,
  `headings` mediumtext NOT NULL,
  `total_img` int(10) unsigned NOT NULL DEFAULT 0,
  `total_alt` int(10) unsigned NOT NULL DEFAULT 0,
  `deprecated` mediumtext NOT NULL,
  `isset_headings` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_document`
--

DROP TABLE IF EXISTS `ca_document`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_document` (
  `wid` int(10) unsigned NOT NULL,
  `doctype` text DEFAULT NULL,
  `lang` varchar(255) DEFAULT NULL,
  `charset` varchar(255) DEFAULT NULL,
  `css` int(10) unsigned NOT NULL DEFAULT 0,
  `js` int(10) unsigned NOT NULL DEFAULT 0,
  `htmlratio` int(10) unsigned NOT NULL DEFAULT 0,
  `favicon` text DEFAULT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_issetobject`
--

DROP TABLE IF EXISTS `ca_issetobject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_issetobject` (
  `wid` int(10) unsigned NOT NULL,
  `flash` tinyint(1) DEFAULT 0,
  `iframe` tinyint(1) DEFAULT 0,
  `nestedtables` tinyint(1) DEFAULT 0,
  `inlinecss` tinyint(1) DEFAULT 0,
  `email` tinyint(1) DEFAULT 0,
  `viewport` tinyint(1) DEFAULT 0,
  `dublincore` tinyint(1) DEFAULT 0,
  `printable` tinyint(1) DEFAULT 0,
  `appleicons` tinyint(1) DEFAULT 0,
  `robotstxt` tinyint(1) DEFAULT 0,
  `gzip` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_links`
--

DROP TABLE IF EXISTS `ca_links`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_links` (
  `wid` int(10) unsigned NOT NULL,
  `links` mediumtext NOT NULL,
  `internal` int(10) unsigned NOT NULL DEFAULT 0,
  `external_dofollow` int(10) unsigned NOT NULL DEFAULT 0,
  `external_nofollow` int(10) unsigned NOT NULL DEFAULT 0,
  `isset_underscore` tinyint(1) NOT NULL,
  `files_count` int(10) unsigned NOT NULL DEFAULT 0,
  `friendly` tinyint(1) NOT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_metatags`
--

DROP TABLE IF EXISTS `ca_metatags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_metatags` (
  `wid` int(10) unsigned NOT NULL,
  `title` mediumtext DEFAULT NULL,
  `keyword` mediumtext DEFAULT NULL,
  `description` mediumtext DEFAULT NULL,
  `ogproperties` mediumtext DEFAULT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_misc`
--

DROP TABLE IF EXISTS `ca_misc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_misc` (
  `wid` int(10) unsigned NOT NULL,
  `sitemap` mediumtext NOT NULL,
  `analytics` mediumtext NOT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_pagespeed`
--

DROP TABLE IF EXISTS `ca_pagespeed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_pagespeed` (
  `wid` int(10) unsigned NOT NULL,
  `data` longtext NOT NULL,
  `lang_id` varchar(5) NOT NULL,
  PRIMARY KEY (`wid`,`lang_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_w3c`
--

DROP TABLE IF EXISTS `ca_w3c`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_w3c` (
  `wid` int(10) unsigned NOT NULL,
  `validator` enum('html') NOT NULL,
  `valid` tinyint(1) NOT NULL DEFAULT 1,
  `errors` smallint(5) unsigned NOT NULL DEFAULT 0,
  `warnings` smallint(5) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ca_website`
--

DROP TABLE IF EXISTS `ca_website`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ca_website` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) DEFAULT NULL,
  `idn` varchar(255) DEFAULT NULL,
  `final_url` mediumtext DEFAULT NULL,
  `md5domain` varchar(32) DEFAULT NULL,
  `added` timestamp NULL DEFAULT NULL,
  `modified` timestamp NULL DEFAULT current_timestamp(),
  `score` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_md5domain` (`md5domain`),
  KEY `ix_rating` (`score`)
) ENGINE=InnoDB AUTO_INCREMENT=61417 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `domain`
--

DROP TABLE IF EXISTS `domain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `domain` (
  `id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT NULL,
  `last_crawl_date` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `feature`
--

DROP TABLE IF EXISTS `feature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `feature` (
  `id` int(11) NOT NULL,
  `name` varchar(32) DEFAULT NULL,
  `value` text DEFAULT NULL,
  `last_crawl_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `domain_id` int(11) DEFAULT NULL,
  `execution_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;