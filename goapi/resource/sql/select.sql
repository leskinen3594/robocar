SELECT ctl_id, "mac_addr"
	FROM "Controllers";
	
SELECT * FROM "Users";


-- SELECT username, mac_addr FILTER BY api_key
SELECT ("Users".uname, "Controllers".mac_addr)
		  FROM (("Users" INNER JOIN "Controllers" ON "Users".uctl_id = "Controllers".ctl_id)
		  				 INNER JOIN "APIkeys" ON "Users".uid = "APIkeys".api_uid)
		  WHERE "APIkeys".api_key LIKE 'ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2';