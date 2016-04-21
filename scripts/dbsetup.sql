CREATE TABLE `budget` (
  `budget_id` int(11) NOT NULL AUTO_INCREMENT,
  `budget_name` varchar(45) NOT NULL,
  `budget_amt` decimal(15,2) NOT NULL,
  PRIMARY KEY (`budget_id`),
  UNIQUE KEY `budget_id_UNIQUE` (`budget_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_fname` varchar(45) NOT NULL,
  `user_lname` varchar(45) NOT NULL,
  `user_email` varchar(100) NOT NULL,
  `user_phone` varchar(15) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

CREATE TABLE `category` (
  `cat_id` int(11) NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(45) NOT NULL,
  PRIMARY KEY (`cat_id`),
  UNIQUE KEY `cat_id_UNIQUE` (`cat_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

CREATE TABLE `expense` (
  `expense_id` int(11) NOT NULL AUTO_INCREMENT,
  `expense_amt` decimal(15,2) NOT NULL,
  `split_id` bigint(20) DEFAULT NULL,
  `expense_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`expense_id`),
  UNIQUE KEY `expense_id_UNIQUE` (`expense_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

CREATE TABLE `group` (
  `group_id` int(11) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(45) NOT NULL,
  PRIMARY KEY (`group_id`),
  UNIQUE KEY `group_id_UNIQUE` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

CREATE TABLE `budget_cat` (
  `cat_id` int(11) NOT NULL,
  `budget_id` int(11) NOT NULL,
  PRIMARY KEY (`cat_id`,`budget_id`),
  KEY `bc_budget_id_idx` (`budget_id`),
  CONSTRAINT `bc_budget_id` FOREIGN KEY (`budget_id`) REFERENCES `budget` (`budget_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `bc_cat_id` FOREIGN KEY (`cat_id`) REFERENCES `category` (`cat_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `budget_expenses` (
  `budget_id` int(11) NOT NULL,
  `expense_id` int(11) NOT NULL,
  PRIMARY KEY (`budget_id`,`expense_id`),
  KEY `be_expense_id_idx` (`expense_id`),
  CONSTRAINT `be_budget_id` FOREIGN KEY (`budget_id`) REFERENCES `budget` (`budget_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `be_expense_id` FOREIGN KEY (`expense_id`) REFERENCES `expense` (`expense_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `group_budgets` (
  `group_id` int(11) NOT NULL,
  `budget_id` int(11) NOT NULL,
  PRIMARY KEY (`group_id`,`budget_id`),
  KEY `budget_id_idx` (`budget_id`),
  CONSTRAINT `gb_budget_id` FOREIGN KEY (`budget_id`) REFERENCES `budget` (`budget_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `gb_group_id` FOREIGN KEY (`group_id`) REFERENCES `group` (`group_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `group_expenses` (
  `group_id` int(11) NOT NULL,
  `expense_id` int(11) NOT NULL,
  PRIMARY KEY (`group_id`,`expense_id`),
  KEY `expense_id_idx` (`expense_id`),
  CONSTRAINT `ge_expense_id` FOREIGN KEY (`expense_id`) REFERENCES `expense` (`expense_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ge_group_id` FOREIGN KEY (`group_id`) REFERENCES `group` (`group_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `group_members` (
  `group_id` int(11) NOT NULL,
  `member_id` int(11) NOT NULL,
  PRIMARY KEY (`group_id`,`member_id`),
  KEY `gm_user_id` (`member_id`),
  CONSTRAINT `gm_group_id` FOREIGN KEY (`group_id`) REFERENCES `group` (`group_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `gm_user_id` FOREIGN KEY (`member_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_budgets` (
  `user_id` int(11) NOT NULL,
  `budget_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`budget_id`),
  KEY `ub_budget_id_idx` (`budget_id`),
  CONSTRAINT `ub_budget_id` FOREIGN KEY (`budget_id`) REFERENCES `budget` (`budget_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ub_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_expenses` (
  `user_id` int(11) NOT NULL,
  `expense_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`expense_id`),
  KEY `user_expense_ibfk_2_idx` (`expense_id`),
  CONSTRAINT `ue_expense_id` FOREIGN KEY (`expense_id`) REFERENCES `expense` (`expense_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ue_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;