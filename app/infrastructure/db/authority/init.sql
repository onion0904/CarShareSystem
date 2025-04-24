-- 全スキーマ・全テーブルから権限をはく奪
REVOKE ALL PRIVILEGES, GRANT OPTION ON *.* FROM 'app_user'@'%';

-- 必要な権限を myapp_db スキーマ内だけに付与
GRANT SELECT, INSERT, UPDATE, DELETE ON myapp_db.* TO 'app_user'@'%';
