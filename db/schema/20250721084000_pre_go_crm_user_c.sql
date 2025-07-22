-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_crm_user_c` (
  `usr_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `usr_email` VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `usr_phone` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'Phone Number',
  `usr_username` VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'Username',
  `usr_password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Password (hashed)',
  `usr_created_at` INT(11) NOT NULL DEFAULT 0 COMMENT 'Creation Timestamp',
  `usr_updated_at` INT(11) NOT NULL DEFAULT 0 COMMENT 'Update Timestamp',
  `usr_create_ip_at` VARCHAR(45) NOT NULL DEFAULT '' COMMENT 'Creation IP',
  `usr_last_login_at` INT(11) NOT NULL DEFAULT 0 COMMENT 'Last Login Timestamp',
  `usr_last_login_ip_at` VARCHAR(45) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
  `usr_login_times` INT(11) NOT NULL DEFAULT 0 COMMENT 'Login Times',
  `usr_status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Status: 1=enabled, 0=disabled, -1=deleted',
  PRIMARY KEY (`usr_id`),
  KEY `idx_email` (`usr_email`),
  KEY `idx_phone` (`usr_phone`),
  KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User Account Table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_crm_user_c`;
-- +goose StatementEnd
