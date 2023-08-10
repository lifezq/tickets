-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.6.50-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- 导出  表 tickets.my_data 结构
CREATE TABLE IF NOT EXISTS `my_data` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `numbers` varchar(20) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_n` (`numbers`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- 正在导出表  tickets.my_data 的数据：~1 rows (大约)
INSERT INTO `my_data` (`id`, `numbers`) VALUES
                                            (10, '01,04,07,19,20,24,08'),
                                            (11, '02,06,12,16,17,26,15'),
                                            (12, '02,10,13,17,27,32,03'),
                                            (6, '03,10,11,18,19,33,05'),
                                            (9, '04,06,09,21,22,25,14'),
                                            (15, '04,06,12,13,19,29,09'),
                                            (3, '04,07,11,13,21,31,04'),
                                            (1, '04,09,18,28,30,31,06'),
                                            (5, '05,06,10,12,23,28,08'),
                                            (7, '05,13,18,25,29,30,16'),
                                            (13, '06,09,13,20,27,28,10'),
                                            (14, '06,18,19,25,29,33,01'),
                                            (4, '07,11,13,23,24,28,02'),
                                            (2, '09,11,14,19,30,31,07'),
                                            (8, '12,13,18,19,21,23,13');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
