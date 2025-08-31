-- +goose Up
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;

CREATE TABLE `pre_go_acc_user_verify_9999` (
  `verify_id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `verify_otp` varchar(6) NOT NULL,
  `verify_key` varchar(255) NOT NULL,
  `verify_key_hash` varchar(255) NOT NULL,
  `verify_type` int DEFAULT '1' COMMENT '1: email, 2: phone',
  `is_verified` int DEFAULT '0',
  `is_deleted` int DEFAULT '0',
  `verify_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `verify_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `unique_verify_key` (`verify_key`),
  KEY `idx_verify_otp` (`verify_otp`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='account_user_verify';

-- +goose Down
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
