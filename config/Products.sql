/*
 Navicat Premium Data Transfer

 Source Server         : MyLocal
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : Restaurant

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 31/03/2020 19:21:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Products
-- ----------------------------
DROP TABLE IF EXISTS `Products`;
CREATE TABLE `Products` (
  `ID_PRO` varchar(16) NOT NULL,
  `PRO_NAME` varchar(50) DEFAULT NULL,
  `QTE_IN_STOCK` int(11) DEFAULT NULL,
  `PRO_IMG` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID_PRO`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of Products
-- ----------------------------
BEGIN;
INSERT INTO `Products` VALUES ('aCCCCICINGC', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('bCCCCICINGC', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CCCCICINGC', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CCCCICINGC2', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CCCCICINGC4', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CCCCICINGC67', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CCCCICINGC89', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CICING', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('CICINGC', 'SUMBAWA', 123, 'POER');
INSERT INTO `Products` VALUES ('lagu', 'ketan', 16, 'aslkfl');
INSERT INTO `Products` VALUES ('pisang12', 'pisang bakar', 12, 'balisojdoeidjdl');
INSERT INTO `Products` VALUES ('pisang13', 'sebakul nasi', 14, 'aslkdlfl');
INSERT INTO `Products` VALUES ('XCCCCICINGC', 'SUMBAWA', 123, 'POER');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
