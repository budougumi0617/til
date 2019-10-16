-- 4.9 練習問題
/**
 * 問題3 健康診断の順番
 * - ID順
 * - 男女別。男が先
 */
SELECT *
FROM members
ORDER BY gender DESC, id;

/**
 * 問題4 一覧ページ
 * - 身長を高い順
 * - 1ページに10人表示
 */
SELECT *
FROM members
ORDER BY height DESC
LIMIT 10 OFFSET 0; -- offsetを指定すると、それ以降の結果を取得できる

SELECT *
FROM members
ORDER BY height DESC
LIMIT 20 OFFSET 10; -- 10人目以降の結果を表示

/**
 * $ PGPASSWORD=secret psql -h localhost -U dev < ch04/practice.sql
 *  id  |   name   | height | gender
 * -----+----------+--------+--------
 *  101 | エレン   |    170 | M
 *  103 | アルミン |    163 | M
 *  104 | ジャン   |    175 | M
 *  106 | コニー   |    158 | M
 *  102 | ミカサ   |    170 | F
 *  105 | サシャ   |    168 | F
 * (6 rows)
 *
 *  id  |   name   | height | gender
 * -----+----------+--------+--------
 *  104 | ジャン   |    175 | M
 *  101 | エレン   |    170 | M
 *  102 | ミカサ   |    170 | F
 *  105 | サシャ   |    168 | F
 *  103 | アルミン |    163 | M
 *  106 | コニー   |    158 | M
 * (6 rows)
 *
 *  id | name | height | gender
 * ----+------+--------+--------
 * (0 rows)
 */
