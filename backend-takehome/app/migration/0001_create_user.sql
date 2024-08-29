-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user`
(
    id          bigint          not null primary key AUTO_INCREMENT,
    name        varchar(50)     not null,
    email       varchar(100)    not null,
    password    varchar(1000)   not null,
    created_at  datetime        default current_timestamp,
    updated_at  datetime        default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user`;
-- +goose StatementEnd
