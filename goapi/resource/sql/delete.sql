DELETE FROM "Users"
	WHERE uid=4;
	
-- Reset SERIAL
TRUNCATE TABLE "Users" RESTART IDENTITY;