# Bicycle

... is a marketplace for dishes from different restaurants.

Bicycle is designed as a platform to meet the demand and supply between consumers and sellers by bringing them to a
common platform.

At this first phase, we haven't implemented the delivery service and has built platform with the consideration that the
delivery will be managed by the seller.

## Requirements for development

Application works with Go 1.16.6 and MySQL 8.0.26

## Installation

1. Get Demo application source files from Git repository:

```
$ git clone https://github.com/YRIDZE/Bicycle-delivery-service.git
```

2. Create `logs` directory

```
$ cd Bicycle-delivery-service
$ mkdir logs
```

3. Add project configs:

- Create the `.evn` file; add and fill fields as suggested in the `.env.dist` file.
- Open `conf/config.yml` and replace **_port_** and **_db_** with you own data.


4. Migrate db

```
$ go build -o cli cmd/migrate/main.go
$ ./cli migrate create
$ ./cli migrate up
```

5. Build

```
$ cd public/bicycle
$ npm install
$ npm run build
```

## License

Distributed under the MIT License. See `LICENSE` for more information.