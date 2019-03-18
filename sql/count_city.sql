SELECT MAX(country.Name) AS name, COUNT(city.ID) AS count
FROM country
  INNER JOIN city
    ON country.Code = city.CountryCode
GROUP BY city.CountryCode
ORDER BY count DESC
LIMIT 10;
