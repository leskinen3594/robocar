SELECT * FROM `Users`;


-- SELECT username, mac_addr FILTER BY api_key
SELECT `Users`.uname, `Robots`.mac_addr
	FROM `Users`
	INNER JOIN `Robots`
		ON `Users`.urbt_id = `Robots`.rbt_id
	INNER JOIN `APIkeys`
		ON `Users`.uid = `APIkeys`.api_uid
	WHERE `APIkeys`.api_key LIKE 'ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2';