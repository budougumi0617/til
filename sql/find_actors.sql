SELECT actor.first_name, actor.last_name
FROM film_actor, film, actor
WHERE film_actor.film_id = film.film_id
AND film_actor.actor_id = actor.actor_id
AND film.title = 'ALI FOREVER';

