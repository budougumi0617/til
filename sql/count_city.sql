SELECT MAX(country.country) AS name, COUNT(city.city_id) AS count
FROM country
  INNER JOIN city
    ON country.country_id = city.country_id
GROUP BY city.country_id
ORDER BY count DESC
LIMIT 10;
