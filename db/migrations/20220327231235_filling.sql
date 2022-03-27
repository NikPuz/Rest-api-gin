-- +goose Up

--
-- Дамп данных таблицы `albums`
--

INSERT INTO `albums` (`id`, `Title`, `Artist`, `Price`) VALUES
(5, 'Blue Train', 'John Coltrane', 666),
(4, 'Blue Train', 'John Coltrane', 56.99),
(13, 'oooow', 'aed evr', 12),
(19, 'sddaaaaaaaaa', 'aed evr', 12),
(11, 'saaaaaaaaa', 'uuuuuuuu ssssssssss', 0),
(10, 'uuuuuuuuuuuuu', 'rrrrrrrrrrrrrr rrrrrrrrrrrrrrr', 0),
(20, 'sddaaaaaaaaa', 'aed evr', 0),
(18, 'sssaaaaaaaaa', 'aad eer', 15),
(21, 'sddaaaaaaaaa', '', 0);

--
-- Дамп данных таблицы `user`
--

INSERT INTO `user` (`id`, `Name`, `Password`, `RefreshToken`, `ExpiresATToken`) VALUES
(4, 'mish', '$2a$15$XhkfvMtopj87Wg2P3Z8f9e.M6qXDOTTM9zflwEX.c9CLupI/p42MK', 'ed48b6e5d9e396b9c9a651786071ccb6b32aca7d0203c2ab48700dd46974738f4', '2021-12-17 03:46:40'),
(6, 'misha', '$2a$15$PMkIggMGHC9WOOncHFvFyuPgdmO9sYiiQRR6uie.BFLvGB8NEAglC', '899a9becef01afebfe94e100859a8b17c71679197c964a57babae7f3a1573c0b6', '2021-12-14 15:46:57'),
(17, 'nikk', '$2a$10$GIB2qnU.Pf2IUQvPVMtpguPTQiiewpLem4z5xXWaTJxuRlJgcheey', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDUyNjMwMjQsInVzZXJfaWQiOiIxNyJ9.v9mPuot41onj6Qlko9ymL2qXleidwpnmEVUg-g4YSqs', '2022-01-20 11:06:44'),
(10, 'nik', '$2a$10$8htS6L6xprUiCDwmc9GvFuk4rZ4FjuYkwQ6kBT/fIQ1JkNdUIp1RO', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDUzMDEzNzMsInVzZXJfaWQiOiIxMCJ9.mZEkCgIrgzhvwmdu2JOfF4XcfPVYCPAoVzmV_8gjC3w', '2022-01-20 11:00:31'),
(16, 'nikitos', '$2a$10$dWu6fSoTi7Vws/NE7qsg3OOuVjM/8e0WhHZen.zoN7k6JocqOp5T.', '54d2f1c7f34cc8b162d4abf3a773e06079270b037d2e7e57e9ba96eb895e5dcc16', '2021-12-24 14:25:21');

-- +goose Down

