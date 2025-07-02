-- +goose Up
-- +goose StatementBegin
-- please delete this migration cause for example purpose
CREATE TABLE `example` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `phone` varchar(20) DEFAULT '',
  `address` text DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- +goose StatementEnd
INSERT INTO `example` (name, email, phone, address) VALUES ("jhon doe", "jhon.doe@mail.com","08123456789","PONDOK CABE");

-- +goose Down
-- +goose StatementBegin
DROP TABLE `example`;
-- +goose StatementEnd
