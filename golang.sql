-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 22, 2022 at 05:58 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 7.3.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang`
--

-- --------------------------------------------------------

--
-- Table structure for table `category`
--

CREATE TABLE `category` (
  `category_id` int(11) NOT NULL,
  `category` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `category`
--

INSERT INTO `category` (`category_id`, `category`) VALUES
(1, 'Plain Color'),
(2, 'Pattern'),
(3, 'Figure');

-- --------------------------------------------------------

--
-- Table structure for table `gender`
--

CREATE TABLE `gender` (
  `gender_id` int(11) NOT NULL,
  `gender` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `gender`
--

INSERT INTO `gender` (`gender_id`, `gender`) VALUES
(1, 'Male'),
(2, 'Female');

-- --------------------------------------------------------

--
-- Table structure for table `order_buy`
--

CREATE TABLE `order_buy` (
  `order_buy_id` int(11) NOT NULL,
  `address` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `status` varchar(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `order_buy`
--

INSERT INTO `order_buy` (`order_buy_id`, `address`, `date`, `status`) VALUES
(7, '99 หมู่9 บ.ร่มเย็น ต.เป็นสุข อ.สุขสม จ.ดังใจ', '2022-04-20', '0'),
(8, '999 หมู่9 บ.ร่มเย็น ต.เป็นสุข อ.สุขสม จ.ดังใจ', '2022-04-22', '0'),
(10, '999 หมู่9 บ.ร่มเย็น ต.เป็นสุข อ.สุขสม จ.ดังใจ', '2022-04-22', '0');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `product_id` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `size_id` int(11) NOT NULL,
  `gender_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`product_id`, `product_name`, `size_id`, `gender_id`, `category_id`) VALUES
(1, 'product01', 1, 1, 1),
(2, 'product02', 2, 1, 3),
(3, 'product03', 3, 1, 1),
(4, 'product04', 4, 1, 1),
(5, 'product05', 5, 1, 1),
(6, 'product06', 1, 2, 2),
(7, 'product07', 2, 2, 2),
(8, 'product08', 3, 2, 2),
(9, 'product09', 4, 2, 3),
(10, 'product10', 5, 2, 3);

-- --------------------------------------------------------

--
-- Table structure for table `product_order_buy`
--

CREATE TABLE `product_order_buy` (
  `product_order_buy_id` int(11) NOT NULL,
  `order_buy_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `product_amount` varchar(7) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `product_order_buy`
--

INSERT INTO `product_order_buy` (`product_order_buy_id`, `order_buy_id`, `product_id`, `product_amount`) VALUES
(1, 7, 1, '2'),
(2, 7, 2, '4'),
(3, 8, 3, '1'),
(4, 8, 4, '1'),
(5, 10, 3, '1'),
(6, 10, 4, '1');

-- --------------------------------------------------------

--
-- Table structure for table `size`
--

CREATE TABLE `size` (
  `size_id` int(11) NOT NULL,
  `size` varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `size`
--

INSERT INTO `size` (`size_id`, `size`) VALUES
(1, 'XS'),
(2, 'S'),
(3, 'M'),
(4, 'L'),
(5, 'XL');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`category_id`);

--
-- Indexes for table `gender`
--
ALTER TABLE `gender`
  ADD PRIMARY KEY (`gender_id`);

--
-- Indexes for table `order_buy`
--
ALTER TABLE `order_buy`
  ADD PRIMARY KEY (`order_buy_id`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`product_id`);

--
-- Indexes for table `product_order_buy`
--
ALTER TABLE `product_order_buy`
  ADD PRIMARY KEY (`product_order_buy_id`);

--
-- Indexes for table `size`
--
ALTER TABLE `size`
  ADD PRIMARY KEY (`size_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `category_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `gender`
--
ALTER TABLE `gender`
  MODIFY `gender_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `order_buy`
--
ALTER TABLE `order_buy`
  MODIFY `order_buy_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `product_order_buy`
--
ALTER TABLE `product_order_buy`
  MODIFY `product_order_buy_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `size`
--
ALTER TABLE `size`
  MODIFY `size_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
