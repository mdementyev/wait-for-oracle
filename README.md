# wait-for-oracle

It's like [wait-for-it.sh](https://github.com/vishnubob/wait-for-it), but for Oracle databases.

## Requirements
* Go v1.18 or later

## Usage
To run the program, you need to provide the following information:

* Oracle database user
* Oracle database password
* Oracle database host
* Oracle database port (defaults to 1521)
* Oracle database SID
* Timeout in seconds (defaults to 120)

You can pass the information as command-line arguments:

```bash
./wait-for-oracle --user=<username> --password=<password> --host=<host> [--port=<port>] --sid=<sid> [--timeout=<timeout>]
```

You can also use environment variables to pass the information:

```bash
export ORACLE_USER=<username>
export ORACLE_PASSWORD=<password>
export ORACLE_HOST=<host>
export ORACLE_PORT=<port>
export ORACLE_SID=<sid>
export ORACLE_TIMEOUT=<timeout>

./wait-for-oracle
```

## Options
* `--user` or `ORACLE_USER`: Oracle database user
* `--password` or `ORACLE_PASSWORD`: Oracle database password
* `--host` or `ORACLE_HOST`: Oracle database host
* `--port` or `ORACLE_PORT`: Oracle database port (default: 1521)
* `--sid` or `ORACLE_SID`: Oracle database SID
* `--timeout` or `ORACLE_TIMEOUT`: Timeout in seconds (default: 120)

## Output
The program will either print "Connected" if the connection to the database is successful and exit with code 0, or "Timeout" if the connection could not be established within the specified timeout (exit code 1).

## License
This program is licensed under the MIT License.