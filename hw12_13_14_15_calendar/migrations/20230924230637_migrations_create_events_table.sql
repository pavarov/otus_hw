-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events
(
    id                uuid PRIMARY KEY,
    title             varchar(255) not null,
    start             timestamp    not null,
    "end"             timestamp    not null,
    description       text         not null,
    user_id           uuid         not null,
    notification_time int8
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF exists events;
-- +goose StatementEnd
