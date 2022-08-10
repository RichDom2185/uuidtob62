# uuidtob62

`uuidtob62` is a command-line utility to convert a UUID in standard (hexadecimal) representation to a base-62 string. It prints the output base-62 string to standard output. `uuidtob62` is written in Golang.

The main purpose of this utility is to provide a more compact representation of a UUID in use cases where there may be limited space available. A normal UUID consists of a total of 36 characters: 32 hexadecimal characters and 4 hyphens. The positions of these 4 hyphens are fixed, hence they can be dropped. The remaining characters are parsed as a base-16 number and converted to base 62 with the following radix, as specified in the [big package documentation](https://pkg.go.dev/math/big?utm_source=gopls#Int.Text):

```
0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
```

Base 62 was chosen to not have to worry about URL-safe characters, while still maintaining sufficient compatibility for most use-cases. Converting the UUID to base 62 results in a 22-character string, saving 39% space. In small, low-resolutions screens, this can make all the difference e.g. when trying to display a unique QR code.

**Note:** If all you need is guaranteed uniqueness, you can just use the base-62 string as a primary key. However, if you are intending to parse the base-62 string back to hexadecimal format, ensure that it handles large values nicely, as the the parsed value of a UUID overflows a standard 32-bit integer.

## Building

```bash
$ go build
```

## Usage

- Output base-62 string to a file:

  ```bash
  $ uuidtob62 input_file.txt > output_file.txt
  ```

- Read (hexadecimal) UUID from a file:

  ```bash
  $ uuidtob62 <file_name>
  ```

- Pass UUID string to standard input:

  ```bash
  $ echo '123e4567-e89b-12d3-a456-426614174000' | uuidtob62
  ```

When both a string STDIN and an file name argument are specified, the latter takes precedence.
