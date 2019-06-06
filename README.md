## global-forward -auth

[![Build Status](https://cloud.drone.io/api/badges/devops-israel/global-forward-auth/status.svg)](https://cloud.drone.io/devops-israel/global-forward-auth)

A forward authentication service that provides Google and Github OAuth based login and authentication for Traefik and Nginx reverse proxies.
inspired by [traefik-forward-auth](https://github.com/thomseddon/traefik-forward-auth).

Use case: use your organization Google email or Githhub user to have the option to authenticate/authorize to any application you run in your organization that doesn't support a native authentication and authorization.

![Alt text](pic/diagram.png?raw=true "Title")

**Features**:
1. Central auth login - no need to add every subdomain to Google.
2. Use Github or Google OAuth.
3. Support both Nginx and Traefik reverse proxies which both work as K8s ingress (Nginx & Traefik)
4. Configure application via Environment Variables or Command line arguments or Config File.
5. Support a full domain and it’s subdomains by saving a cookie for the whole domain - the default behavior.
6. Check the user with Google/Github every 6 hours to make sure it is still authorized. (we will save the last time we checked in the cookie)
6. Authorization based on Email filter.
7. Authorization based on Github teams/organizations or Google groups.

**Technology Stack**:
1. Golang
2. Goth - https://github.com/markbates/goth
3. Iris - https://github.com/kataras/iris
4. Viper - https://github.com/spf13/viper
5. JWT - https://github.com/dgrijalva/jwt-go - used to sign the cookie value we save.


**Development Requirements**:
1. All development will be inside a Docker container.
2. Public Drone as CI server.
3. Add Testing to the project.
