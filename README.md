# Convert string to JSON object and return converted to Bytes[]

Objective:

- Write a go program that calls a function and prints the return value of that function. Feel free to have some creativity with this :-).

- Next, write a unit test for the function that is called.

- Lastly, write a Dockerfile that executes the go program.

- Please share a link to a Github repo with your completed code. 


Solution summary:

- Write a function that takes a `key` string and a `value` string eg. `text` and `Hello Slack!`
- Funtion converts the strings to a json object eg. `"{'text','Hello Slack!'}"` and converts it to byts bytes[] eg. `[123 34 116 101 120 116 34 58 34 72 101 108 108 111 32 83 108 97 99 107 33 34 125]`


Run the solution:

- Build and run the Docker image

```bash
docker build -t tobytes .
docker run --rm -it tobytes
```

- Output

```
[123 34 116 101 120 116 34 58 34 72 101 108 108 111 32 83 108 97 99 107 33 34 125]
```

- Watch the recording of it working!

```bash
python3 -m pip install asciinema
asciinema play asciinema-rec.cast
```

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem


