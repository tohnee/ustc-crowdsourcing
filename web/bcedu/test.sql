-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-06-07 21:21:44
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `test`
--

-- --------------------------------------------------------

--
-- 表的结构 `connects`
--

CREATE TABLE `connects` (
  `ID` int(11) NOT NULL,
  `orderid` int(11) NOT NULL,
  `connect` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `connects`
--

INSERT INTO `connects` (`ID`, `orderid`, `connect`) VALUES
(1, 37, '1763083159'),
(2, 37, '1763083159'),
(4, 38, '1763083159'),
(5, 38, '123456'),
(6, 38, '1234567'),
(7, 28, '123456'),
(8, 37, '123456'),
(9, 29, '54321'),
(10, 37, '54321'),
(11, 42, '123456'),
(12, 40, '123456');

-- --------------------------------------------------------

--
-- 表的结构 `orders`
--

CREATE TABLE `orders` (
  `ID` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `grade` varchar(20) NOT NULL,
  `gender` varchar(20) NOT NULL,
  `address` varchar(20) NOT NULL,
  `detailed` text NOT NULL,
  `tel` varchar(20) NOT NULL,
  `subject` varchar(80) NOT NULL,
  `timepay` text NOT NULL,
  `want` text NOT NULL,
  `time` datetime NOT NULL,
  `connect` varchar(200) NOT NULL,
  `succeed` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `orders`
--

INSERT INTO `orders` (`ID`, `name`, `grade`, `gender`, `address`, `detailed`, `tel`, `subject`, `timepay`, `want`, `time`, `connect`, `succeed`) VALUES
(28, 'Mary', '六年级', '女生', '包河区', '', '18551412373', '小学数学,高中数学,其它', '', '男女不限男女不限男女不限男女不限男女不限男女不限', '0000-00-00 00:00:00', '', 0),
(29, '张春花', '工作', '男生', '蜀山区', '', '18551412373', '区块', '', '', '0000-00-00 00:00:00', '', 0),
(30, '李小萌', '', '女生', '蜀山区', '', '18551412373', '', '', '', '0000-00-00 00:00:00', '', 0),
(34, 'hello', '大学', '男生', '蜀山区', '高数', '187', '', '数学竞赛', '', '2020-05-22 20:15:24', '', 0),
(42, '哈哈', '一年级', '男生', '蜀山区', '', '145', '', '', '', '2016-07-20 13:06:21', '', 0);

-- --------------------------------------------------------

--
-- 表的结构 `registrants`
--

CREATE TABLE `registrants` (
  `ID` int(11) NOT NULL,
  `user` varchar(20) NOT NULL,
  `pass` varchar(20) NOT NULL,
  `name` varchar(20) NOT NULL,
  `gender` varchar(20) NOT NULL,
  `school` varchar(20) NOT NULL,
  `grade` varchar(20) NOT NULL,
  `major` varchar(20) NOT NULL,
  `subject` varchar(80) NOT NULL,
  `aboutme` text NOT NULL,
  `push` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `registrants`
--

INSERT INTO `registrants` (`ID`, `user`, `pass`, `name`, `gender`, `school`, `grade`, `major`, `subject`, `aboutme`, `push`) VALUES
(36, '123456', '123', 'haha', '女生', '其它院校', '大一', 'haha', '小学语文', 'haha', 1),
(37, '123', '123', 'haha', '男生', '其它院校', '博士', 'lala ', '艺术类', 'haha', 1);

--
-- 转储表的索引
--

--
-- 表的索引 `connects`
--
ALTER TABLE `connects`
  ADD PRIMARY KEY (`ID`);

--
-- 表的索引 `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`ID`);

--
-- 表的索引 `registrants`
--
ALTER TABLE `registrants`
  ADD PRIMARY KEY (`ID`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `connects`
--
ALTER TABLE `connects`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `orders`
--
ALTER TABLE `orders`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=43;

--
-- 使用表AUTO_INCREMENT `registrants`
--
ALTER TABLE `registrants`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
