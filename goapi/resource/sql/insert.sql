-- INSERT TO Controllers
INSERT INTO `Robots`(mac_addr)
	VALUES	('00:00:00:00:00:00'),
			('94:B9:7E:D5:AD:F4'),
			('01:01:01:01:01:01'),
			('02:02:02:02:02:02'),
			('03:03:03:03:03:03'),
			('04:04:04:04:04:04'),
			('05:05:05:05:05:05'),
			('06:06:06:06:06:06'),
			('07:07:07:07:07:07'),
			('08:08:08:08:08:08'),
			('09:09:09:09:09:09');


-- INSERT TO Users
INSERT INTO `Users`(urbt_id, uname, passwd, email, phone)
	VALUES	(1, 'dolly', '00000000', 'dolly0@radionoise.co', null),
			(2, 'last_order', 'LastOder20001', 'last_order20001@radionoise.co', null),
			(4, 'lester', 'abc456789', 'lester4869@protonmail.com', "0629310903"),
			(5, 'vermouth', '0464759045', 'chris_vineyard@blackorg.com', "0965484169"),
			(3, 'chianti', 'd_viper300', 'chianti300@blackorg.com', "0925300053"),
			(7, 'rum', 'sushi256', 'asaka56@blackorg.com', "0607502562"),
			(6, 'gin', 'p_356a1955', 'gin43680@blackorg.com', "0594869076"),
			(8, 'davidson', 'v-rod2003', 'davidson20030@blackorg.com', "0772733088"),
			(9, 'bourbon', 'fd_rx-7330', 'zero7310@blackorg.com', "0733100099"),
			(10, 'martini', 'aptx4869', 'martin4869@blackorg.com', "0486900010");


-- INSERT TO APIkeys
INSERT INTO `APIkeys`(api_uid, api_key)
	VALUES	(2, 'ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2'),
			(4, 'XS2-P9VR5XS2mJzJvpwl9bSrU3_bhg==');