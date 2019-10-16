SELECT gender, MAX(height), MIN(height), SUM(height)
             , COUNT(*), TO_CHAR(AVG(height), '999.99')
FROM members
GROUP BY gender -- 性別でグループ化
ORDER BY gender DESC;

/*
 * $ PGPASSWORD=secret psql -h localhost -U dev < ch05/aggregate_functions.sql
 *  gender | max | min | sum | count | to_char
 * --------+-----+-----+-----+-------+---------
 *  M      | 175 | 158 | 666 |     4 |  166.50
 *  F      | 170 | 168 | 338 |     2 |  169.00
 * (2 rows)
 */
