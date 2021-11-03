# Pulu CLI
## Usage
### Login
You can easly login into the production or staging servers by running `pulu login [production/staging]` \
When logging in, the [proxy gateway](https://github.com/vives-projectwerk-2021/proxy-gateway) will check if you have the right permissions. (you must be in the devops group)

## Install from release
### Linux
```sh
sudo curl -L https://github.com/vives-projectwerk-2021/pulu-cli/releases/latest/download/pulu-linux -o /usr/local/bin/pulu && sudo chmod +x /usr/local/bin/pulu
```

### MacOS
```sh
sudo curl -L https://github.com/vives-projectwerk-2021/pulu-cli/releases/latest/download/pulu-macos -o /usr/local/bin/pulu && sudo chmod +x /usr/local/bin/pulu
```

### Windows

![make folder](./img/folderpulu.png)

![make folder](./img/pulu-exe.png)

![make folder](./img/search.png)

![make folder](./img/systemproperties.png)

![make folder](./img/editenv.png)


## Install from source
```sh
make
```
