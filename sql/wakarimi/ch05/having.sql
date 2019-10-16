SELECT gender, COUNT(*)
FROM members
GROUP BY gender -- 性別でグループ化
HAVING COUNT(*) >= 3; -- HAVING句はグループに対して実行される

/**
 * $ PGPASSWORD=secret psql -h localhost -U dev < ch05/having.sql
 *  gender | count
 * --------+-------
 *  M      |     4
 * (1 row)
 */
