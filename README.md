# Movie Finder

## Architecture

![Architecture diagram](movies-finder.drawio.svg)

## Solution

- Frontend
  - [Svelte](https://svelte.dev/)
    - [shadcn](https://www.shadcn-svelte.com/)
- API
  - [FastAPI](https://fastapi.tiangolo.com/)
- Recommendation Engine
  - [Go](https://golang.org/)
- Database
  - [MariaDB](https://mariadb.org/)
- Caching
  - [Redis](https://redis.io/)
  - [Go-Proxy-Cache](https://github.com/fabiocicerchia/go-proxy-cache)
- Auth & Idp
  - [Keycloak](https://www.keycloak.org/)
- Deployment
  - [Docker](https://www.docker.com/)
  - [Docker Compose](https://docs.docker.com/compose/)
  - [Traefik](https://www.nginx.com/)
- Observability
  - [Prometheus](https://prometheus.io/)
  - [Grafana](https://grafana.com/)
  - [Loki](https://grafana.com/loki/)
  - [Grafana-Agent](https://grafana.com/oss/grafana-agent/)

## Instructions
To start the service, get a [TMDB](https://www.themoviedb.org/) api key by logging in and going to [this page](https://www.themoviedb.org/settings/api). 
Then you can execute export TMDB_API_KEY=YOUR_TMDB_API_KEY. 

Then, to start the service, simply execute docker compose up --build. Wait until all the services are started, then head to http://what2watch.localhost/.

## Website navigation
There you will be prompted to login. You can register a new account, and then you will be redirected to the website. 

The first tab is the movie tab. It allows you to browse the top and popular movies, and to rate them by clicking "View details", then selecting a note from 1 to 5 and then clicking "Add to my movies".

You can also add your preferred genres by clicking on the icon at the top right of the page, and clicking "Preferences". In this page you can add genres to your favourite genres.

In the second tab you can browse the movies you added to your list. You can also remove a movie from your list.

Then, there is the recommendations page. But first, you need to create a group in the "Group" tab. To do that simply enter the group name, and click the "Create group" button. Then you can add people to the group by clicking "Add someone to this group", and entering the email address they used to register to this website. If you want locally to test with multiple users, you can create multiple accounts.

Once your group is created, you can go to the recommendations page, and click on the "Get recommendations" button in the group that you want to get recommendations for. 
