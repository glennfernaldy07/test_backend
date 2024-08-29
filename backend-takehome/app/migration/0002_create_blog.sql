-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `blog`
(
    id          bigint          not null primary key AUTO_INCREMENT,
    title       varchar(1000)   not null,
    content     text            not null,
    author_id   int             not null,
    created_at  datetime        default current_timestamp,
    updated_at  datetime        default current_timestamp on update current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `blog`;
-- +goose StatementEnd
